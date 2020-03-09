package Entre

import (
	"Util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	WEB_SCRAPE = false
	URL        = "https://vasakronan.foodbycoor.se/restaurang-entre/restaurangen/restaurangens-meny"
	NAME       = "Entre"
)

type Entre struct{}

func (z *Entre) PrintMenu() {
	doc := Util.GetHtmlBody(URL)
	doc.Find("body .menu-row").EachWithBreak(func(_ int, sBody *goquery.Selection) bool {
		bodyText := Util.GetText(sBody)
		if Util.CheckToday(bodyText) {
			fmt.Printf("%s\n\n", bodyText)
			sBody.NextUntil("h2").Find(".element.description ").Each(func(_ int, sFiltered *goquery.Selection) {
				fmt.Println(Util.GetText(sFiltered))
			})
			return false
		}
		return true
	})
}

func (z *Entre) Name() string {
	return NAME
}

func (z *Entre) WebScrape() bool {
	return WEB_SCRAPE
}
