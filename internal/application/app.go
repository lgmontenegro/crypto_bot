package application

import (
	"fmt"
	"lgmontenegro/crypto_bot/internal/application/alert"
	"lgmontenegro/crypto_bot/internal/application/crawler"
	"lgmontenegro/crypto_bot/internal/application/data_processor"
	"lgmontenegro/crypto_bot/internal/config"
	"runtime"
	"sync"
	"time"
)

type pairCrawler struct {
	crawler *crawler.Crawler
	pair    string
}

type Application struct {
	pairsCrawlers []pairCrawler
	dataProcessor data_processor.DataProcessor
}

func (a *Application) Bootstrap(config config.Config) {
	for _, pair := range config.Pairs {
		c := crawler.Crawler{}
		c.URL = fmt.Sprintf("%s%s%s", config.URL, config.Endpoint, pair)
		pairSet := pairCrawler{}
		pairSet.crawler = &c
		pairSet.pair = pair

		a.pairsCrawlers = append(a.pairsCrawlers, pairSet)		
	}

	alertsCfg := alert.Alerts{}
	for _, alertCfg := range config.Alerts {
		alertConfig := alert.Alert{
			Pair: alertCfg.Pair,
			Perc: alertCfg.Perc,
		}
		alertsCfg = append(alertsCfg, alertConfig)
	}

	a.dataProcessor = data_processor.DataProcessor{
		Times:  config.Times * time.Second,
		Alerts: alertsCfg,
	}
}

func (a *Application) Start() {
	runtime.GOMAXPROCS(len(a.pairsCrawlers))
	var wg sync.WaitGroup
	wg.Add(len(a.pairsCrawlers))
	for _, pairSet := range a.pairsCrawlers {
		go a.dataProcessor.Process(pairSet.crawler, pairSet.pair, wg)
	}
	wg.Wait()
}
