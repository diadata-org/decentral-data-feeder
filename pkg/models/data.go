package models

import "time"

type Data struct {
	Symbol    string
	Value     float64
	Source    string
	Timestamp time.Time
}
