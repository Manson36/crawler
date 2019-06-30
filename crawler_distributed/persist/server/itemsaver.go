package main

import (
	"flag"
	"fmt"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/persist"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{//Save方法是指针接受者，这里也要使用指针
		Client: client,
		Index: index,
	})
}
