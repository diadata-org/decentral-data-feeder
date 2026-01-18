package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

// API Configuration
const (
	PARTICULA           = "Particula"
	PARTICULA_BASE_URL  = "https://api.particula.io"
	PARTICULA_AUTH_URL  = PARTICULA_BASE_URL + "/oauth/token"
	PARTICULA_PRECISION = 1e8
)

var (
	particulaUpdateSeconds int64
)

type particulaAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type particulaIssuer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Image  string `json:"image"`
}

type particulaIssuersResponse []particulaIssuer

type particulaTokenRatingResponse struct {
	TokenName                string  `json:"token_name"`
	IssuerName               string  `json:"issuer_name"`
	UnderlyingName           string  `json:"underlying_name"`
	FinalRating              string  `json:"final_rating"`
	FirstRatedAt             string  `json:"first_rated_at"`
	LastRatedAt              string  `json:"last_rated_at"`
	EnvironmentalImpactScore float64 `json:"environmental_impact_score"`
	ComplianceScore          float64 `json:"compliance_score"`
	TechnologyScore          float64 `json:"technology_score"`
	EconomicsScore           float64 `json:"economics_score"`
	TransparencyScore        float64 `json:"transparency_score"`
}

type ParticulaScraper struct {
	dataChannel       chan []byte
	updateDoneChannel chan bool
	ticker            *time.Ticker
	stockSymbols      []string
	stockMarketOpen   bool
	apiClientID       string
	apiKey            string
}

func init() {
	var err error
	particulaUpdateSeconds, err = strconv.ParseInt(utils.Getenv("UPDATE_SECONDS", "30"), 10, 64)
	if err != nil {
		log.Error("Parse UPDATE_SECONDS: ", err)
	}
}

func NewParticulaScraper() *ParticulaScraper {

	scraper := &ParticulaScraper{
		ticker:      time.NewTicker(time.Duration(particulaUpdateSeconds) * time.Second),
		apiClientID: utils.Getenv("PARTICULA_API_CLIENT_ID", ""),
		apiKey:      utils.Getenv("PARTICULA_API_KEY", ""),
	}

	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	go scraper.mainLoop()

	return scraper

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *ParticulaScraper) mainLoop() {

	// Initial run.
	err := scraper.UpdateRatings()
	if err != nil {
		log.Error(err)
	}
	scraper.updateDoneChannel <- true

	// Update every @particulaUpdateSeconds.
	for range scraper.ticker.C {
		err := scraper.UpdateRatings()
		if err != nil {
			log.Error(err)
			continue
		}
		scraper.updateDoneChannel <- true
	}

}

func (scraper *ParticulaScraper) UpdateRatings() error {
	log.Info("Authenticating...")

	accessToken, err := getAccessToken(scraper.apiClientID, scraper.apiKey)
	if err != nil || accessToken == "" {
		fmt.Println("Failed to get access token:", err)
		return err
	}

	log.Info("Authentication successful!")
	log.Info("Fetching issuers...")

	issuerResponse, err := listIssuers(accessToken)
	if err != nil || len(issuerResponse) == 0 {
		log.Info("Failed to retrieve issuers or no issuers found:", err)
		return err
	}

	log.Infof("Total issuers found: %d", len(issuerResponse))

	for range issuerResponse {
		for k, v := range issuerResponse {
			log.Infof("issuer %v: %v", k+1, v)
		}
	}

	// Iterate over issuers
	for _, issuer := range issuerResponse {

		log.Infof("Fetching tokens for issuer  with ID %s", issuer.ID)

		tokens, err := getIssuerToken(issuer.ID, accessToken)
		if err != nil || len(tokens) == 0 {
			return fmt.Errorf("Failed to retrieve issuer tokens or no tokens found: %v", err)
		}

		log.Infof("Total tokens found: %d", len(tokens))

		for i, token := range tokens {
			log.Infof("token %d :: name -- symbol -- image -- id: %s -- %s -- %s -- %s", i, token.Name, token.Symbol, token.Image, token.ID)

			rating, err := getTokenRatings(token.ID, accessToken)
			if err != nil {
				return err
			}

			m, err := makeParticulaRatingMap(token, rating, PARTICULA_PRECISION)
			if err != nil {
				return err
			}

			ratingBytes, err := json.Marshal(m)
			if err != nil {
				log.Error("marshal rating data: ", err)
			}

			scraper.dataChannel <- ratingBytes

		}
	}

	return nil
}

