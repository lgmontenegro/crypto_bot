package data_processor

import (
	"encoding/json"
	"fmt"
	"lgmontenegro/crypto_bot/internal/application/alert"
	"lgmontenegro/crypto_bot/internal/application/ticker_info"
	"log"
	"sync"
	"time"
)

type DataProcessor struct {
	Times  time.Duration
	Alerts alert.Alerts
}

type Data interface {
	DataRetriever() (body []byte, err error)
}

func (d *DataProcessor) Process(exec Data, pair string, wg sync.WaitGroup) (err error) {
	defer wg.Done()

	var setAlarm bool
	firstTicker := ticker_info.Ticker{}
	firstTickerLoop := true

	for t := range time.Tick(d.Times) {
		body, err := exec.DataRetriever()
		if err != nil {
			log.Fatal(err)
			return err
		}

		newTicker := ticker_info.Ticker{}
		err = json.Unmarshal(body, &newTicker)
		if err != nil {
			log.Fatal(err)
			return err
		}
		newTicker.CreatedAt = t

		if firstTickerLoop {
			firstTicker = newTicker
			firstTickerLoop = false
		}

		for _, alarm := range d.Alerts {
			if alarm.Pair == pair {
				setAlarm = alarm.Watch(firstTicker, newTicker)
			}

			if setAlarm {
				fmt.Println(alarm.Warning.Ticker.CreatedAt.Date())
			}
		}

		fmt.Println(string(body))
	}

	return nil
}
