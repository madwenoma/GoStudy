package parser

import (
	"GoStudy/Chapter16/crawlerPro/engine"
	"GoStudy/Chapter16/crawlerPro/model"
	"regexp"
	"strconv"
)

//<span class="label">性别：</span><span field="">女</span></td>
//<span class="label">年龄：</span>29岁</td>
//<span class="label">身高：</span><span field="">169CM</span></td>
//<span class="label">体重：</span><span field="">52KG</span></td>
//<span class="label">月收入：</span>3001-5000元</td>
//<span class="label">婚况：</span>未婚</td>
//<span class="label">学历：</span>大学本科</td>
//<span class="label">职业： </span>会计</td>
//<span class="label">籍贯：</span>四川阿坝</td>
//<span class="label">星座：</span>魔羯座</td>
//<span class="label">住房条件：</span><span field="">和家人同住</span></td>
//<span class="label">是否购车：</span><span field="">未购车</span></td>
//const name =

//var nameReg = regexp.MustCompile(``)
var genderReg = regexp.MustCompile(`<span class="label">性别：</span><span field="">([^<]+)</span>`)
var ageReg = regexp.MustCompile(`<span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightReg = regexp.MustCompile(`<span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)
var weightReg = regexp.MustCompile(`<span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`)
var incomeReg = regexp.MustCompile(`<span class="label">月收入：</span>([^<]+)</td>`)
var marriageReg = regexp.MustCompile(`<span class="label">婚况：</span>([^<]+)</td>`)
var educationReg = regexp.MustCompile(`<span class="label">学历：</span>([^<]+)</td>`)
var occupationReg = regexp.MustCompile(`<span class="label">职业： </span>([^<]+)</td>`)
var hokouReg = regexp.MustCompile(`<span class="label">籍贯：</span>([^<]+)</td>`)
var xingzuoReg = regexp.MustCompile(`<span class="label">星座：</span>([^<]+)</td>`)
var houseReg = regexp.MustCompile(`<span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carReg = regexp.MustCompile(`<span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var idUrlReg = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	user := model.Profile{}
	user.Name = name
	user.Gender = extractStr(contents, genderReg)
	age, err := strconv.Atoi(extractStr(contents, ageReg))
	if err == nil {
		user.Age = age
	}

	height, err := strconv.Atoi(extractStr(contents, heightReg))
	if err == nil {
		user.Height = height
	}

	weight, err := strconv.Atoi(extractStr(contents, weightReg))
	if err == nil {
		user.Weight = weight
	}
	user.Income = extractStr(contents, incomeReg)
	user.Marriage = extractStr(contents, marriageReg)
	user.Education = extractStr(contents, educationReg)
	user.Occupation = extractStr(contents, occupationReg)
	user.Hokou = extractStr(contents, hokouReg)
	user.Xingzuo = extractStr(contents, xingzuoReg)
	user.House = extractStr(contents, houseReg)
	user.Car = extractStr(contents, carReg)

	// fmt.Println("get user:",user)

	return engine.ParseResult{Items: []engine.Item{{
		Url:     "",
		Id:      extractStr([]byte(url), idUrlReg),
		Type:    "zhenai",
		Payload: user,
	}}}
}

func extractStr(contents []byte, reg *regexp.Regexp) string {

	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		str := string(match[1])
		return str
	}

	return ""
}
