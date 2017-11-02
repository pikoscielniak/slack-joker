package fetcher

import (
	"time"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func Fetch() (string, error) {
	doc, err := goquery.NewDocument("https://www.dowcipy.jeja.pl/")
	if err != nil {
		return "", err
	}
	jokeDiv := doc.Find("div.dow-left-text").First()
	if jokeDiv == nil {
		return "", errors.New("Not found")
	}
	return jokeDiv.Find("p").Text(), nil
}
