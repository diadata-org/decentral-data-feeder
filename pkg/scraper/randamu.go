package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	requestor "github.com/diadata-org/decentral-data-feeder/contracts/VRFRequestor"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	requestorContract string
	restClient        *ethclient.Client
	wsClient          *ethclient.Client
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

// Type 0
type RandamuBytesRequest struct {
	// request parameters for array of random bytes
	RequestID *big.Int
	Seed      string
	LenBytes  uint8
	NumBytes  uint8
}

type RandamuBytesResponse struct {
	RequestID  *big.Int        `json:"request_id"`
	Metadata   RandamuMetadata `json:"metadata"`
	Randomness []string        `json:"randomness"`
}

// Type 1
type RandamuIntRangeRequest struct {
	// request parameters for a range of ints with limits
	RequestID *big.Int
	Seed      string
	Min       uint8
	Max       uint8
	NumInts   uint8
}

type RandamuIntRangeResponse struct {
	RequestID *big.Int        `json:"request_id"`
	Metadata  RandamuMetadata `json:"metadata"`
	Ints      []string        `json:"randomness"`
}

// Type 2
type RandamuIntRequest struct {
	// request parameters for array of random ints
	RequestID *big.Int
	Seed      string
	BitSize   uint8
	NumInts   uint8
}

type RandamuIntResponse struct {
	RequestID *big.Int        `json:"request_id"`
	Metadata  RandamuMetadata `json:"metadata"`
	Ints      []string        `json:"randomness"`
}

func NewRandamuScraper() *RandamuScraper {

	var scraper RandamuScraper
	scraper.apiKey = utils.Getenv("RANDAMU_API_KEY", "")
	scraper.dataChannel = make(chan []byte)
	scraper.requestChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)
	scraper.requestorContract = utils.Getenv("REQUESTOR_CONTRACT", "0x3E02966F02b51f395F296E308A301076CfEd0930")
	restClient, err := ethclient.Dial(utils.Getenv("LUMINA_REST_CLIENT", ""))
	if err != nil {
		log.Fatal("get lumina rest client: ", err)
	}
	wsClient, err := ethclient.Dial(utils.Getenv("LUMINA_WS_CLIENT", ""))
	if err != nil {
		log.Fatal("get lumina ws client: ", err)
	}
	scraper.restClient = restClient
	scraper.wsClient = wsClient

	go scraper.listen()

	// // // Test IntRequest ------------------------------------------------------
	// var requestInt RandamuIntRequest
	// requestInt.BitSize = uint8(8)
	// requestInt.NumInts = uint8(10)
	// requestInt.RequestID = "testID"
	// requestBody, err := json.Marshal(requestInt)
	// if err != nil {
	// 	log.Error("marshal requestInt: ", err)
	// }
	// var request RandamuRequest
	// request.Body = requestBody
	// request.Type = 2
	// requestByte, err := json.Marshal(request)
	// if err != nil {
	// 	log.Error("marshal request: ", err)
	// }

	// scraper.requestChannel <- requestByte
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
	requestFilterer, err := requestor.NewVRFRequestorFilterer(common.HexToAddress(scraper.requestorContract), scraper.wsClient)
	if err != nil {
		log.Error("NewVRFRequestorFilterer: ", err)
		return
	}

	requestSink := make(chan *requestor.VRFRequestorRequestReceived)
	sub, err := requestFilterer.WatchRequestReceived(&bind.WatchOpts{}, requestSink, nil)
	if err != nil {
		log.Error("WatchRequestReceived: ", err)
	}

	go func() {
		for {
			select {
			case r, ok := <-requestSink:
				if !ok {
					log.Warn("Event channel closed")
					return
				}
				log.Infof("got request: numValues -- requestId -- seed: %v -- %s -- %s ", r.NumOfValues.Int64(), r.RequestId.String(), r.Seed)
				response, err := scraper.processRequestEvent(r)
				if err != nil {
					log.Error("requestData: ", err)
				} else {
					scraper.DataChannel() <- response
				}
			case err := <-sub.Err():
				log.Error("Subscription error: ", err)
				return
			}
		}
	}()

}

func (scraper *RandamuScraper) processRequestEvent(requestEvent *requestor.VRFRequestorRequestReceived) ([]byte, error) {
	var randamuRequest RandamuRequest

	// For the MVP we're currently restricted to int arrays.
	randamuRequest.Type = 2

	// Fill request types and marshal them.
	var requestInt RandamuIntRequest
	requestInt.BitSize = 64
	requestInt.NumInts = uint8(requestEvent.NumOfValues.Int64())
	requestInt.RequestID = requestEvent.RequestId
	requestInt.Seed = requestEvent.Seed
	requestBody, err := json.Marshal(requestInt)
	if err != nil {
		log.Error("marshal requestInt: ", err)
	}
	randamuRequest.Body = requestBody
	requestByte, err := json.Marshal(randamuRequest)
	if err != nil {
		log.Error("marshal request: ", err)
	}

	// make randamu request using marshalled request data and send the obtained response to main.
	return scraper.requestData(requestByte)

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
