package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/persist"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}

