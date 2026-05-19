package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
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
	rwawsConfigUpdateSeconds int
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

	hkStocks    []string
	usStocks    []string
	fxTickers   []string
	commodities []string
	usEtfs      []string

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
	deviationThresholds map[string]float64

	publishTrigger chan struct{}

	lastSubmittedTs int64
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

	rwawsConfigUpdateSeconds, err = strconv.Atoi(utils.Getenv("RWA_WS_CONFIG_UPDATE_SECONDS", "86400"))
	if err != nil {
		log.Errorf("parse RWA_WS_CONFIG_UPDATE_SECONDS: %v", err)
		rwawsConfigUpdateSeconds = 86400
	}

	s := &RWAWSScraper{
		ctx:                     ctx,
		cancel:                  cancel,
		apiKey:                  utils.Getenv("TWELVEDATA_API_KEY", ""),
		wsURL:                   utils.Getenv("TWELVEDATA_WS_URL", rwaWSURL),
		closed:                  make(chan struct{}),
		pendingQuotes:           make(map[string]RWAWSQuote),
		publishCooldown:         time.Duration(publishIntervalMs) * time.Millisecond,
		hkLoc:                   hkLoc,
		hkex:                    calendar.XHKG(),
		auth:                    auth,
		contractAny:             contractAny,
		chainId:                 chainId,
		source:                  source,
		lastPublishedPrices:     make(map[string]float64),
		lastPublishedTimes:      make(map[string]time.Time),
		forcePublishAfter:       time.Duration(forcePublishAfterSec) * time.Second,
		decimals:                decimals,
	}

	if s.apiKey == "" {
		log.Fatal("TWELVEDATA_API_KEY not set")
	}

	if err := s.updateConfig(RAW_WS_CONFIG); err != nil {
		log.Fatal("Could not load configuration file: ", err)
	}

	s.heartbeatTicker = time.NewTicker(10 * time.Second)
	s.publishTrigger = make(chan struct{}, 1)

	go s.mainLoop()
	go s.heartbeatLoop()
	go s.publishLoop()
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

func (scraper *RWAWSScraper) heartbeatLoop() {
	for {
		select {
		case <-scraper.ctx.Done():
			return
		case <-scraper.closed:
			return
		case <-scraper.heartbeatTicker.C:
			if err := scraper.sendHeartbeat(); err != nil {
				log.Debugf("RWAWS sendHeartbeat: %v", err)
			}
		}
	}
}

func (scraper *RWAWSScraper) mainLoop() {
	if err := scraper.connectAndSubscribe(); err != nil {
		log.Fatal("RWAWS connectAndSubscribe: ", err)
	}

	scraper.configTicker = time.NewTicker(time.Duration(rwawsConfigUpdateSeconds) * time.Second)

	for {
		select {
		case <-scraper.ctx.Done():
			return
		case <-scraper.closed:
			return
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
				scraper.handleDisconnect()
			}
		}
	}
}

func (scraper *RWAWSScraper) publishLoop() {
    for {
        select {
        case <-scraper.ctx.Done():
            return
        case <-scraper.closed:
            return
        case <-scraper.publishTrigger:
            scraper.publishPendingBatch()
        }
    }
}

func (scraper *RWAWSScraper) triggerPublish() {
    select {
    case scraper.publishTrigger <- struct{}{}:
    default:
    }
}

func (scraper *RWAWSScraper) publishData(keys []string, values []*big.Float) {
	log.Infof("collected %v responses. make oracle update...", len(values))

	switch contract := scraper.contractAny.(type) {
	case diaoraclev3.DIAOracleV3:
		err := scraper.updateOracleMultiValuesForRWAWS(
			contract, scraper.auth,
			keys, values,
			scraper.decimals,
		)
		if err != nil {
			log.Warnf("updater - Failed to update Oracle: %v.", err)
		}
	default:
		log.Errorf("RWAWS - unexpected contract type: %T", contract)
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
		return fmt.Errorf("failed to dial WebSocket: %v", err)
	}

	scraper.connMu.Lock()
	scraper.conn = conn
	scraper.connMu.Unlock()

	log.Infof("RWAWS - WebSocket connected to %s", u.String())
	return scraper.subscribeAll()
}

