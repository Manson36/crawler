package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type request struct {
	url string
	parseFunc func([]byte) parseResult
}

type parseResult struct {
	items []interface{}
	requests []request
}

func nilParser([]byte) parseResult {
	return parseResult{}
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

func parseCityList(contents []byte) parseResult {
	const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	parseResult := parseResult{}
	for _, m := range matches {
		parseResult.items = append(parseResult.items, string(m[2]))
		parseResult.requests = append(parseResult.requests, request{
			url: string(m[1]),
			parseFunc: nilParser,
		})
	}
	return parseResult
}

func Run(seed ...request) {
	var requests []request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching url", r.url)
		body, err := fetch(r.url)
		if err != nil {
			log.Printf("fetching err: url %s: err %v", r.url, err)
			continue
		}

		parseResult := r.parseFunc(body)
		requests = append(requests, parseResult.requests...)

		for _, item := range parseResult.items {
			log.Printf("Got item %v", item)
		}
	}
}

func main() {
	Run(request{
		url: "http://www.zhenai.com/zhenghun",
		parseFunc:parseCityList,
	})
}
