package scraper

import (
	"fmt"
	"strings"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	TWELVEDATA = "TwelveData"
	PARTICULA  = "Particula"
	BELO       = "Belo"
	RWAWS      = "RWAWS"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func GetSourcesFromEnv(envName string) ([]string, error) {
	srcStr := utils.Getenv(envName, "")
	srcNames := strings.Split(srcStr, ",")
	var sources []string
	for _, name := range srcNames {
		src, ok := sourceFromString(name) // Implement this conversion
		if !ok {
			return []string{}, fmt.Errorf("source %s does not exist", name)
		}
		sources = append(sources, src)
	}
	return sources, nil
}

func sourceFromString(s string) (string, bool) {
	switch s {
	case TWELVEDATA:
		return TWELVEDATA, true
	case PARTICULA:
		return PARTICULA, true
	case BELO:
		return BELO, true
	case RWAWS:
		return RWAWS, true
	}
	return "", false
}
