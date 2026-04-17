package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaoraclev3 "github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ws "github.com/gorilla/websocket"
	calendar "github.com/scmhub/calendar"
)

const (
	RWAWS         = "RWAWS"
	RAW_WS_CONFIG = "rawWSConfig.json"
	rwaWSURL      = "wss://ws.twelvedata.com/v1/quotes/price"
)

var (
	lastTimestampMu sync.Mutex
	lastTimestamp   int64
)

type RWAWSQuote struct {
	Symbol        string    `json:"Symbol"`
	Name          string    `json:"Name"`
	Price         float64   `json:"Price"`
	Time          time.Time `json:"Time"`
	ReceivedAt    time.Time `json:"ReceivedAt"`
	Source        string    `json:"Source"`
	Type          dataType  `json:"Type"`
	MarketHoliday bool      `json:"MarketHoliday"`
	MarketOpen    bool      `json:"MarketOpen"`
	Exchange      string    `json:"Exchange,omitempty"`
	MICCode       string    `json:"MICCode,omitempty"`
	CurrencyBase  string    `json:"CurrencyBase,omitempty"`
	CurrencyQuote string    `json:"CurrencyQuote,omitempty"`
	Currency      string    `json:"Currency,omitempty"`
}

type rwaWSMessage struct {
	Event         string          `json:"event"`
	Status        string          `json:"status"`
	Code          int             `json:"code"`
	Message       string          `json:"message"`
	Success       interface{}     `json:"success"`
	Fails         interface{}     `json:"fails"`
	Symbol        string          `json:"symbol"`
	Price         json.RawMessage `json:"price"`
	Timestamp     int64           `json:"timestamp"`
	Time          string          `json:"time"`
	Type          string          `json:"type"`
	CurrencyBase  string          `json:"currency_base"`
	CurrencyQuote string          `json:"currency_quote"`
	Currency      string          `json:"currency"`
	Exchange      string          `json:"exchange"`
	MICCode       string          `json:"mic_code"`
}

type rwaWSSubscribeMessage struct {
	Action string               `json:"action"`
	Params rwaWSSubscribeParams `json:"params"`
}

type rwaWSSubscribeParams struct {
	Symbols string `json:"symbols"`
}

type rwaWSHeartbeatMessage struct {
	Action string `json:"action"`
}

type RWAWSScraper struct {
	ctx    context.Context
	cancel context.CancelFunc

	apiKey string
	wsURL  string

	conn    *ws.Conn
	connMu  sync.RWMutex
	writeMu sync.Mutex

	closeOnce sync.Once
	closed    chan struct{}

	stockSymbols []string
	fxTickers    []string
	commodities  []string
	etfs         []string

	pendingMu     sync.Mutex
	pendingQuotes map[string]RWAWSQuote

	heartbeatTicker *time.Ticker
	configTicker    *time.Ticker

	publishCooldown time.Duration
	lastPublishTime time.Time
	flushTimer      *time.Timer
	flushTimerMu    sync.Mutex

	hkLoc *time.Location
	hkex  *calendar.Calendar

	// for oracle update
	auth        *bind.TransactOpts
	contractAny any
	chainId     int64
	source      string

	lastPublishedPrices map[string]float64
	lastPublishedTimes  map[string]time.Time
	forcePublishAfter   time.Duration

	decimals int64
}

