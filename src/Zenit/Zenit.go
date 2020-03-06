package Zenit

import (
	"Util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	WEB_SCRAPE = false
	URL        = "https://vasakronan.foodbycoor.se/zenit/sv/meny"
)

type Zenit struct{}

func (z *Zenit) PrintMenu() (result bool) {
	doc := Util.GetHtmlBody(URL)

	doc.Find("body .menu-row").EachWithBreak(func(_ int, sBody *goquery.Selection) bool {
		bodyText := Util.GetText(sBody)
		if Util.CheckToday(bodyText) {
			fmt.Println("Day:", bodyText)
			sBody.NextUntil("h2").Find(".element.title").Each(func(_ int, sFiltered *goquery.Selection) {
				fmt.Println(Util.GetText(sFiltered))
			})
			return false
		}
		return true
	})
	result = true
	return
}
