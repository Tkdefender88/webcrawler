package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getH1FromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		expected  string
	}{
		{
			name:      "simple html",
			inputHtml: `<html><head><title>Hello</title></head><body><h1>Hello World</h1></body></html>`,
			expected:  "Hello World",
		},
		{
			name:      "simple html with multiple h1",
			inputHtml: `<html><head><title>Hello</title></head><body><h1>Hello World</h1><h1>Hello World 2</h1></body></html>`,
			expected:  "Hello World",
		},
		{
			name:      "no h1 tag",
			inputHtml: `<html><head><title>Hello</title></head><body><p>Hello World</p></body></html>`,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h1Tag := getH1FromHTML(tt.inputHtml)
			require.Equal(t, tt.expected, h1Tag)
		})
	}
}

func Test_getFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		expected  string
	}{
		{
			name:      "simple html",
			inputHtml: `<html><head><title>Hello</title></head><body><p>Hello World</p></body></html>`,
			expected:  "Hello World",
		},
		{
			name:      "simple html with multiple p",
			inputHtml: `<html><head><title>Hello</title></head><body><p>Hello World</p><p>Hello World 2</p></body></html>`,
			expected:  "Hello World",
		},
		{
			name:      "no p tag",
			inputHtml: `<html><head><title>Hello</title></head><body><h1>Hello World</h1></body></html>`,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pTag := getFirstParagraphFromHTML(tt.inputHtml)
			require.Equal(t, tt.expected, pTag)
		})
	}
}

func Test_getURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		baseUrl   string
		expected  []string
	}{
		{
			name:      "simple html",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><a href="https://blog.boot.dev/path/to/something">Hello World</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path/to/something"},
		},
		{
			name:      "simple html with multiple a",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><a href="https://blog.boot.dev/path/to/something">Hello World</a><a href="https://blog.boot.dev/path/to/something">Hello World 2</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path/to/something", "https://blog.boot.dev/path/to/something"},
		},
		{
			name:      "no a tag",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><h1>Hello World</h1></body></html>`,
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseUrl, err := url.Parse(tt.baseUrl)
			require.NoError(t, err)
			links, err := getURLsFromHTML(tt.inputHtml, baseUrl)
			require.NoError(t, err)
			require.Equal(t, tt.expected, links)
		})
	}
}

func Test_getImageUrls(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		baseUrl   string
		expected  []string
	}{
		{
			name:      "simple html",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><img src="https://blog.boot.dev/path/to/something"/></body></html>`,
			expected:  []string{"https://blog.boot.dev/path/to/something"},
		},
		{
			name:      "simple html with multiple img",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><img src="https://blog.boot.dev/path/to/something"></img><img src="https://blog.boot.dev/path/to/something"></img></body></html>`,
			expected:  []string{"https://blog.boot.dev/path/to/something", "https://blog.boot.dev/path/to/something"},
		},
		{
			name:      "no img tag",
			baseUrl:   "https://blog.boot.dev",
			inputHtml: `<html><head><title>Hello</title></head><body><h1>Hello World</h1></body></html>`,
			expected:  []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func Test_extractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		expected  PageData
	}{
		{
			name: "simple html",
			inputHtml: `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				Images:         []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageData := extractPageData(tt.inputHtml, "https://blog.boot.dev")
			require.Equal(t, tt.expected, pageData)
		})
	}
}
