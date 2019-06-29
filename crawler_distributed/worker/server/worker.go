package main

import (
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"github.com/crawler/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		config.WorkerPort0, worker.CrawlService{}))
}
