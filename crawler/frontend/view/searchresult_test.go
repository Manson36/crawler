package view

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/frontend/model"
	common "github.com/crawler/crawler/model"
	"os"
	"testing"
)

//将template.test 修改，换为searchresult中的函数
func TestSearchResultView_Render(t *testing.T) {
	template := NewSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Id: "108906739",
		Type:"zhenai",
		Payload:common.Profile{ //两个model包重名了，需要将其中的一个重命名
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
	//我们在其中插十个记录,再运行 查看结果
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = template.Render(out, page)
	if err != nil {
		panic(err)
	}
}
