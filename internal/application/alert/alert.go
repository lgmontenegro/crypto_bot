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
	oscilationBid := ticker.Bid - firstTicker.Bid
	oscilationAsk := ticker.Ask - firstTicker.Ask

	percBid := firstTicker.Bid * a.Perc
	percAsk := firstTicker.Ask * a.Perc

	direction := UP
	if oscilationBid < 0 {
		oscilationBid = oscilationBid * -1
		direction = DOWN
	}

	if oscilationAsk < 0 {
		oscilationAsk = oscilationAsk * -1
		direction = DOWN
	}

	if (oscilationBid >= percBid) || (oscilationAsk >= percAsk) {
		var oscilation float64
		switch {
			case oscilationBid >= percBid:
				oscilation = oscilationBid
			case oscilationAsk >= percAsk:
				oscilation = oscilationAsk
			default:
				oscilation = 0.0
		}
		
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
