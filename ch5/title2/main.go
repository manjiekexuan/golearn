package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s,not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	fmt.Println(doc)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML:%v", url, err)
	}
	return nil
}
