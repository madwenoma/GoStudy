package parser

import (
	"testing"
	"io/ioutil"
	"GoStudy/crawlerPro/model"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_info.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s",contents)
	user := ParseProfile(contents, "小顺儿")
	if len(user.Items) != 1 {
		t.Errorf("wrong result ,expect 1, but got%d", len(user.Items))
	}
	profile := user.Items[0].(model.Profile)
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
	if expectUserProfile != profile {
		t.Errorf("wrong profile ,expect %v,but got %v", expectUserProfile, profile)
	}
	fmt.Println("get profile",profile)
}
