package controller

import (
	"context"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/frontend/model"
	"github.com/crawler/crawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

//第二次操作：首先我们要初始化一个searchResultHander
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))//docker中一定要SetSniff
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view: view.NewSearchResultView(template),
		client: client, //我们要在前面先建一下
	}
}


//我们要使用它实现一个Handler in net/http/server.go
//这个handler，我们想做这样一件事情。比如访问localhost：8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//req.FormValue()//首先拿到上面提到的参数，在地道一点，把空格给去掉
	q := strings.TrimSpace(req.FormValue("q"))
	//from后面是一个整数，Atoi可能会出错，加一个err
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0 //忽略掉这个错误
	}

	//fmt.Fprintf(w, "q=%s, from=%d", q, from)
	//我们q和from都有了
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)//统一的错误处理还没有做，就先使用这种方法
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult//首先将result定义出来
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)). //NewQuery中有很多选项，我们只需要最简单的stringQuery
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	//没有错误，我们开始包装
	result.Hits = resp.TotalHits()//totalHits中都有包装，是int64，需要转换为int或者将Hits改为int64
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))//复杂的一大堆，查看each方法

	return result, err
}



