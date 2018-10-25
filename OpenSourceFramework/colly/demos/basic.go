package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func main() {

	doFetchBilibili()

}

func bisicOfficDemo() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	c.AllowedDomains = []string{"hackerspaces.org", "wiki.hackerspaces.org"}

	//过滤A标签，对其处理回调
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}

func doFetchBilibili() {
	spider := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36 Core/1.47.933.400 QQBrowser/9.4.8699.400"),
		colly.MaxDepth(1),
	)
	//spider.AllowedDomains = []string{"bilibili.com"}
	spider.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.DOM.Children().Text()
		fmt.Printf("Link found: %q -> %s\n", title, link)
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	spider.OnRequest(func(req *colly.Request) {
		fmt.Println("visiting:", req.URL.String())
	})

	time.Sleep(time.Millisecond * 100)
	spider.Visit("https://www.bilibili.com/v/dance/?spm_id_from=333.334.primary_menu.29")
}
