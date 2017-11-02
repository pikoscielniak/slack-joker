package fetcher

import (
	"time"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func Fetch() {
	doc, err := goquery.NewDocument("https://www.dowcipy.jeja.pl/")
	if err != nil {
		log.Fatal(err)
	}
	jokeDiv := doc.Find("div.dow-left-text").First()
	if jokeDiv != nil {

		joke := jokeDiv.Find("p").Text()
		fmt.Printf("%s\n", joke)
	}
}
