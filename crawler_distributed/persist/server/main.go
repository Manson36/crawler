package main

import (
	"github.com/crawler/crawler_distributed/persist"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(":1234", &persist.ItemSaverService{//Save方法是指针接受者，这里也要使用指针
		Client: client,
		Index: "dating_profile",
	})
}
