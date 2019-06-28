package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{Scheduler:&scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

