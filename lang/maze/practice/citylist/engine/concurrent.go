package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	url string
	parseFunc func([]byte) parseResult
}

type parseResult struct {
	items []interface{}
	requests []request
}

func worker (r request) (parseResult, error) {
	log.Printf("Fetching url", r.url)
	body, err := fetch(r.url)
	if err != nil {
		log.Printf("fetching err: url %s: err %v", r.url, err)
		return parseResult{}, err
	}

	return r.parseFunc(body), nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error statuscode", resp.StatusCode)
		return nil, errors.New("wrong status code")
	}

	return ioutil.ReadAll(resp.Body)
}


type concurrentEngine struct {
	scheduler scheduler
	workerCount int
}

type scheduler interface {
	submit(request)
}

func (e concurrentEngine) Run2(seed ...request) {
	in := make(chan request)
	out := make(chan parseResult)

	for i := 0; i < e.workerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seed {
		e.scheduler.submit(r)
	}

	for {
		result := <- out
		for _, item := range result.items {
			log.Printf("Got item :%v", item)
		}

		for _, request := range result.requests {
			e.scheduler.submit(request)
		}
	}
}

func createWorker(in chan request, out chan parseResult) {
	go func() {
		for {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
