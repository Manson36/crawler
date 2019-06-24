package main

import (
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
	requests []request
	items []interface{}
}

func fetcher(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Fetcher status code error :%v", err)
		return nil, fmt.Errorf("fetcher status code err :%v", err)
	}

	return ioutil.ReadAll(resp.Body)
}

func nilParser([]byte) parseResult {
	return parseResult{}
}

func parseCityList(contents []byte) parseResult {
	const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := parseResult{}
	for _, r := range matches {
		result.items = append(result.items, string(r[2]))
		result.requests = append(result.requests, request{
			url: string(r[1]),
			parseFunc:nilParser,
		})
	}

	return result
}

func run(seed ...request) {
	var requests []request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching url %s", r.url)
		 body, err := fetcher(r.url)
		 if err != nil{
		 	log.Printf("fetching err: url%s : %v", r.url, err)
		 	continue
		 }

		parseResult := parseCityList(body)
		requests = append(requests, parseResult.requests...)

		for _, item := range parseResult.items {
			log.Printf("got item %v", item)
		}
	}
}

func main() {
	run(request{
		url:"http://www.zhenai.com/zhenghun",
		parseFunc:nilParser,
	})
}