func NewRWAWSScraper(auth *bind.TransactOpts, contractAny any, chainId int64, source string, decimals int64) *RWAWSScraper {
	ctx, cancel := context.WithCancel(context.Background())

	publishIntervalMs, err := strconv.Atoi(utils.Getenv("RWA_WS_PUBLISH_INTERVAL_MS", "600"))
	if err != nil {
		log.Errorf("parse RWA_WS_PUBLISH_INTERVAL_MS: %v", err)
		publishIntervalMs = 600
	}

	hkLoc, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		log.Errorf("load Asia/Hong_Kong location: %v", err)
		hkLoc = time.FixedZone("HKT", 8*3600)
	}

	forcePublishAfterSec, err := strconv.Atoi(utils.Getenv("RWA_WS_FORCE_PUBLISH_AFTER_SECONDS", "30"))
	if err != nil {
		forcePublishAfterSec = 30
	}

	s := &RWAWSScraper{
		ctx:                 ctx,
		cancel:              cancel,
		apiKey:              utils.Getenv("TWELVEDATA_API_KEY", ""),
		wsURL:               utils.Getenv("TWELVEDATA_WS_URL", rwaWSURL),
		closed:              make(chan struct{}),
		pendingQuotes:       make(map[string]RWAWSQuote),
		publishCooldown:     time.Duration(publishIntervalMs) * time.Millisecond,
		hkLoc:               hkLoc,
		hkex:                calendar.XHKG(),
		auth:                auth,
		contractAny:         contractAny,
		chainId:             chainId,
		source:              source,
		lastPublishedPrices: make(map[string]float64),
		lastPublishedTimes:  make(map[string]time.Time),
		forcePublishAfter:   time.Duration(forcePublishAfterSec) * time.Second,
		decimals:            decimals,
	}

	if s.apiKey == "" {
		log.Fatal("TWELVEDATA_API_KEY not set")
	}

	if err := s.updateConfig(RAW_WS_CONFIG); err != nil {
		log.Fatal("Could not load configuration file: ", err)
	}

	go s.mainLoop()
	return s
}

func (scraper *RWAWSScraper) Close() error {
	var err error
	scraper.closeOnce.Do(func() {
		close(scraper.closed)
		scraper.cancel()

		if scraper.heartbeatTicker != nil {
			scraper.heartbeatTicker.Stop()
		}
		if scraper.configTicker != nil {
			scraper.configTicker.Stop()
		}

		scraper.flushTimerMu.Lock()
		if scraper.flushTimer != nil {
			scraper.flushTimer.Stop()
			scraper.flushTimer = nil
		}
		scraper.flushTimerMu.Unlock()

		scraper.connMu.Lock()
		defer scraper.connMu.Unlock()
		if scraper.conn != nil {
			err = scraper.conn.Close()
			scraper.conn = nil
		}
	})
	return err
}

func (scraper *RWAWSScraper) mainLoop() {
	if err := scraper.connectAndSubscribe(); err != nil {
		log.Fatal("RWAWS connectAndSubscribe: ", err)
	}

	scraper.heartbeatTicker = time.NewTicker(10 * time.Second)
	scraper.configTicker = time.NewTicker(time.Duration(configUpdateSeconds) * time.Second)

	for {
		select {
		case <-scraper.ctx.Done():
			return
		case <-scraper.closed:
			return

		case <-scraper.heartbeatTicker.C:
			if err := scraper.sendHeartbeat(); err != nil {
				log.Errorf("RWAWS sendHeartbeat: %v", err)
			}

		case <-scraper.configTicker.C:
			if err := scraper.updateConfig(RAW_WS_CONFIG); err != nil {
				log.Errorf("RWAWS updateConfig: %v", err)
			}

		default:
			if err := scraper.readMessage(); err != nil {
				log.Errorf("RWAWS readMessage: %v", err)
				select {
				case <-scraper.ctx.Done():
					return
				case <-scraper.closed:
					return
				default:
				}
				time.Sleep(2 * time.Second)
				if err := scraper.reconnect(); err != nil {
					log.Errorf("RWAWS reconnect: %v", err)
				}
			}
		}
	}
}

func (scraper *RWAWSScraper) connectAndSubscribe() error {
	u, err := url.Parse(scraper.wsURL)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("apikey", scraper.apiKey)
	u.RawQuery = q.Encode()

	conn, _, err := ws.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	scraper.connMu.Lock()
	scraper.conn = conn
	scraper.connMu.Unlock()

	log.Infof("RWAWS - WebSocket connected")
	return scraper.subscribeAll()
}

func (scraper *RWAWSScraper) reconnect() error {
	scraper.connMu.Lock()
	if scraper.conn != nil {
		_ = scraper.conn.Close()
		scraper.conn = nil
	}
	scraper.connMu.Unlock()

	return scraper.connectAndSubscribe()
}

func (scraper *RWAWSScraper) subscribeAll() error {
	symbols := scraper.allSymbols()
	if len(symbols) == 0 {
		return fmt.Errorf("no symbols configured")
	}

	msg := rwaWSSubscribeMessage{
		Action: "subscribe",
		Params: rwaWSSubscribeParams{
			Symbols: strings.Join(symbols, ","),
		},
	}

	b, _ := json.Marshal(msg)
	log.Infof("RWAWS - subscribing payload: %s", string(b))

	return scraper.writeJSON(msg)
}

