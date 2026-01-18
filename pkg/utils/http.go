package utils

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetRequestAuth(url string, accessToken string) (data []byte, statusCode int, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode

	if statusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("Failed to get issuer tokens: %s %s", resp.Status, string(body))
		return
	}
	data, err = io.ReadAll(resp.Body)
	return
}

// GetRequest performs a get request on @url and returns the response body
// as a slice of byte data.
func GetRequest(url string) ([]byte, int, error) {

	response, err := http.Get(url) //nolint:noctx,gosec
	if err != nil {
		log.Error("get request: ", err)
		return []byte{}, 0, err
	}

	// Close response body after function
	defer func() {
		cerr := response.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		return []byte{}, response.StatusCode, fmt.Errorf("HTTP Response Error %d", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	return XMLdata, response.StatusCode, err
}
