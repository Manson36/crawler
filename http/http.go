package main

import (
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
}
