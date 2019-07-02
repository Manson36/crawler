package engine

import (
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Get item 	%v", item)
		}
	}
}