func (scraper *RWAWSScraper) sendHeartbeat() error {
	return scraper.writeJSON(rwaWSHeartbeatMessage{Action: "heartbeat"})
}

func (scraper *RWAWSScraper) writeJSON(v interface{}) error {
	scraper.writeMu.Lock()
	defer scraper.writeMu.Unlock()

	scraper.connMu.RLock()
	defer scraper.connMu.RUnlock()

	if scraper.conn == nil {
		return fmt.Errorf("ws connection is nil")
	}
	return scraper.conn.WriteJSON(v)
}

func (scraper *RWAWSScraper) readMessage() error {
	scraper.connMu.RLock()
	conn := scraper.conn
	scraper.connMu.RUnlock()

	if conn == nil {
		return fmt.Errorf("ws connection is nil")
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	log.Infof("RWAWS - raw ws payload: %s", string(message))

	var msg rwaWSMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil
	}

	switch strings.ToLower(strings.TrimSpace(msg.Event)) {
	case "heartbeat":
		return nil
	case "subscribe-status":
		log.Infof("RWAWS - subscribe-status: status=%s code=%d msg=%s success=%v fails=%v",
			msg.Status, msg.Code, msg.Message, msg.Success, msg.Fails)
		return nil
	case "price":
		return scraper.handlePriceMessage(msg)
	default:
		return nil
	}
}

func (scraper *RWAWSScraper) handlePriceMessage(msg rwaWSMessage) error {
	price, err := parseRawFloat(msg.Price)
	if err != nil {
		return err
	}

	quoteTime, validTs := parseTimestamp(msg.Timestamp)
	if !validTs {
		log.Warnf("RWAWS - dropping quote for %s: missing or invalid timestamp", msg.Symbol)
		return nil
	}

	quote := RWAWSQuote{
		Symbol:        msg.Symbol,
		Name:          chooseName(msg),
		Price:         price,
		Time:          quoteTime,
		ReceivedAt:    time.Now().UTC(),
		Source:        RWAWS,
		Exchange:      msg.Exchange,
		MICCode:       msg.MICCode,
		CurrencyBase:  msg.CurrencyBase,
		CurrencyQuote: msg.CurrencyQuote,
		Currency:      msg.Currency,
	}

	switch {
	case contains(scraper.stockSymbols, msg.Symbol):
		quote.Type = Equities

		marketHoliday, marketOpen, knownMarket := scraper.getStockMarketStatus(msg)
		if knownMarket {
			quote.MarketHoliday = marketHoliday
			quote.MarketOpen = marketOpen
			if !marketOpen {
				return nil
			}
		} else {
			// Unknown market: do not block publishing in v1.
			quote.MarketHoliday = false
			quote.MarketOpen = true
		}

	case contains(scraper.etfs, msg.Symbol):
		quote.Type = ETF

		marketHoliday, marketOpen, knownMarket := scraper.getStockMarketStatus(msg)
		if knownMarket {
			quote.MarketHoliday = marketHoliday
			quote.MarketOpen = marketOpen
			if !marketOpen {
				return nil
			}
		} else {
			quote.MarketHoliday = false
			quote.MarketOpen = true
		}

	case contains(scraper.fxTickers, msg.Symbol):
		quote.Type = Fiat

	case contains(scraper.commodities, msg.Symbol):
		quote.Type = Commodities

	default:
		log.Warnf("RWAWS - received unconfigured symbol: %s", msg.Symbol)
		return nil
	}

	scraper.pendingMu.Lock()
	scraper.pendingQuotes[quote.Symbol] = quote
	scraper.pendingMu.Unlock()

	scraper.maybePublishNowOrSchedule()
	return nil
}

// Publish immediately if cooldown already passed.
// Otherwise ensure a one-shot timer exists to publish once cooldown ends.
func (scraper *RWAWSScraper) maybePublishNowOrSchedule() {
	now := time.Now()

	scraper.pendingMu.Lock()
	hasPending := len(scraper.pendingQuotes) > 0
	lastPublish := scraper.lastPublishTime
	scraper.pendingMu.Unlock()

	if !hasPending {
		return
	}

	if lastPublish.IsZero() || now.Sub(lastPublish) >= scraper.publishCooldown {
		scraper.publishPendingBatch()
		return
	}

	remaining := scraper.publishCooldown - now.Sub(lastPublish)
	scraper.scheduleFlushAfter(remaining)
}

