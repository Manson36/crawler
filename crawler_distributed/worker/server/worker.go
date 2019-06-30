package main

import (
	"flag"
	"fmt"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"github.com/crawler/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {//这里port类型是一个*int
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", port), worker.CrawlService{}))
}
