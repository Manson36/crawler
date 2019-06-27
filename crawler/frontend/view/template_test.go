package view

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/frontend/model"
	common "github.com/crawler/crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(					//我们自己写的，一定认为它是合法的,，如果出错，must会panic
		template.ParseFiles("template.html"))	//在template中有自己的语法，返回值是*template和error

	out, err := os.Create("template.test.html")

	page := model.SearchResult{} //要填写的内容
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

	err = template.Execute(out, page)//有两个参数，io.writer和data。
	//首先我们先什么都不填，简单的方法：往屏幕输出，template.Execute(os.Stdout,page)
	//返回error，我们在数据合的过程中肯定会出现各种错误
	//我们想展示给别人看，我们把输出改一下，创建一个文件，此时，这个文件，我们就可以用浏览器查看了
	//然后，我们手动填写一些数据进去，查看
	if err != nil {
		panic(err)
	}
}
