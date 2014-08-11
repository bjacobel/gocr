package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
)

var urls, _ = regexp.Compile("(http(s)?://)?([\\w-]+\\.)+[\\w-]+(/[\\w- ;,./?%&=]*)?")
var resizer, _ = regexp.Compile("236")

func main() {
	var doc *goquery.Document
	var e error

	search := "inspirational quotes"
	spaces, _ := regexp.Compile(" ")
	search = string(spaces.ReplaceAll([]byte(search), []byte("%20"))) // lol golang's regexp library is fucked

	if doc, e = goquery.NewDocument("http://www.pinterest.com/search/pins/?q=" + search); e != nil {
		log.Fatal(e)
	}

	doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
		go getImgSrcFromNode(s)
		fmt.Println()
	})
}

func getImgSrcFromNode(s *goquery.Selection) string {
	html, _ := (*s).Html()
	return string(resizer.ReplaceAll(urls.Find([]byte(html)), []byte("736")))
}
