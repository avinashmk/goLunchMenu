package Util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	Today = time.Now().Weekday().String()

	swedishDay = map[string]string{
		"Monday":    "Måndag",
		"Tuesday":   "Tisdag",
		"Wednesday": "Onsdag",
		"Thursday":  "Torsdag",
		"Friday":    "Fredag"}
)

func GetText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}

func CheckToday(day string) bool {
	if strings.Contains(day, Today) || strings.Contains(day, swedishDay[Today]) {
		return true
	}
	return false
}

func GetHtmlBody(url string) (doc *goquery.Document) {

	doc = nil

	// Create a client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Describe the HTTP request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("UserAgent", "Kinda_Lunch_Menu_Consolidator_App")

	// Send request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read the body to a goquery Document
	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func PrintHeader(header string) {
	fmt.Println("\n================================================================")
	fmt.Printf("%s: ", header)
}
