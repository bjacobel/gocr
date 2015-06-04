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
	"time"
)

var urls, _ = regexp.Compile("(http(s)?://)([\\w-]+\\.)+[\\w-]+(/[\\w- ;,./?%&=]*)?")
var resizer, _ = regexp.Compile("236")

func main() {

	search := "inspirational quotes"
	spaces, _ := regexp.Compile(" ")
	search = string(spaces.ReplaceAll([]byte(search), []byte("%20"))) // lol golang's regexp library is fucked

	if doc, e := goquery.NewDocument("https://www.pinterest.com/search/pins/?q=" + search); e != nil {
		log.Fatal(e)
	} else {

		client, _ := gosseract.NewClient()
		var syncStart time.Time
		var asyncStart time.Time
		var syncDeltas float64
		var asyncDeltas float64
		iters := 10

		// do image retrieval and OCR 100 times without goroutines, avg the runtime
		for i := 1; i <= iters; i++ {
			syncStart = time.Now()

			doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
				OcrFromNode(s, client)
			})

			syncDeltas += time.Since(syncStart).Seconds()

			fmt.Print("\rCompleted synchronous iteration ", i)
		}

		fmt.Println("... done. Synchronous execution completed in avg ", syncDeltas/float64(iters), "s")

		// now do it with goroutines
		for i := 1; i <= iters; i++ {
			asyncStart = time.Now()

			var wg sync.WaitGroup

			doc.Find(".fadeContainer").Each(func(i int, s *goquery.Selection) {
				wg.Add(1)
				go OcrFromNode(s, client, &wg)
			})

			wg.Wait()

			asyncDeltas += time.Since(asyncStart).Seconds()

			fmt.Print("\rCompleted goroutine iteration ", i)
		}

		fmt.Println("... done. Goroutine execution completed in avg ", asyncDeltas/float64(iters), "s")
	}
}

func OcrFromNode(s *goquery.Selection, client *gosseract.Client, wg ...*sync.WaitGroup) {
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

	client.Src(path).Out()

	// fmt.Println("---------------\nText OCRed from " + imgLoc + ":\n")
	// fmt.Println(out + "\n")

	if wg != nil {
		wg[0].Done()
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
