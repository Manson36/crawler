package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	//Start TestItemSaver
	go serveRpc(host, "test1")
	time.Sleep(time.Second)//goroutine中sever还没有起来，client就连上，会出错误

	//Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//Call Save
	item := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Id: "108906739",
		Type:"zhenai",
		Payload:model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc,
		item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s; err: %s", result, err)
	}
}
