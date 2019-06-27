package view

import (
	"github.com/crawler/crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func NewSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(fileName)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {//data 我们认为一定是searchresult的类型
	return s.template.Execute(w, data) //此时，template view已经包装好了
}