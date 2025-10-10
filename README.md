# Go Web Crawler

A simple, educational web crawler written in Go, built while following the Boot.dev course on building a web crawler. It crawls a target website, extracts links, normalizes URLs, and can generate a CSV report of discovered pages.

## Description

This project demonstrates practical web crawling concepts in Go:
- Fetching and parsing HTML
- Extracting and normalizing links
- Respecting same-origin constraints
- Tracking visited pages to avoid cycles
- Generating a CSV report of discovered URLs and their counts
- Utilizing goroutines to speed up crawling with concurrency

The codebase is intentionally small and approachable, making it a good starting point for learners. Key files include:
- main.go: CLI entrypoint
- crawler.go: Crawl logic and concurrency control
- html.go: HTML parsing utilities
- normalize_url.go: URL normalization helpers
- csv_report.go: CSV report generation

## Quick start

Prerequisites:
- Go 1.25+ installed

Cloen the repo and then:
```bash

# Build
go build -o bin/webcrawler

# Run (crawl a site up to a depth)
./bin/webcrawler https://example.com 2

# Or run directly with go run
go run . https://example.com 2
```

## Usage

Basic usage:

```bash
webcrawler <start_url> <max_goroutines> <max_pages>
```

- start_url: The root URL to begin crawling
- max_goroutines: Non-negative integer number of concurrent goroutines to use
- max_depth: Non-negative integer depth limit for the crawl

Example:

```bash
webcrawler https://boot.dev 2
```

CSV report:
- After a crawl, a report.csv file is generated in the project root, listing discovered pages and counts.
- You can open it in a spreadsheet app or process it programmatically.

Notes:
- The crawler restricts itself to the origin of start_url.
- URLs are normalized to avoid duplicates caused by fragments, trailing slashes, etc.

## Contributing

Development tips:
- Use `go test ./...` to run tests
- Keep functions small and focused

License: MIT
