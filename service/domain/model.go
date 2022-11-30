package domain

import "time"

type PriceNote struct {
	TimeStamp time.Time
	Currency  string  // crypty currency
	Price     float64 // in USD
}