func (scraper *RWAWSScraper) scheduleFlushAfter(delay time.Duration) {
	scraper.flushTimerMu.Lock()
	defer scraper.flushTimerMu.Unlock()

	if scraper.flushTimer != nil {
		return
	}

	scraper.flushTimer = time.AfterFunc(delay, func() {
		scraper.publishPendingBatch()

		scraper.flushTimerMu.Lock()
		scraper.flushTimer = nil
		scraper.flushTimerMu.Unlock()

		// In case more messages arrived while publishing and cooldown already passed
		// by the time we finished, check again.
		scraper.maybePublishNowOrSchedule()
	})
}

func (scraper *RWAWSScraper) publishPendingBatch() {
	scraper.pendingMu.Lock()
	if len(scraper.pendingQuotes) == 0 {
		scraper.pendingMu.Unlock()
		return
	}

	toFlush := make([]RWAWSQuote, 0, len(scraper.pendingQuotes))
	for _, q := range scraper.pendingQuotes {
		toFlush = append(toFlush, q)
	}
	scraper.pendingQuotes = make(map[string]RWAWSQuote)
	scraper.lastPublishTime = time.Now()
	scraper.pendingMu.Unlock()

	log.Infof("RWAWS - publishing batch of %d quotes", len(toFlush))

	now := time.Now()
	filtered := []RWAWSQuote{}

	for _, quote := range toFlush {
		last, exists := scraper.lastPublishedPrices[quote.Symbol]
		lastTime := scraper.lastPublishedTimes[quote.Symbol]

		priceChanged := !exists || last != quote.Price
		timedOut := exists && now.Sub(lastTime) >= scraper.forcePublishAfter

		if priceChanged || timedOut {
			filtered = append(filtered, quote)
			scraper.lastPublishedPrices[quote.Symbol] = quote.Price
			scraper.lastPublishedTimes[quote.Symbol] = now
		}
	}

	if len(filtered) == 0 {
		log.Infof("RWAWS - all prices unchanged, skipping oracle update")
		return
	}

	keys := []string{}
	values := []*big.Float{}
	marketStatusAdded := false

	for _, quote := range filtered {
		k, v := scraper.preparePublishData(quote, &marketStatusAdded)
		keys = append(keys, k...)
		values = append(values, v...)
	}

	scraper.publishData(keys, values)
}

func (scraper *RWAWSScraper) preparePublishData(rwaResponse RWAWSQuote, marketStatusAdded *bool) (keys []string, values []*big.Float) {
	log.Info("got rwa ws data: ", rwaResponse)

	if (rwaResponse.Type == Equities || rwaResponse.Type == ETF) && !*marketStatusAdded {
		keys = append(keys, "Market_Open")
		marketOpen := new(big.Float).SetFloat64(0)
		if rwaResponse.MarketOpen {
			marketOpen = new(big.Float).SetFloat64(1)
		}
		values = append(values, marketOpen)

		keys = append(keys, "Market_Holiday")
		marketHoliday := new(big.Float).SetFloat64(0)
		if rwaResponse.MarketHoliday {
			marketHoliday = new(big.Float).SetFloat64(1)
		}
		values = append(values, marketHoliday)

		*marketStatusAdded = true
	}

	keys = append(keys, rwaResponse.Symbol)
	values = append(values, new(big.Float).SetFloat64(rwaResponse.Price))
	return keys, values
}

func (scraper *RWAWSScraper) publishData(keys []string, values []*big.Float) {
	log.Infof("collected %v responses. make oracle update...", len(values))

	switch contract := scraper.contractAny.(type) {
	case *diaoraclev3.DiaOracleV3MultiupdateService:
		keysCopy := make([]string, len(keys))
		copy(keysCopy, keys)
		valuesCopy := make([]*big.Float, len(values))
		copy(valuesCopy, values)

		ts := nextTimestamp()

		go func() {
			err := updateOracleMultiValuesForRWAWS(*contract, scraper.auth, keysCopy, valuesCopy, ts, scraper.decimals)
			if err != nil {
				log.Warnf("updater - Failed to update Oracle: %v.", err)
			}
		}()
	}
}

