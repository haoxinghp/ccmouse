package engine

import (
	"dali.cc/ccmouse/crawler/fetcher"
	"github.com/gpmgo/gopm/modules/log"
)

func worker (r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	//log.Warn("Fetching %s", r.Url)
	if err != nil {
		log.Error("请求[%s]失败：%s", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body, r.Url), nil
}
