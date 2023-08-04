package src

import (
	"fmt"
	"giao/pkg/util"
	"giao/pkg/util/custom_http"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

type Chapter struct {
	Url     string
	NextUrl string
	Title   string
	Content string
}

type EBook struct {
	host      string
	nextUrl   string
	name      string
	G         int // 限制协程数
	LinkCount int
	menuMap   map[int]*Chapter
	Site      SiteT
}

type SiteT interface {
	GetHost() string
	GetBookName(reader *goquery.Document) string
	GetMenuNextPageUrl(reader *goquery.Document) string
	GetMenuMap(e *EBook, reader *goquery.Document)

	// GetChapterContent
	// content := reader.Find("#rtext").Text()
	//nextHref, _ := reader.Find("#linkNext").Attr("href")
	//
	//c.NextUrl = ""
	//str1 := strings.Replace(nextHref, ".html", "", 1)
	//str2 := strings.Replace(c.Url, ".html", "", 1)
	//
	//if strings.HasPrefix(str1, str2) {
	//c.NextUrl = e.GetHost() + nextHref
	//}
	//
	//if c.Content == "" {
	//c.Content = strings.TrimSuffix(content, "\n")
	//} else {
	//c.Content = c.Content + strings.TrimSpace(content)
	//}
	GetChapterContent(c *Chapter, reader *goquery.Document, e *EBook)
}

func (e *EBook) Run(EbookUrl string) {
	start := time.Now()
	e.FetchMenuList(EbookUrl)

	e.goFetchData()
	fmt.Println("内容已获取完成")

	e.download()

	fmt.Println("花费：", time.Now().Sub(start).Seconds())
}

func (e *EBook) FetchMenuList(EbookUrl string) {
	u, _ := url.Parse(EbookUrl)
	e.host = u.Scheme + "://" + u.Hostname()
	e.nextUrl = u.String()

	e.LinkCount = 0
	e.menuMap = make(map[int]*Chapter)

	for e.nextUrl != "" {
		e.fetchMenu()
	}
	fmt.Println("章节：", e.LinkCount)
}

func (e *EBook) fetchMenu() {
	resp := custom_http.Fetch(e.nextUrl, nil)
	defer resp.Body.Close()

	utf8, err := charset.NewReader(resp.Body, "UTF-8")
	util.CheckErr(err)
	e.nextUrl = ""

	reader, err := goquery.NewDocumentFromReader(utf8)
	util.CheckErr(err)

	if e.name == "" {
		//bookName, _ := reader.Find("meta[property=\"og:novel:book_name\"]").Attr("content")
		e.name = e.Site.GetBookName(reader)
	}

	//nextPage, _ := reader.Find(".listpage .right>a.onclick").Attr("href")
	nextPage := e.Site.GetMenuNextPageUrl(reader)
	if nextPage != "" {
		e.nextUrl = e.host + nextPage
	}

	//chapterList := reader.Find(".book_last").Last()
	//chapterList.Find("a").Each(func(i int, selection *goquery.Selection) {
	//	//href, _ := selection.Attr("href")
	//	//title, _ := selection.Html()
	//	//e.LinkCount += 1
	//	//count := e.LinkCount
	//	//e.AddMenuMap(count, &src.Chapter{
	//	//	Url:   href,
	//	//	Title: title,
	//	//})
	//})
	e.Site.GetMenuMap(e, reader)
}

func (e *EBook) download() {
	dirName := "./download/"
	if ok, _ := util.IsExists(dirName); !ok {
		_ = os.Mkdir(dirName, 0644)
	}
	content := ""

	for i := 0; i < e.LinkCount+1; i++ {
		c, ok := e.menuMap[i]
		if ok {
			content = content + "\n\n" + c.Title + "\n\n" + c.Content
			//content = content + "\n\n" + "\n\n" + c.Content
		}
	}

	content = strings.Replace(content, "<br/><br/>", "\n", -1)

	fileName := dirName + e.name + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	util.CheckErr(err)
	_, _ = file.Write([]byte(content))
}

func (e *EBook) goFetchData() {
	g := e.G
	if g == 0 {
		g = 20
	}
	fmt.Println("开启协程:", g)

	var opts int32 = 0

	gLimit := util.NewGLimit(g)
	gLimit.Run(e.LinkCount, func(i int) {
		//gLimit.Run(1, func(i int) {
		atomic.AddInt32(&opts, 1)
		c, ok := e.menuMap[i+1]
		if ok {
			e.fetchContent(c)
		}

		fmt.Println("获取进度:", (float32(opts)/float32(e.LinkCount))*100)
	})
}

func (e *EBook) fetchContent(c *Chapter) {
	c.NextUrl = e.host + c.Url
	for c.NextUrl != "" {
		c.fetchPage(e)
	}
}

func (c *Chapter) fetchPage(e *EBook) {
	resp := custom_http.Fetch(c.NextUrl, nil)
	defer resp.Body.Close()

	utf8, err := charset.NewReader(resp.Body, "UTF-8")
	util.CheckErr(err)

	reader, err := goquery.NewDocumentFromReader(utf8)
	util.CheckErr(err)

	reader.Find("*").Each(func(i int, selection *goquery.Selection) {
		// 去除js
		if selection.Is("script") {
			selection.Remove()
		}
	})

	// html, err := reader.Find("#chaptercontent").Html()
	// if err != nil {
	// 	return
	// }
	// util.DD(html)

	//content, _ := reader.Find("#chaptercontent").Html()
	//nextHref, _ := reader.Find("#pb_next").Attr("href")
	//
	//c.nextUrl = ""
	//str1 := strings.Replace(nextHref, ".html", "", 1)
	//str2 := strings.Replace(c.url, ".html", "", 1)
	//
	//if strings.HasPrefix(str1, str2) {
	//	c.nextUrl = e.host + nextHref
	//}
	//compile, _ := regexp.Compile(`第\(.*?\)页`)
	//util.CheckErr(err)
	//all := compile.ReplaceAllString(content, "\n")
	//
	//compile, _ = regexp.Compile(`：\w.+?.com`)
	//util.CheckErr(err)
	//all = compile.ReplaceAllString(all, "\n")
	//
	//content = strings.ReplaceAll(all, "<br/>", "\n")
	//if c.content == "" {
	//	c.content = strings.TrimSuffix(content, "\n")
	//} else {
	//	c.content = c.content + strings.TrimSpace(content)
	//}

	e.Site.GetChapterContent(c, reader, e)
}

func (e *EBook) AddMenuMap(k int, c *Chapter) {
	e.menuMap[k] = c
}

func (e *EBook) GetHost() string {
	return e.host
}
