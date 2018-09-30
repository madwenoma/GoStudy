package controller

import (
	"reflect"
	"GoStudy/Chapter17/crawlerPro/engine"
	"regexp"
	"GoStudy/Chapter17/crawlerPro/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"GoStudy/Chapter17/crawlerPro/frontend/model"
	"context"
)

//type SearchResultHandler struct {
//	view   view.SearchResultView
//	client *elastic.Client
//}
//
//func CreateSearchResultHandler(tem string) SearchResultHandler {
//	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"),
//		elastic.SetSniff(false))
//	if err != nil {
//		panic(err)
//	}
//	return SearchResultHandler{
//		view:   view.CreateSearchResultView(tem),
//		client: client,
//	}
//}
//
//func (srHandler SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	q := strings.TrimSpace(req.FormValue("q"))
//	from, err := strconv.Atoi(req.FormValue("from"))
//	if err != nil {
//		from = 0
//	}
//	//fmt.Fprintf(w, "q=%s,from=%d", q, from)
//	var page model.SearchResult
//	page,err = srHandler.getSearchResult(q, from)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	err = srHandler.view.Render(w, page)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (srHandler SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
//	var result model.SearchResult
//	resp, err := srHandler.client.Search("dating_profile").
//		Query(elastic.NewQueryStringQuery(q)).
//		From(from).Do(context.Background())
//	if err != nil {
//		return result, err
//	}
//	result.Hits = resp.TotalHits()
//	result.Start = from
//	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
//	return result, nil
//}
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"),
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.
	//Search(config.ElasticIndex).
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
		rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom = (result.Start - 1) / pageSize * pageSize
	}

	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:") //$1能拿到（）里的字段
}
