package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]PageData
	maxPages           int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(rawBaseUrl string, maxConcurrency, maxPages int) (*config, error) {
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, err
	}

	cfg := &config{
		pages:              make(map[string]PageData),
		maxPages:           maxPages,
		baseUrl:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	return cfg, nil
}

func (this *config) addPageVisit(normalizedUrl string) (isFirst bool) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, visited := this.pages[normalizedUrl]; visited {
		return false
	}

	this.pages[normalizedUrl] = PageData{URL: normalizedUrl}
	return true
}

func (this *config) checkPageDepth() (isMaxDepth bool) {
	this.mu.Lock()
	defer this.mu.Unlock()

	return len(this.pages) >= this.maxPages
}

func (this *config) setPageData(normalizedUrl string, data PageData) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.pages[normalizedUrl] = data
}

func (this *config) crawlPage(rawCurrentURL string) {
	this.concurrencyControl <- struct{}{}
	defer func() {
		<-this.concurrencyControl
		this.wg.Done()
	}()

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if currentUrl.Hostname() != this.baseUrl.Hostname() {
		return
	}

	normUrl, err := NormalizeURL(currentUrl.String())
	if err != nil {
		fmt.Println("could not normalize url")
		return
	}

	firstVisit := this.addPageVisit(normUrl)
	if !firstVisit {
		return
	}

	if this.checkPageDepth() {
		return
	}

	fmt.Printf("Crawling: %s\n", normUrl)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		//fmt.Printf("error fetching html for the page: %v\n", err)
		return
	}

	pageData := extractPageData(html, rawCurrentURL)
	this.setPageData(normUrl, pageData)

	for _, url := range pageData.OutgoingLinks {
		this.wg.Add(1)
		go this.crawlPage(url)
	}
}

func getHTML(rawURL string) (string, error) {

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	req.Header.Add("User-Agent", "BootCrawler/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("problem making request: %w", err)
	}
	if resp.StatusCode > 400 {
		return "", fmt.Errorf("request failed with response: %s", resp.Status)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("page didn't respond with html")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(data), nil
}