func (scraper *ParticulaScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}

func (scraper *ParticulaScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}

func (scraper *ParticulaScraper) Close() error {
	return nil
}

func getAccessToken(clientID, clientSecret string) (string, error) {
	data := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}
	req, err := http.NewRequest("POST", PARTICULA_AUTH_URL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("authentication failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Auth failed: %s\n%s", resp.Status, string(body))
	}

	var authData particulaAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authData); err != nil {
		return "", err
	}
	return authData.AccessToken, nil
}

func listIssuers(accessToken string) (issuers particulaIssuersResponse, err error) {
	issuersURL := PARTICULA_BASE_URL + "/v1/issuer"
	data, _, err := utils.GetRequestAuth(issuersURL, accessToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &issuers)
	return
}

func getIssuerToken(issuerID, accessToken string) (issuerTokens particulaIssuersResponse, err error) {
	issuerTokenURL := fmt.Sprintf("%s/v1/issuer/%s/token", PARTICULA_BASE_URL, issuerID)
	data, _, err := utils.GetRequestAuth(issuerTokenURL, accessToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &issuerTokens)
	return

}

func getTokenRatings(tokenID, accessToken string) (rating particulaTokenRatingResponse, err error) {
	ratingsURL := fmt.Sprintf("%s/v1/token/%s/rating", PARTICULA_BASE_URL, tokenID)
	data, _, err := utils.GetRequestAuth(ratingsURL, accessToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &rating)
	return
}

// makeParticulaRatingMap returns a map of rating quantities associated with @token that can be used
// in a key-value store of types string-int64. Floats have precision @precision.
func makeParticulaRatingMap(token particulaIssuer, tokenRating particulaTokenRatingResponse, precision float64) (m map[string]int64, err error) {
	m = make(map[string]int64)

	m[PARTICULA+":"+token.Symbol+":FinalRating"] = encodeRating(tokenRating.FinalRating)

	firstRatedAt, err := time.Parse("2006-01-02 15:04:05", tokenRating.FirstRatedAt)
	if err != nil {
		return
	}
	lastRatedAt, err := time.Parse("2006-01-02 15:04:05", tokenRating.LastRatedAt)
	if err != nil {
		return
	}

	m[PARTICULA+":"+token.Symbol+":FirstRatedAt"] = firstRatedAt.Unix()
	m[PARTICULA+":"+token.Symbol+":LastRatedAt"] = lastRatedAt.Unix()
	m[PARTICULA+":"+token.Symbol+":EnvironmentalImpactScore"] = int64(math.Round(tokenRating.EnvironmentalImpactScore * precision))
	m[PARTICULA+":"+token.Symbol+":ComplianceScore"] = int64(math.Round(tokenRating.ComplianceScore * precision))
	m[PARTICULA+":"+token.Symbol+":TechnologyScore"] = int64(tokenRating.TechnologyScore * precision)
	m[PARTICULA+":"+token.Symbol+":EconomicsScore"] = int64(math.Round(tokenRating.EconomicsScore * precision))
	m[PARTICULA+":"+token.Symbol+":TransparencyScore"] = int64(math.Round(tokenRating.TransparencyScore * precision))

	return

}

// encodeRating encodes a rating given by a string to an integer.
// An encoding always consists of 4 digits.
// First 3 digits are the rating. Each letter of the rating is encoded to its position in the alphabet.
// 4th digit is encoded as -:1 , +:2 , :0
func encodeRating(rating string) (intRating int64) {
	for i, l := range rating {
		intRating += int64(1000 / pow10(i) * letterToPosition(l))
	}
	letters := strings.Split(rating, "")
	if letters[len(letters)-1] == "-" {
		intRating += 1
	}
	if letters[len(letters)-1] == "+" {
		intRating += 2
	}
	return
}

func letterToPosition(c rune) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	return 0 // Not a letter
}

func pow10(n int) int {
	result := 1
	for range n {
		result *= 10
	}
	return result
}
