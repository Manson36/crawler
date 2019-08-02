package engine

import (
	"github.com/crawler/restart/singletask/fetch"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetch.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher error: %s, %v", r.Url, err)
			continue
		}

		ParseResult := r.ParseFunc(body)
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got items %s", item)
		}
		}

}
