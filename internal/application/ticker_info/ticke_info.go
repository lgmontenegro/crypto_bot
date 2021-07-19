package ticker_info

import "time"

type Ticker struct {
	Ask       float64   `json:"ask,string"`
	Bid       float64   `json:"bid,string"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"-"`
}
