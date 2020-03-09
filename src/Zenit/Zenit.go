package Zenit

import (
	"Util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	WEB_SCRAPE = false
	URL        = "https://vasakronan.foodbycoor.se/zenit/sv/meny"
	NAME       = "Zenit"
)

type Zenit struct{}

func (z *Zenit) PrintMenu() {
	doc := Util.GetHtmlBody(URL)
	doc.Find("body .menu-row").EachWithBreak(func(_ int, sBody *goquery.Selection) bool {
		bodyText := Util.GetText(sBody)
		if Util.CheckToday(bodyText) {
			fmt.Printf("%s\n\n", bodyText)
			sBody.NextUntil("h2").Find(".element.title").Each(func(_ int, sFiltered *goquery.Selection) {
				fmt.Println(Util.GetText(sFiltered))
			})
			return false
		}
		return true
	})
}

func (z *Zenit) Name() string {
	return NAME
}

func (z *Zenit) WebScrape() bool {
	return WEB_SCRAPE
}
