package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

const (
	RANDAMU              = "Randamu"
	RANDAMU_API_BASE_URL = "https://vrf-api.dcipher.network/v1/dia/"
)

type RandamuScraper struct {
	apiKey            string
	requestChannel    chan []byte
	dataChannel       chan []byte
	updateDoneChannel chan bool
}

type RandamuMetadata struct {
	Round     int64  `json:"round"`
	Seed      string `json:"seed"`
	Signature string `json:"signature"`
}

type RandamuRequest struct {
	Type int
	Body []byte
}

type RandamuResponse struct {
	Type     int
	Response []byte
}

type RandamuBytesRequest struct {
	// request parameters for array of random bytes
	RequestID string
	Seed      string
	LenBytes  uint8
	NumBytes  uint8
}

type RandamuBytesResponse struct {
	RequestID  string          `json:"request_id"`
	Metadata   RandamuMetadata `json:"metadata"`
	Randomness []string        `json:"randomness"`
}

type RandamuIntRangeRequest struct {
	// request parameters for a range of ints with limits
	RequestID string
	Seed      string
	Min       uint8
	Max       uint8
	NumInts   uint8
}

type RandamuIntRangeResponse struct {
	RequestID string          `json:"request_id"`
	Metadata  RandamuMetadata `json:"metadata"`
	Ints      []string        `json:"randomness"`
}

type RandamuIntRequest struct {
	// request parameters for array of random ints
	RequestID string
	Seed      string
	BitSize   uint8
	NumInts   uint8
}

type RandamuIntResponse struct {
	RequestID string          `json:"request_id"`
	Metadata  RandamuMetadata `json:"metadata"`
	Ints      []string        `json:"randomness"`
}

func NewRandamuScraper() *RandamuScraper {

	var scraper RandamuScraper
	scraper.apiKey = utils.Getenv("RANDAMU_API_KEY", "")
	scraper.dataChannel = make(chan []byte)
	scraper.requestChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	go scraper.listen()

	// // Test IntRequest ------------------------------------------------------
	var requestInt RandamuIntRequest
	requestInt.BitSize = uint8(8)
	requestInt.NumInts = uint8(10)
	requestInt.RequestID = "testID"
	requestBody, err := json.Marshal(requestInt)
	if err != nil {
		log.Error("marshal requestInt: ", err)
	}
	var request RandamuRequest
	request.Body = requestBody
	request.Type = 2
	requestByte, err := json.Marshal(request)
	if err != nil {
		log.Error("marshal request: ", err)
	}

	scraper.requestChannel <- requestByte
	// // Test ----------------------------------------------------------------

	// // Test ByteRequest ------------------------------------------------------
	// var requestByte RandamuBytesRequest
	// requestByte.LenBytes = uint8(10)
	// requestByte.NumBytes = uint8(4)
	// requestByte.RequestID = "testID"
	// requestBody, err := json.Marshal(requestByte)
	// if err != nil {
	// 	log.Error("marshal requestInt: ", err)
	// }
	// var request RandamuRequest
	// request.Body = requestBody
	// request.Type = 0
	// requestMarshalled, err := json.Marshal(request)
	// if err != nil {
	// 	log.Error("marshal request: ", err)
	// }

	// scraper.requestChannel <- requestMarshalled
	// // Test ----------------------------------------------------------------

	// Test IntRangeRequest ------------------------------------------------------
	// testTicker := time.NewTicker(time.Duration(10 * time.Second))
	// go func() {
	// 	for range testTicker.C {
	// 		var requestIntRange RandamuIntRangeRequest
	// 		requestIntRange.Min = uint8(2)
	// 		requestIntRange.Max = uint8(10)
	// 		requestIntRange.NumInts = uint8(4)
	// 		requestIntRange.RequestID = "testID"
	// 		requestBody, err := json.Marshal(requestIntRange)
	// 		if err != nil {
	// 			log.Error("marshal requestInt: ", err)
	// 		}
	// 		var request RandamuRequest
	// 		request.Body = requestBody
	// 		request.Type = 1
	// 		requestMarshalled, err := json.Marshal(request)
	// 		if err != nil {
	// 			log.Error("marshal request: ", err)
	// 		}

	// 		scraper.requestChannel <- requestMarshalled
	// 	}
	// }()
	// Test ----------------------------------------------------------------

	return &scraper
}

