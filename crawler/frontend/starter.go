package main

import (
	"github.com/crawler/crawler/frontend/controller"
	"net/http"
)

func main() {
	//在这选用http函数，前面我们创建了handler接口内ServeHTTP，所以在这使用http.handle；还可以使用http.HandleFunc
	http.Handle("/search", controller.CreateSearchResultHandler(
		"crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil )//第二个handler我们一般传nil
	if err != nil {
		panic(err)//这里出错，直接就挂掉了
	}
	//此时，我们就可以在网页上使用localhost：8888/search 访问

}
