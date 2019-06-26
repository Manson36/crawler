package engine

import (
	"github.com/crawler/crawler/fetcher"
	"log"
)

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher err: fetching %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}
