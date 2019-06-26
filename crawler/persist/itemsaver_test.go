package persist

import (
	"context"
	"encoding/json"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	//通过index，type，id 查找内容，是否可以匹配
	//TODO:Try to start up elastic search here using docker go client
	//目的：不依赖外界环境
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	//首先save expect item
	err = save(client,  expected, index)
	if err != nil {
		panic(err)
	}

	//然后Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%+v", resp)
	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//Verify result
	if expected != actual {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
