package scraper

type DataScraper interface {
	// DataChannel enables a scraper to transport scraped data.
	DataChannel() chan []byte
	UpdateDoneChannel() chan bool
	Close() error
}

func NewDataScraper(source string) DataScraper {
	switch source {
	case COINMARKETCAP:
		return NewCMCScraper()
	case COINGECKO:
		return NewCGScraper()
	case RANDAMU:
		return NewRandamuScraper()
	case TWELVEDATA:
		return NewTwelvedataScraper()
	case PARTICULA:
		return NewParticulaScraper()
	default:
		return nil
	}
}
