package joker

import (
	"github.com/PuerkitoBio/goquery"
	"math/rand"
)

const (
	maxJokes = 3
)

var maxJokesExceededMsgs = []string{
	"foch",
	"dość już",
	"nie, nie i jeszcze raz nie",
	"nieczynne",
	"dość tego",
	"hm... ta",
	"przerwa",
	"nie wiem jak Wy ale ja już mam na dzisiaj fajrant",
	"a może by tak popracować",
}

var readJokes = newRotateQueue(maxJokes)

func getDefaultMsg() string {
	i := rand.Intn(len(maxJokesExceededMsgs))
	return maxJokesExceededMsgs[i]
}

func fetch() (string, error) {
	doc, err := goquery.NewDocument("https://www.dowcipy.jeja.pl/")
	if err != nil {
		return "", err
	}
	returnJoke := getDefaultMsg()
	sel := doc.Find("div.dow-left-text")
	sel.EachWithBreak(func(i int, s *goquery.Selection) bool {
		stop := false
		next := !stop
		if i >= maxJokes {
			return stop
		}
		jokeSel := s.Find("p")
		if jokeSel == nil {
			return stop
		}
		joke := jokeSel.Text()
		if readJokes.hasJoke(joke) {
			return next
		}
		readJokes.add(joke)
		returnJoke = joke
		return stop
	})
	return returnJoke, nil
}