func (scraper *RandamuScraper) listen() {

	// TO DO: Listen to requests from requestor.
	log.Info("listen")

	// For now, assume @response is a response from listening to the requestor.
	for request := range scraper.requestChannel {
		response, err := scraper.requestData(request)
		if err != nil {
			log.Error("requestData: ", err)
		} else {
			scraper.DataChannel() <- response
		}
	}

}

func (scraper *RandamuScraper) requestData(requestBytes []byte) ([]byte, error) {
	var request RandamuRequest
	err := json.Unmarshal(requestBytes, &request)
	if err != nil {
		return []byte{}, err
	}
	switch request.Type {
	case 0:
		var randomBytesRequest RandamuBytesRequest
		err = json.Unmarshal(request.Body, &randomBytesRequest)
		if err != nil {
			return []byte{}, err
		}
		randomBytes, err := scraper.getRandomBytes(randomBytesRequest)
		if err != nil {
			return randomBytes, err
		}

		// make response struct
		var resp RandamuResponse
		resp.Type = request.Type
		resp.Response = randomBytes

		return json.Marshal(resp)

	case 1:
		var intRangeRequest RandamuIntRangeRequest
		err = json.Unmarshal(request.Body, &intRangeRequest)
		if err != nil {
			return []byte{}, err
		}
		randomIntRange, err := scraper.getRandomIntRange(intRangeRequest)
		if err != nil {
			return randomIntRange, err
		}

		// make response struct
		var resp RandamuResponse
		resp.Type = request.Type
		resp.Response = randomIntRange

		return json.Marshal(resp)

	case 2:
		var intRequest RandamuIntRequest
		err = json.Unmarshal(request.Body, &intRequest)
		if err != nil {
			log.Error("Unmarshal intRequest body: ", err)
		}
		randomInts, err := scraper.getRandomInts(intRequest)
		if err != nil {
			return randomInts, err
		}

		// make response struct
		var resp RandamuResponse
		resp.Type = request.Type
		resp.Response = randomInts

		return json.Marshal(resp)
	}
	return []byte{}, errors.New("unknown request type")
}

func (scraper *RandamuScraper) getRandomBytes(requestData RandamuBytesRequest) ([]byte, error) {
	log.Info("getRandomBytes")

	url := RANDAMU_API_BASE_URL + "random/bytes/" + strconv.Itoa(int(requestData.LenBytes))
	var body []byte
	if requestData.Seed != "" {
		body = []byte(fmt.Sprintf(`{"seed":%s}`, requestData.Seed))
	}
	return scraper.queryRandamuAPI(url, body)

}

func (scraper *RandamuScraper) getRandomIntRange(requestData RandamuIntRangeRequest) ([]byte, error) {
	log.Info("getRandomIntRange")

	url := RANDAMU_API_BASE_URL + "random/in-range"
	var body []byte
	min := strconv.Itoa(int(requestData.Min))
	max := strconv.Itoa(int(requestData.Max))
	if requestData.Seed == "" {
		body = []byte(fmt.Sprintf(`{"min":"%s","max":"%s","n":%v}`, min, max, requestData.NumInts))
	} else {
		body = []byte(fmt.Sprintf(`{"seed":%s,"min":"%s","max":"%s","n":%v}`, requestData.Seed, min, max, requestData.NumInts))
	}
	return scraper.queryRandamuAPI(url, body)

}

func (scraper *RandamuScraper) getRandomInts(requestData RandamuIntRequest) ([]byte, error) {
	log.Info("getRandomInts")

	url := RANDAMU_API_BASE_URL + "random/int/" + strconv.Itoa(int(requestData.BitSize))
	var body []byte
	if requestData.Seed == "" {
		body = []byte(fmt.Sprintf(`{"type":"int","n":%v}`, requestData.NumInts))
	} else {
		body = []byte(fmt.Sprintf(`{"type":"int","n":%v,"seed":"%s"}`, requestData.NumInts, requestData.Seed))
	}

	return scraper.queryRandamuAPI(url, body)

}

// queryRandamuAPI queries Randamu API using a POST query. It returns the raw response body.
func (scraper *RandamuScraper) queryRandamuAPI(url string, requestBody []byte) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return []byte{}, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+scraper.apiKey)

	log.Info("request: ", request)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func (scraper *RandamuScraper) Close() error {
	log.Info("TO DO.")
	return nil
}

func (scraper *RandamuScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}

func (scraper *RandamuScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
