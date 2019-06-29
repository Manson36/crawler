package worker

import (
	"errors"
	"fmt"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}
func SerializeRequest(r engine.Request) Request {
	 name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name:name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParserResult {
	result :=  engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("deserializing request error: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

//我们如何将字符串转化为一个函数呢？我们有两种做法：
// 一种复杂的是把Parser的名字都注册到一个全局的map中，然后我们从map中找到对应的parser函数
//第二种就是switch case
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}

	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

