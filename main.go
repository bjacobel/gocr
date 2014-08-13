package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dchest/uniuri"
	"github.com/otiai10/gosseract"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var urls, _ = regexp.Compile("(http(s)?://)?([\\w-]+\\.)+[\\w-]+(/[\\w- ;,./?%&=]*)?")
var resizer, _ = regexp.Compile("236")

func main() {
	var wg sync.WaitGroup

	search := "inspirational quotes"
	spaces, _ := regexp.Compile(" ")
	search = string(spaces.ReplaceAll([]byte(search), []byte("%20"))) // lol golang's regexp library is fucked

	if doc, e := goquery.NewDocument("http://www.pinterest.com/search/pins/?q=" + search); e != nil {
		log.Fatal(e)
	} else {

		client, _ := gosseract.NewClient()

		doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
			wg.Add(1)
			go OcrFromNode(s, &wg, client)
		})
	}

	wg.Wait()
}

func OcrFromNode(s *goquery.Selection, wg *sync.WaitGroup, client *gosseract.Client) {
	html, err := (*s).Html()
	handleErr(err)

	imgLoc := string(resizer.ReplaceAll(urls.Find([]byte(html)), []byte("736")))

	imgData, err := http.Get(imgLoc)
	handleErr(err)

	defer func() {
		err := imgData.Body.Close()
		handleErr(err)
	}()

	body, err := ioutil.ReadAll(imgData.Body)
	handleErr(err)

	path := "images/" + uniuri.New() + ".jpg"

	err = ioutil.WriteFile(path, body, 0644)
	handleErr(err)

	out, _ := client.Src(path).Out()

	fmt.Println("---------------\nText OCRed from " + imgLoc + ":\n")
	fmt.Println(out + "\n")

	wg.Done()
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
