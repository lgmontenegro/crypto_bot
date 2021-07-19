package config

import (
	"time"
)

type Config struct {
	Times    time.Duration
	URL      string
	Endpoint string
	Pairs    []string
	Alerts   []Alert
}

type Alert struct {
	Pair string
	Perc float64
}
