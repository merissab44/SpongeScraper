package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gocolly/colly"
)

type SpongebobEpisodes struct {
	Title   string
	Viewers string
	Episode string
}

func createJson(characters []SpongebobEpisodes) {
	jsonFile, _ := json.MarshalIndent(characters, "", " ")
	_ = ioutil.WriteFile("Spongebob.json", jsonFile, 0644)
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	character := make([]SpongebobEpisodes, 0)
	//create callback for links
	c.OnHTML("table.general", func(e *colly.HTMLElement) {
		e.ForEach("tr.general-header", func(_ int, e *colly.HTMLElement) {
			newCharacter := SpongebobEpisodes{}
			newCharacter.Title = e.ChildText("a")
			newCharacter.Viewers = e.ChildText("center")
			newCharacter.Episode = e.ChildText("b")
			character = append(character, newCharacter)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	// start scraping
	c.Visit("https://spongebob.fandom.com/wiki/List_of_episodes")

	createJson(character)
}
