package parser

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler_distributed/config"
	"regexp"
)

var (
	ProfileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParserResult {//定义的ParserFunc（）中添加了string，
																// 这里也需要添加，不使用，_
	matches := ProfileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	result = engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser:engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})
	}

	return result
}