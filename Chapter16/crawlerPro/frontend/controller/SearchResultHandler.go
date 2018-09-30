package controller

import (
	"GoStudy/Chapter16/crawlerPro/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strconv"
	"strings"
	"GoStudy/Chapter16/crawlerPro/frontend/model"
	"context"
	"reflect"
	"GoStudy/Chapter16/crawlerPro/engine"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(tem string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(tem),
		client: client,
	}
}

func (srHandler SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s,from=%d", q, from)
	var page model.SearchResult
	page,err = srHandler.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = srHandler.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srHandler SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := srHandler.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	return result, nil
}
