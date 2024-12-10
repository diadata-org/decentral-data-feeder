package models

// Asset is a generic data type for assets, ranging from fiat to crypto.
type Asset struct {
	Symbol     string `json:"Symbol"`
	Name       string `json:"Name"`
	Address    string `json:"Address"`
	Decimals   uint8  `json:"Decimals"`
	Blockchain string `json:"Blockchain"`
}

func (a *Asset) AssetIdentifier() string {
	return a.Blockchain + "-" + a.Address
}
