package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
)

func main() {
	var doc *goquery.Document
	var e error

	search := "inspirational quotes"
	r, _ := regexp.Compile(" ")
	search = string(r.ReplaceAll([]byte(search), []byte("%20"))) // lol golang's regexp library is fucked

	if doc, e = goquery.NewDocument("http://www.pinterest.com/search/pins/?q=" + search); e != nil {
		log.Fatal(e)
	}

	doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
		html, _ := (*s).Html()
		fmt.Println(html)
	})
}
