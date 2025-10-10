package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("no website provided")
	}
	if len(args) > 4 {
		return fmt.Errorf("too many arguments provided")
	}

	baseUrl := args[1]
	maxConcurrency := 5
	maxPages := 10
	if len(args) > 2 && args[2] != "" {
		var err error
		maxConcurrency, err = strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("could not parse max concurrency: %w", err)
		}
		if maxConcurrency < 1 {
			return fmt.Errorf("max concurrency must be greater than 0")
		}
	}

	if len(args) > 3 && args[3] != "" {
		var err error
		maxPages, err = strconv.Atoi(args[3])
		if err != nil {
			return fmt.Errorf("could not parse max pages: %w", err)
		}
		if maxPages < 1 {
			return fmt.Errorf("max pages must be greater than 0")
		}
	}

	fmt.Printf("Starting Crawl of %s with max concurrency %d and max pages %d\n", baseUrl, maxConcurrency, maxPages)

	cfg, err := configure(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		return err
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	for page := range cfg.pages {
		fmt.Printf("%s\n", page)
	}

	return nil
}
