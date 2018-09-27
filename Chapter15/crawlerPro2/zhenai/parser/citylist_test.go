package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("zhenai_citylist.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Print(contents)
	const resultSize = 470

	result := ParseCityList(contents)
	//fmt.Println(result.Items)
	if len(result.Items) != resultSize {
		t.Errorf("result size shoule be %d,but got%d", resultSize, len(result.Items))
	}

	expectCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	expectUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	for i, value := range expectCities {
		if result.Items[i].(string) != value {
			t.Errorf("wrong city")
		}
	}

	for i, value := range expectUrls {
		if result.Requests[i].Url != value {
			t.Errorf("wrong url")
		}
	}
}
