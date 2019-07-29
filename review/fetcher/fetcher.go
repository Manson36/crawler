package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second/20)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter

	log.Printf("Fetching url %s", url)
	 resp, err := http.Get(url)
	 if err != nil {
	 	return nil, err
	 }

	 defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
	 	fmt.Println("Error: resp code")
	 	return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	 }

	 return ioutil.ReadAll(resp.Body)
}
