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
	spaces, _ := regexp.Compile(" ")
	search = string(spaces.ReplaceAll([]byte(search), []byte("%20"))) // lol golang's regexp library is fucked

	if doc, e = goquery.NewDocument("http://www.pinterest.com/search/pins/?q=" + search); e != nil {
		log.Fatal(e)
	}

	urls, _ := regexp.Compile("(http(s)?://)?([\\w-]+\\.)+[\\w-]+(/[\\w- ;,./?%&=]*)?")
	resizer, _ := regexp.Compile("236")

	doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
		html, _ := (*s).Html()
		url := string(resizer.ReplaceAll(urls.Find([]byte(html)), []byte("736")))
		fmt.Println(url)
	})
}
