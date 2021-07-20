package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Crawler struct {
	URL      string
	Response []byte
}

func (c *Crawler) get() (err error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		log.Fatal("Error getting response. ", err)
		return err
	}
	defer resp.Body.Close()

	err = c.responseBody(resp)
	if err != nil {
		log.Fatal("Error getting response. ", err)
		return err
	}
	return nil
}

func (c *Crawler) responseBody(resp *http.Response) (err error) {
	c.Response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return err
	}

	return nil
}

func (c *Crawler) DataRetriever() (body []byte, err error) {
	err = c.get()
	if err != nil {
		return []byte{}, err
	}

	return c.Response, nil
}
