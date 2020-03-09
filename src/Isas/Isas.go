package Isas

import (
	"Util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	WEB_SCRAPE = true
	URL        = "https://isasspis.gastrogate.com/lunch/"
	NAME       = "Isas Spis"
)

type Isas struct {
	today string
}

func (z *Isas) PrintMenu() {
	doc := Util.GetHtmlBody(URL)
	doc.Find("body .table.lunch_menu.animation").Contents().EachWithBreak(func(_ int, sBody *goquery.Selection) bool {
		bodyText := Util.GetText(sBody)
		if Util.CheckToday(bodyText) {
			fmt.Printf("%s\n\n", bodyText)
			sBody.NextUntil(".lunch-day-header").Find(".td_title").Each(func(_ int, sFiltered *goquery.Selection) {
				fmt.Println(Util.GetText(sFiltered))
			})
			return false
		}
		return true
	})
}

func (z *Isas) Name() string {
	return NAME
}

func (z *Isas) WebScrape() bool {
	return WEB_SCRAPE
}
