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

type RWAWSConfig struct {
	HK_Stocks   []string `json:"HK_Stocks"`
	US_Stocks   []string `json:"US_Stocks"`
	FX          []string `json:"FX"`
	Commodities []string `json:"Commodities"`
	US_ETF      []string `json:"US_ETF"`
}

type BeloConfig struct {
	Pairs []string `json:"Pairs"`
}

func GetRWAConfig(filename string) (c RWAConfig, err error) {
	data, err := utils.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	return
}

func GetRWAWSConfig(filename string) (c RWAWSConfig, err error) {
	data, err := utils.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	return
}

func GetBeloConfig(filename string) (c BeloConfig, err error) {
	data, err := utils.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	return
}