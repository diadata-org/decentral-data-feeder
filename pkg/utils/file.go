package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type GitHubContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

// readFile reads a gitcoin submission json file and returns the slice of items.
func readFromFile(filename string) (data []byte, err error) {
	var (
		jsonFile *os.File
	)
	path := os.Getenv("GOPATH") + "/src/github.com/diadata-org/decentral-data-feeder/config/rwa/" + filename
	jsonFile, err = os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	return io.ReadAll(jsonFile)

}

func readFromRemote(filename string) (data []byte, err error) {

	url := "https://api.github.com/repos/diadata-org/decentral-data-feeder/contents/config/rwa/" + filename

	req, _ := http.NewRequest("GET", url, nil)

	//Optional authentication
	githubToken := Getenv("GITHUB_TOKEN", "")
	if githubToken != "" {
		log.Info("Set github token for API requests")
		req.Header.Set("Authorization", "token "+githubToken)
	}

	time.Sleep(350 * time.Millisecond)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rateLimitRemaining, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	if err != nil {
		log.Error("rateLimitRemaining for github API calls: ", err)
	}
	log.Info("remaining github API calls: ", rateLimitRemaining)

	//Only applies for anonymous calls or exhausted PAT limits
	if rateLimitRemaining == 0 {
		rateLimitReset, errParseInt := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)
		if errParseInt != nil {
			log.Error("rateLimitReset for github API calls: ", errParseInt)
		}

		timeWait := rateLimitReset - time.Now().Unix()
		if timeWait < 0 {
			timeWait = 0
		}

		log.Warnf("Rate limit reached, waiting for refresh in %v", time.Duration(timeWait)*time.Second)
		time.Sleep(time.Duration(timeWait+30) * time.Second)

		resp.Body.Close()
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
		return nil, err
	}

	var gh GitHubContent
	if err = json.NewDecoder(resp.Body).Decode(&gh); err != nil {
		return nil, err
	}
	data, err = base64.StdEncoding.DecodeString(gh.Content)
	return

}

func ReadFile(filename string) ([]byte, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/root" || dir == "/home" {
		return readFromRemote(filename)
	}

	return readFromFile(filename)
}
