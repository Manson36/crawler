package view

import (
	"github.com/crawler/practice/htmltemplate/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func NewSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template.Must(template.ParseFiles(fileName)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}