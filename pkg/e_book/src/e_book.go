package src

import (
	"fmt"
	"giao/pkg/util"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
)

type chapter struct {
	title   string
	content string
}

type EBook struct {
	host    string
	nextUrl string
	name    string
	menuUrl string
	chapter
	content []chapter
}

func (e *EBook) Run(EbookUrl string) {
	u, _ := url.Parse(EbookUrl)
	e.host = u.Scheme + "://" + u.Hostname()
	e.nextUrl = u.String()

	for e.nextUrl != "" {
		e.FetchPage()
	}

	e.download()

}

func (e *EBook) FetchPage() {
	res, _ := http.Get(e.nextUrl)
	e.nextUrl = ""
	defer res.Body.Close()
	reader, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	if e.name == "" {
		e.name = reader.Find("#_bqgmb_h1").Text()
	}
	title := reader.Find("#nr_title").Text()
	content := reader.Find("#nr1").Text()
	menuHref, _ := reader.Find("#pt_mulu").Attr("href")
	nextHref, _ := reader.Find("#pt_next").Attr("href")

	fmt.Println(title)

	if title != e.title {
		e.content = append(e.content, e.chapter)
		e.chapter.content = ""
	}
	e.chapter.title = title
	e.chapter.content = e.chapter.content + content
	if menuHref != nextHref {
		e.nextUrl = e.host + nextHref
	} else {
		e.content = append(e.content, e.chapter)
	}

}

func (e *EBook) download() {
	dirName := "./download/"
	if ok, _ := util.IsExists(dirName); !ok {
		_ = os.Mkdir(dirName, 0644)
	}
	content := ""

	for _, c := range e.content {
		content = content + "\n\n" + c.title + "\n\n" + c.content
	}

	fileName := dirName + e.name + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	util.CheckErr(err)
	_, _ = file.Write([]byte(content))
}
