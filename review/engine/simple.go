package engine

import "log"

type SimpleEngine struct {

}

func (e *SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParserResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, ParserResult.Requests...)

		for _, item := range ParserResult.Items {
			log.Printf("Get item %v", item)
		}
	}
}