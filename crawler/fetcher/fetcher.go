package fetcher

import (
	"fmt"
	"github.com/crawler/crawler_distributed/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter

	log.Printf("Fetching url %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf(
			"wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
