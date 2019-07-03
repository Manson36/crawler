package persist

import (
	"context"
	"errors"
	"fmt"
	"github.com/crawler/crawler/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//sniff 是维护elastic的集群的状态，但是elastic没有运行在本机上，实在docker上，docker只是一个内网，外面看不见，无法sniff
	client, err := elastic.NewClient(elastic.SetSniff(false))//Must turn off sniff in docker
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

			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item saver error: saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {

	if 	item.Type == "" {
		return errors.New("must supply Type")
	}

	//Index()表示存储信息，后面参数是存储的位置
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	//%+v可以将结构体的字段名打印出
	fmt.Printf("%+v", resp)
	return nil
}