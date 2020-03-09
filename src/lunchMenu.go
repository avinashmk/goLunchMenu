package main

import (
	"Entre"
	"Glaze"
	"Isas"
	"Util"
	"Zenit"
	"fmt"
)

type Restaurant interface {
	PrintMenu()
	Name() string
	WebScrape() bool
}

func main() {
	fmt.Println("Hello", Util.Today, "Lunch!")

	restaurants := []Restaurant{
		&Zenit.Zenit{},
		&Glaze.Glaze{},
		&Entre.Entre{},
		&Isas.Isas{},
	}
	for _, r := range restaurants {
		Util.PrintHeader(r.Name())
		if r.WebScrape() {
			r.PrintMenu()
		} else {
			fmt.Println("Web-scraping disabled.")
		}
	}
}
