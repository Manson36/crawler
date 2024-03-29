package parser

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"github.com/crawler/crawler_distributed/config"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
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
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
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
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte,url string, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload:profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),				//之前：string(m[2])不能直接填写，需先定义，是因为定义m值在变化，最后只能拿到一个值
			Parser: NewProfileParser(string(m[2])),//string(m[2])可以直接填进去，不用先定义name，是因为在这里是函数调用，是值拷贝
		})
	}

	return result
}

func extractString (contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) > 2 {
		return string(match[1])
	} else {
		return ""
	}
}

//func ProfileParser(name string) engine.ParserFunc {
//	return func(c []byte, url string) engine.ParserResult {
//		return ParseProfile(c, url, name)
//	}
//}

//分布式，自己定义一个ProfileParser, 前面的代码就不需要了
type ProfileParser struct {
	userName string
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{name}
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParserResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}
