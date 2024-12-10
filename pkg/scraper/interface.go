package scraper

import (
	"time"

	models "github.com/diadata-org/decentral-data-feeder/pkg/models"
)

type DataScraper interface {
	// DataChannel enables a scraper to transport scraped data.
	DataChannel() chan models.Data
	Close() error
}

func NewDataScraper(source string, triggerChannel *chan time.Time) DataScraper {
	switch source {
	case COINMARKETCAP:
		return NewCMCScraper(triggerChannel)
	case COINGECKO:
		return NewCGScraper(triggerChannel)
	default:
		return nil
	}
}