func (scraper *RWAWSScraper) reconnect() error {
	log.Infof("RWAWS - closing existing connection before reconnecting")
	scraper.connMu.Lock()
	if scraper.conn != nil {
		_ = scraper.conn.Close()
		scraper.conn = nil
	}
	scraper.connMu.Unlock()

	return scraper.connectAndSubscribe()
}

func (scraper *RWAWSScraper) handleDisconnect() {
	const maxRetries = 3
	backoff := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		select {
		case <-scraper.ctx.Done():
			return
		case <-scraper.closed:
			return
		default:
		}

		log.Warnf("RWAWS - reconnecting... (attempt %d of %d) after %v", i+1, maxRetries, backoff)
		time.Sleep(backoff)

		if err := scraper.reconnect(); err != nil {
			log.Errorf("RWAWS - reconnect attempt %d failed: %v", i+1, err)
			if backoff < 60*time.Second {
				backoff *= 2
			}
			continue
		}

		log.Infof("RWAWS - reconnect successful on attempt %d", i+1)
		return
	}

	log.Fatalf("RWAWS - failed to reconnect after %d attempts, giving up", maxRetries)
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
	log.Infof("RWAWS - subscribing to %d symbols: %s", len(symbols), string(b))

	if err := scraper.writeJSON(msg); err != nil {
		return fmt.Errorf("failed to subscribe to symbols: %v", err)
	}

	return nil
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

		if msg.Fails != nil {
			switch v := msg.Fails.(type) {
			case []interface{}:
				if len(v) > 0 {
					log.Errorf("RWAWS - subscribe partially failed, %d symbols failed: %v - will reconnect", len(v), v)
					scraper.handleDisconnect()
				}
			case map[string]interface{}:
				if len(v) > 0 {
					log.Errorf("RWAWS - subscribe partially failed, %d symbols failed: %v - will reconnect", len(v), v)
					scraper.handleDisconnect()
				}
			}
		}
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
	case contains(scraper.hkStocks, msg.Symbol):
		quote.Type = Equities

		if scraper.applyMarketStatus(msg, &quote) {
			return nil
		}

	case contains(scraper.usStocks, msg.Symbol):
		quote.Type = Equities

		if scraper.applyMarketStatus(msg, &quote) {
			return nil
		}

	case contains(scraper.usEtfs, msg.Symbol):
		quote.Type = ETF

		if scraper.applyMarketStatus(msg, &quote) {
			return nil
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

func (scraper *RWAWSScraper) applyMarketStatus(msg rwaWSMessage, quote *RWAWSQuote) (drop bool) {
	marketHoliday, marketOpen, knownMarket := scraper.getStockMarketStatus(msg)
	if knownMarket {
		quote.MarketHoliday = marketHoliday
		quote.MarketOpen = marketOpen
		if !marketOpen {
			return true
		}
	} else {
		quote.MarketHoliday = false
		quote.MarketOpen = true
	}
	return false
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
		scraper.triggerPublish()
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
		scraper.triggerPublish()

		scraper.flushTimerMu.Lock()
		scraper.flushTimer = nil
		scraper.flushTimerMu.Unlock()
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

		priceChanged := !exists || scraper.hasPriceDeviatedEnough(quote.Symbol, last, quote.Price)
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

func (scraper *RWAWSScraper) hasPriceDeviatedEnough(symbol string, oldPrice, newPrice float64) bool {
	threshold, hasThreshold := scraper.deviationThresholds[symbol]
	if !hasThreshold {
		// any price change triggers update
		return oldPrice != newPrice
	}
	if oldPrice == 0 {
		return true
	}
	deviation := math.Abs(newPrice-oldPrice) / oldPrice
	return deviation >= threshold
}

func (scraper *RWAWSScraper) preparePublishData(rwaResponse RWAWSQuote, marketStatusAdded *bool) (keys []string, values []*big.Float) {
	log.Info("got rwa ws data: ", rwaResponse)

	if (rwaResponse.Type == Equities || rwaResponse.Type == ETF) && !*marketStatusAdded {
		keys = append(keys, fmt.Sprintf("%s_Market_Open", rwaResponse.Symbol))
		marketOpen := new(big.Float).SetFloat64(0)
		if rwaResponse.MarketOpen {
			marketOpen = new(big.Float).SetFloat64(1)
		}
		values = append(values, marketOpen)

		keys = append(keys, fmt.Sprintf("%s_Market_Holiday", rwaResponse.Symbol))
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

func (scraper *RWAWSScraper) updateConfig(filename string) error {
	c, err := models.GetRWAWSConfig(filename)
	if err != nil {
		return err
	}

	oldSymbols := scraper.allSymbols()

	scraper.usStocks = c.US_Stocks
	scraper.hkStocks = c.HK_Stocks
	scraper.fxTickers = c.FX
	scraper.commodities = c.Commodities
	scraper.usEtfs = c.US_ETF
	scraper.deviationThresholds = c.DeviationThresholds

	log.Infof("RWAWS - loaded config: hkstocks=%d usstocks=%d fx=%d commodities=%d usetfs=%d",
		getNumSymbols(scraper.hkStocks),
		getNumSymbols(scraper.usStocks),
		getNumSymbols(scraper.fxTickers),
		getNumSymbols(scraper.commodities),
		getNumSymbols(scraper.usEtfs),
	)

	if len(oldSymbols) == 0 {
		return nil
	}

	newSymbols := scraper.allSymbols()
	if !symbolsEqual(oldSymbols, newSymbols) {
		log.Infof("RWAWS - symbols changed, subscribing to new symbols")
		if err := scraper.subscribeAll(); err != nil {
			log.Errorf("RWAWS - failed to subscribe to new symbols: %v", err)
		}
	}

	return nil
}

func symbolsEqual(oldSymbols []string, newSymbols []string) bool {
	if len(oldSymbols) != len(newSymbols) {
		return false
	}
	seen := make(map[string]struct{}, len(oldSymbols))
	for _, v := range oldSymbols {
		seen[v] = struct{}{}
	}
	for _, v := range newSymbols {
		if _, ok := seen[v]; !ok {
			return false
		}
	}
	return true
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

	appendUnique(scraper.hkStocks)
	appendUnique(scraper.usStocks)
	appendUnique(scraper.fxTickers)
	appendUnique(scraper.commodities)
	appendUnique(scraper.usEtfs)

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

	// NYSE / NASDAQ
	if mic == "XNYS" || mic == "XNAS" || mic == "BATS" || exchange == "NYSE" || exchange == "NASDAQ" || exchange == "CBOE" {
		holiday := scraper.isNYSEHoliday(now)
		open := scraper.isNYSEMarketOpen(now)
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

func (scraper *RWAWSScraper) isNYSEHoliday(now time.Time) bool {
	nyse := calendar.XNYS()
	return nyse.IsHoliday(now)
}

func (scraper *RWAWSScraper) isNYSEMarketOpen(now time.Time) bool {
	if scraper.isNYSEHoliday(now) {
		return false
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Errorf("load America/New_York location: %v", err)
		return false
	}
	now = now.In(loc)

	minutes := now.Hour()*60 + now.Minute()
	return minutes >= 9*60+30 && minutes < 16*60
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

func (scraper *RWAWSScraper) updateOracleMultiValuesForRWAWS(
	contract diaoraclev3.DIAOracleV3,
	auth *bind.TransactOpts,
	keys []string,
	values []*big.Float,
	decimals int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int

	multiplier := new(big.Float).SetInt(
		new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil),
	)

	submitTimestamp := time.Now().Unix()

	if submitTimestamp <= scraper.lastSubmittedTs {
		submitTimestamp = scraper.lastSubmittedTs + 1
	}
	scraper.lastSubmittedTs = submitTimestamp

	for i := range keys {
		scaled := new(big.Float).Mul(values[i], multiplier)
		scaledInt, _ := scaled.Int(nil)

		cValue := new(big.Int).Lsh(scaledInt, 128)
		cValue = cValue.Add(cValue, big.NewInt(submitTimestamp))
		cValues = append(cValues, cValue)
	}

	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	if err != nil {
		return err
	}

	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
	return nil
}
