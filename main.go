package main

import(
	"flag"
	"net/url"
	"strings"

	"github.com/toqueteos/webbrowser"
)

// sample url
// https://www.google.co.jp/search?q=test&ie=utf-8&oe=utf-8&hl=ja

type Ggr struct {
    q string
    ie string
    oe string
    hl string
}

func NewGgr() Ggr{
    g := Ggr{
        q : "",
        ie : "utf-8",
        oe : "utf-8",
        hl : "ja",
        tbm : ""
    }
    return g
}

func (p Ggr) getSearchURL() string {
	values := url.Values{}
	values.Add("q", p.q)
	values.Add("ie", p.ie)
	values.Add("oe", p.oe)
	values.Add("hl", p.hl)
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "www.google.co.jp"
	u.Path = "search"
	u.RawQuery = values.Encode()
    return u.String()
}

func main() {
	flag.Parse()
	args := flag.Args()
    g := NewGgr()
    g.q = strings.Join(args, " ")
	webbrowser.Open(g.getSearchURL())
}
