package www_vbiquge_co

import (
	"giao/pkg/e_book/src"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Site struct {
}

func (s *Site) GetHost() string {
	//https://www.ddyueshu.com/12_12416/
	return "www.vbiquge.co"
}

func (s *Site) GetBookName(reader *goquery.Document) string {
	bookName, ok := reader.Find("meta[property=\"og:novel:book_name\"]").Attr("content")
	if !ok {
		panic("书名未找到")
	}
	return bookName
}

func (s *Site) GetMenuNextPageUrl(reader *goquery.Document) string {
	return ""
}

func (s *Site) GetMenuMap(e *src.EBook, reader *goquery.Document) {
	chapterList := reader.Find("#list-chapterAll")

	chapterList.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Html()
		e.LinkCount += 1
		count := e.LinkCount

		e.AddMenuMap(count, &src.Chapter{
			Url:   href,
			Title: title,
		})
	})
}

func (s *Site) GetChapterContent(c *src.Chapter, reader *goquery.Document, e *src.EBook) {

	content := reader.Find("#rtext").Text()
	nextHref, _ := reader.Find("#linkNext").Attr("href")

	c.NextUrl = ""
	str1 := strings.Replace(nextHref, ".html", "", 1)
	str2 := strings.Replace(c.Url, ".html", "", 1)

	if strings.HasPrefix(str1, str2) {
		c.NextUrl = e.GetHost() + nextHref
	}

	if c.Content == "" {
		c.Content = strings.TrimSuffix(content, "\n")
	} else {
		c.Content = c.Content + strings.TrimSpace(content)
	}
}
