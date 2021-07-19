package alert

import (
	ticker "lgmontenegro/crypto_bot/internal/application/ticker_info"
	"time"
)

const (
	UP   = "up"
	DOWN = "down"
)

type Alerts []Alert

type Alert struct {
	Pair    string
	Perc    float64
	Times   time.Duration
	Warning Alarm
}

type Alarm struct {
	Ticker     ticker.Ticker
	Oscilation float64
	CreatedAt  time.Time
	Direction  string
}

func (a *Alert) Watch(firstTicker ticker.Ticker, ticker ticker.Ticker) (alarmTicker bool) {
	oscilation := ticker.Bid - firstTicker.Bid

	direction := UP
	if oscilation < 0 {
		direction = DOWN
		oscilation = oscilation * -1.0
	}

	if oscilation >= a.Perc {
		a.Warning = Alarm{
			Ticker:     ticker,
			Oscilation: oscilation,
			CreatedAt:  time.Now(),
			Direction:  direction,
		}
		return true
	}

	return false
}
