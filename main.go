package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	if len(args) < 2 {
		return fmt.Errorf("no website provided")
	}
	if len(args) > 2 {
		return fmt.Errorf("too many arguments provided")
	}

	url := args[1]

	fmt.Fprintf(out, "starting crawl of: %s\n", url)
	pages := make(map[string]int)
	crawlPage(url, url, pages)

	for key, value := range pages {
		fmt.Printf("page %s referenced %d times\n", key, value)
	}

	return nil
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

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if currentUrl.Host != baseUrl.Host {
		//fmt.Printf("not the same host, exiting. %s : %s\n", currentUrl.Host, baseUrl.Host)
		return
	}

	normUrl, err := NormalizeURL(currentUrl.String())
	if err != nil {
		fmt.Println("could not normalize url")
		return
	}

	count, ok := pages[normUrl]
	if !ok {
		pages[normUrl] = 1
	} else {
		pages[normUrl] = count + 1
		return
	}

	fmt.Printf("Crawling: %s\n", normUrl)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		//fmt.Printf("error fetching html for the page: %v\n", err)
		return
	}

	pageData := extractPageData(html, rawCurrentURL)

	for _, url := range pageData.OutgoingLinks {
		crawlPage(rawBaseURL, url, pages)
	}
}
