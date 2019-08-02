package main

import (
	"github.com/crawler/restart/singletask/engine"
	"github.com/crawler/restart/singletask/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
