package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
		return
	}

	var gh GitHubContent
	if err = json.NewDecoder(resp.Body).Decode(&gh); err != nil {
		return
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
