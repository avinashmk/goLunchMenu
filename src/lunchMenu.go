package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"
	"strings"
	"time"
)

import (
	"github.com/PuerkitoBio/goquery"
)

var (
	swedishDays = map[string]string{
		"Monday":    "Måndag",
		"Tuesday":   "Tisdag",
		"Wednesday": "Onsdag",
		"Thursday":  "Torsdag",
		"Friday":    "Fredag"}
)

func main() {
	fmt.Println("Hello World!")

	today := time.Now().Weekday().String()
	fmt.Println("Today:", today)

	// Create a client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Describe the HTTP request
	request, err := http.NewRequest("GET", "https://vasakronan.foodbycoor.se/zenit/sv/meny", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Kinda Lunch Menu Consolidator App")

	// Send request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Open a file
	//file, _ := os.Open("../smaple.html")

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find("body .menu-row").EachWithBreak(func(_ int, sBody *goquery.Selection) bool {
		if text(sBody) == swedishDays[today] {
			fmt.Println("Day:", text(sBody))
			sBody.NextUntil("h2").Find(".element.title").Each(func(_ int, sFiltered *goquery.Selection) {
				fmt.Println(text(sFiltered))
			})
			return false
		}
		return true
	})
}

func text(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}
