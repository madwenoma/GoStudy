package persist

import (
	"testing"
	"GoStudy/Chapter16/crawlerPro/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"GoStudy/Chapter16/crawlerPro/engine"
)

func TestItemSaver(t *testing.T) {
	var expectItem = engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "小顺儿",
			Gender:     "女",
			Age:        29,
			Height:     169,
			Weight:     52,
			Income:     "3001-5000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "会计",
			Hokou:      "四川阿坝",
			Xingzuo:    "魔羯座",
			House:      "和家人同住",
			Car:        "未购车",
		},
	}

	err := save(expectItem)
	if err != nil {
		t.Errorf("error:%v", err)
	}

	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"), elastic.SetSniff(false))
	if err != nil {
		t.Errorf("error:%v", err)
	}
	profileFromEs, err := client.Get().Index("dating_profile").Type(expectItem.Type).Id(expectItem.Id).Do(context.Background())
	if err != nil {
		t.Errorf("error:%v", err)
	}
	t.Logf("ss:%s", profileFromEs.Source)

	var actualItem engine.Item
	//expectProfileJson,err:=json.Marshal(expectProfile)
	//t.Logf("%s",expectProfileJson)
	sourceByte, _ := profileFromEs.Source.MarshalJSON()
	err = json.Unmarshal(sourceByte, &actualItem)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actualItem)
	actualItem.Payload = actualProfile
	if expectItem != actualItem {
		t.Errorf("error,expect %v,got %v", expectItem, actualItem)
	}

}
