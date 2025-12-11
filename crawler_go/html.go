package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	Images         []string
}

func extractPageData(html, pageUrl string) PageData {
	title := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	parsedUrl, err := url.Parse(pageUrl)
	if err != nil {
		return PageData{
			URL:            pageUrl,
			H1:             title,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			Images:         nil,
		}
	}

	links, err := getURLsFromHTML(html, parsedUrl)
	if err != nil {
		links = nil
	}

	images, err := getImagesFromHTML(html, parsedUrl)
	if err != nil {
		images = nil
	}

	return PageData{
		URL:            pageUrl,
		H1:             title,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  links,
		Images:         images,
	}
}

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return doc.Find("h1").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return doc.Find("p").First().Text()
}

func getURLsFromHTML(html string, baseUrl *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	links := make([]string, 0)

	doc.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if !exists {
			return
		}

		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		u, err := url.Parse(href)
		if err != nil {
			return
		}

		absolute := baseUrl.ResolveReference(u)
		links = append(links, absolute.String())
	})

	return links, nil
}

func getImagesFromHTML(html string, baseUrl *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	images := make([]string, 0)
	doc.Find("img").Each(func(_ int, selection *goquery.Selection) {
		src, exists := selection.Attr("src")
		if !exists {
			return
		}
		u, err := url.Parse(src)
		if err != nil {
			return
		}
		absolute := baseUrl.ResolveReference(u)
		images = append(images, absolute.String())
	})

	return images, nil
}
