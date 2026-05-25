package models

import (
	"encoding/json"

	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

type RWAConfig struct {
	Stocks      []string `json:"Stocks"`
	FX          []string `json:"FX"`
	Commodities []string `json:"Commodities"`
	ETF         []string `json:"ETF"`
}

type BeloConfig struct {
	Pairs []string `json:"Pairs"`
}

func GetRWAConfig(filename string, branch string) (c RWAConfig, err error) {

	data, err := utils.ReadFile(filename, branch)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &c)
	return

}

func GetBeloConfig(filename string, branch string) (c BeloConfig, err error) {

	data, err := utils.ReadFile(filename, branch)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &c)
	return

}
