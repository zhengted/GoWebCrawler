package parser

import (
	"GoWebCrawler/crawler/engine"
	"GoWebCrawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parserProfile(contents, "厌与深情记得笑i", "")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0]

	expected := engine.Item{
		"http://localhost:8080/mock/album.zhenai.com/u/7143522202848495805",
		"zhenai",
		"7143522202848495805",
		model.Profile{
			Name:       "厌与深情记得笑i",
			Gender:     "女",
			Age:        10,
			Height:     1,
			Weight:     116,
			Income:     "2001-3000元",
			Marriage:   "未婚",
			Education:  "硕士",
			Occupation: "销售",
			Hokou:      "东莞市",
			Xinzuo:     "天秤座",
			House:      "有房",
			Car:        "有豪车",
		},
	}
	if profile != expected {
		t.Errorf("expected:%v but got: %v", expected, profile)
	}
}
