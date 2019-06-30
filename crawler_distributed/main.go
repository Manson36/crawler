package main

import (
	"flag"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
	"github.com/crawler/crawler_distributed/config"
	itemsaver "github.com/crawler/crawler_distributed/persist/client"
	"github.com/crawler/crawler_distributed/rpcsupport"
	worker "github.com/crawler/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String(
		"worker_host", "", "worker hosts (comma separated)")//以逗号分隔
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(//这里是slice，我们要把逗号分隔解析出来
		strings.Split(*workerHosts, ","))

	processor:= worker.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {
	//我们要对这些host一个个去连，练好了之后，形成一个ClientPool，然后通过一个channel发给worker
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	//clients我们已经建好了，接下来我们就要往channel里分发，
	//建一个channel，分发是在一个goroutine里面,这是go语言一种常用的写法
	out := make(chan *rpc.Client)
	go func() {
		//分发怎么分发，可以随机、轮流分发，当然轮流比较简单
		//只有一个for，发一轮就完了，所以要再添加一个for
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}