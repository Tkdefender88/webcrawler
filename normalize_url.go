package main

import (
	"fmt"
	"net/url"
)

func NormalizeURL(in string) (string, error) {

	url, err := url.Parse(in)
	if err != nil {
		return "", err
	}

	normailizedUrl := fmt.Sprintf("%s%s", url.Host, url.Path)

	return normailizedUrl, nil
}
