package main

import (
	"Util"
	"Zenit"
	"fmt"
	//"os"
	//"strings"
)

type Restaurant interface {
	PrintMenu() bool
}

func main() {
	fmt.Println("Hello", Util.Today, "!")

	var restaurants [1]Restaurant
	restaurants[0] = &Zenit.Zenit{}

	for _, r := range restaurants {
		r.PrintMenu()
	}
	// Open a file
	//file, _ := os.Open("../smaple.html")

}
