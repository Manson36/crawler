package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
	itemsaver "github.com/crawler/crawler_distributed/persist/client"
	worker "github.com/crawler/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: processor}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}

