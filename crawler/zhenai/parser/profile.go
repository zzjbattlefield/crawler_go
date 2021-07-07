package parser

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"log"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)

var idUrlRe = regexp.MustCompile(
	`.*album\.zhenai\.com/u/([\d]+)`)

//解析用户数据
func ParseProfile(content []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	if age, err := strconv.Atoi(extractString(content, ageRe)); err != nil {
		log.Printf("get user age error:%v", err)
	} else {
		profile.Age = age
	}
	profile.Marriage = extractString(content, marriageRe)
	height, err := strconv.Atoi(
		extractString(content, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(
		extractString(content, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(
		content, incomeRe)
	profile.Gender = extractString(
		content, genderRe)
	profile.Car = extractString(
		content, carRe)
	profile.Education = extractString(
		content, educationRe)
	profile.Hokou = extractString(
		content, hokouRe)
	profile.House = extractString(
		content, houseRe)
	profile.Occupation = extractString(
		content, occupationRe)
	profile.Xinzuo = extractString(
		content, xinzuoRe)
	result := engine.ParseResult{
		Item: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url), idUrlRe),
				PayLoad: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(content, -1)
	//猜你喜欢页面
	for _, match := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ProfileParser(string(match[2])),
		})
	}

	return result
}

//通过传入的正则表达式从content中获得数据
func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func ProfileParser(name string) func(content []byte, url string) engine.ParseResult {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
