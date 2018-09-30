package parser

import (
	"GoStudy/Chapter17/crawlerPro/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	url := "http://album.zhenai.com/u/108906739"
	contents, err := ioutil.ReadFile("profile_info.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s",contents)
	user := ParseProfile(contents, url, "小顺儿")
	if len(user.Items) != 1 {
		t.Errorf("wrong result ,expect 1, but got%d", len(user.Items))
	}
	item := user.Items[0]
	var expectUserProfile = model.Profile{
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
	}
	if expectUserProfile != item.Payload {
		t.Errorf("wrong profile ,expect %v,but got %v", expectUserProfile, item)
	}
	fmt.Println("get profile", item)
}
