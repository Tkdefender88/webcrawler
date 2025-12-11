package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	csvWriter.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})

	for url, page := range pages {
		urls := strings.Join(page.OutgoingLinks, ";")
		images := strings.Join(page.Images, ";")
		csvWriter.Write([]string{url, page.H1, page.FirstParagraph, urls, images})
	}

	return nil
}