func (scraper *RWAWSScraper) updateConfig(filename string) error {
	c, err := models.GetRWAConfig(filename)
	if err != nil {
		return err
	}

	scraper.stockSymbols = c.Stocks
	scraper.fxTickers = c.FX
	scraper.commodities = c.Commodities
	scraper.etfs = c.ETF

	log.Infof("RWAWS - loaded config: stocks=%d fx=%d commodities=%d etf=%d",
		getNumSymbols(scraper.stockSymbols),
		getNumSymbols(scraper.fxTickers),
		getNumSymbols(scraper.commodities),
		getNumSymbols(scraper.etfs),
	)

	return nil
}

func (scraper *RWAWSScraper) allSymbols() []string {
	seen := make(map[string]struct{})
	out := make([]string, 0)

	appendUnique := func(values []string) {
		for _, v := range values {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}

	appendUnique(scraper.stockSymbols)
	appendUnique(scraper.fxTickers)
	appendUnique(scraper.commodities)
	appendUnique(scraper.etfs)

	return out
}

// Returns (holiday, open, knownMarket)
func (scraper *RWAWSScraper) getStockMarketStatus(msg rwaWSMessage) (bool, bool, bool) {
	mic := strings.ToUpper(strings.TrimSpace(msg.MICCode))
	exchange := strings.ToUpper(strings.TrimSpace(msg.Exchange))
	now := time.Now()

	// HKEX / XHKG
	if mic == "XHKG" || exchange == "HKEX" {
		holiday := scraper.isHKHoliday(now)
		open := scraper.isHKMarketOpen(now)
		return holiday, open, true
	}

	// Unknown market in v1
	return false, false, false
}

func (scraper *RWAWSScraper) isHKHoliday(now time.Time) bool {
	now = now.In(scraper.hkLoc)
	return !scraper.hkex.IsBusinessDay(now)
}

func (scraper *RWAWSScraper) isHKMarketOpen(now time.Time) bool {
	now = now.In(scraper.hkLoc)

	if !scraper.hkex.IsBusinessDay(now) {
		return false
	}

	minutes := now.Hour()*60 + now.Minute()
	inMorning := minutes >= 9*60+30 && minutes < 12*60
	inAfternoon := minutes >= 13*60 && minutes < 16*60

	return inMorning || inAfternoon
}

func parseRawFloat(raw json.RawMessage) (float64, error) {
	s := strings.TrimSpace(string(raw))
	if s == "" || s == "null" {
		return 0, fmt.Errorf("empty price")
	}
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		unquoted, err := strconv.Unquote(s)
		if err != nil {
			return 0, err
		}
		return strconv.ParseFloat(unquoted, 64)
	}
	return strconv.ParseFloat(s, 64)
}

func parseTimestamp(ts int64) (time.Time, bool) {
	if ts <= 0 {
		return time.Time{}, false
	}
	return time.Unix(ts, 0).UTC(), true
}

func contains(values []string, target string) bool {
	target = strings.TrimSpace(target)
	for _, v := range values {
		if strings.TrimSpace(v) == target {
			return true
		}
	}
	return false
}

func chooseName(msg rwaWSMessage) string {
	if strings.TrimSpace(msg.CurrencyBase) != "" {
		return msg.CurrencyBase
	}
	if strings.TrimSpace(msg.Symbol) != "" {
		return msg.Symbol
	}
	return "unknown"
}

func nextTimestamp() int64 {
	lastTimestampMu.Lock()
	defer lastTimestampMu.Unlock()
	now := time.Now().Unix()
	if now <= lastTimestamp {
		lastTimestamp++
	} else {
		lastTimestamp = now
	}
	return lastTimestamp
}

func updateOracleMultiValuesForRWAWS(
	contract diaoraclev3.DiaOracleV3MultiupdateService,
	auth *bind.TransactOpts,
	keys []string,
	values []*big.Float,
	timestamp int64,
	decimals int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int

	multiplier := new(big.Float).SetInt(
		new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil),
	)

	for _, value := range values {
		scaled := new(big.Float).Mul(value, multiplier)
		scaledInt, _ := scaled.Int(nil)

		// Create compressed argument with values/timestamps
		cValue := new(big.Int).Lsh(scaledInt, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	// check if tx is sendable then fgo backup
	if err != nil {
		// backup in here
		return err
	}

	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
	return nil
}
