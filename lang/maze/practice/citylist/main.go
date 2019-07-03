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
			parseFunc: parseCity,
		})
	}
	return parseResult
}

func parseCity(contents []byte) parseResult {
	const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	parseResult := parseResult{}
	for _, m := range matches {
		parseResult.items = append(parseResult.items, "User " + string(m[2]))
		parseResult.requests = append(parseResult.requests, request{
			url: string(m[1]),
			parseFunc:nilParser,
		})
	}

	return parseResult
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

func Run(seed ...request) {
	var requests []request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, result.requests...)

		for _, item := range result.items {
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
