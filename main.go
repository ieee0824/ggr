package main

import (
	"flag"
	"net/url"
	"strings"

	"github.com/toqueteos/webbrowser"
)

// sample url
// https://www.google.co.jp/search?q=test&ie=utf-8&oe=utf-8&hl=ja

func main() {
	flag.Parse()
	values := url.Values{}
	args := flag.Args()
	q := strings.Join(args, " ")
	values.Add("q", q)
	values.Add("ie", "utf-8")
	values.Add("oe", "utf-8")
	values.Add("hl", "ja")
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "www.google.co.jp"
	u.Path = "search"
	u.RawQuery = values.Encode()
	webbrowser.Open(u.String())
}
