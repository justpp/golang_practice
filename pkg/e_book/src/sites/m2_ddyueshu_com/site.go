package m2_ddyueshu_com

import (
	"giao/pkg/e_book/src"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type Site struct {
}

func (s *Site) GetHost() string {
	return "www.ddyueshu.com"
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
	chapterList := reader.Find("#list")

	chapterList.Find("a").Each(func(i int, selection *goquery.Selection) {
		if i < 6 {
			return
		}
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

	content, _ := reader.Find("#content").Html()
	compile := regexp.MustCompile(`<br(/)*>`)
	content = compile.ReplaceAllString(content, "\n")

	compile = regexp.MustCompile(`请记住本书首发域名：ddyueshu.com。顶点小说手机版阅读网址：m.ddyueshu.com`)
	content = compile.ReplaceAllString(content, "")

	c.NextUrl = ""

	if c.Content == "" {
		c.Content = strings.TrimSuffix(content, "\n")
	} else {
		c.Content = c.Content + strings.TrimSpace(content)
	}
}
