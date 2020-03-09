package Glaze

import (
	"Util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	WEB_SCRAPE = false
	URL        = "https://vasakronan.foodbycoor.se/glaze/sv/meny"
	NAME       = "Glaze"
)

type Glaze struct{}

func (z *Glaze) PrintMenu() {
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

func (z *Glaze) Name() string {
	return NAME
}

func (z *Glaze) WebScrape() bool {
	return WEB_SCRAPE
}
