package client

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string ) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			//我们没收到一个元素应该做的事Call rpc
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item saver error: saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

