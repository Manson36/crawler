package serialized

import (
	"github.com/crawler/crawler/engine"
	parser2 "github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
	"log"
	"net"
)

type parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{} )
}

type ParseFunc func (contents []byte, url string) ParserResult

type FuncParser struct {
	parser ParseFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) interface{} {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name: name,
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) interface{} {
	panic("implement me")
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	panic("implement me")
}

func NewProfileParser(userName string) *ProfileParser {
	return &ProfileParser{userName:userName}
}

//序列化和反序列化
type serializedParser struct {
	Name string
	Args interface{}
}

type request struct {
	url string
	Parser serializedParser
}

type parseresult struct {
	requests []request
	items []engine.Item
}

func serializeRequest(r engine.Request) request {
	name, args := r.Parser.Serialize()
	return request{
		url: r.Url,
		Parser:serializedParser{
			Name:name,
			Args:args,
		},
	}
}

func serializeResult(r engine.ParserResult) parseresult {
	result :=  parseresult{
		items:r.Items,
	}

	for _, req := range r.Requests {
		result.requests = append(result.requests, serializeRequest(req))
	}

	return result
}

func deserializeRequest(r request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)

	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:r.url,
		Parser:parser,
	}, nil
}

func deserializeResult(r parseresult) engine.ParserResult {
	result := engine.ParserResult{
		Items:r.items,
	}

	for _, req := range r.requests {
		engineReq, err := deserializeRequest(req)
		if err != nil {
			log.Println(err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p serializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser2.ParseCityList, config.ParseCityList),nil
	}
}