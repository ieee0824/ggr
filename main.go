package main

import(
	"flag"
	"net/url"
	"strings"
    "os"

	"github.com/toqueteos/webbrowser"
)

// sample url
// https://www.google.co.jp/search?q=test&ie=utf-8&oe=utf-8&hl=ja

type Ggr struct {
    q string
    ie string
    oe string
    hl string
    tbm string
}

func NewGgr() Ggr{
    g := Ggr{
        q : "",
        ie : "utf-8",
        oe : "utf-8",
        hl : "ja",
        tbm : "",
    }
    return g
}

func (p Ggr) getSearchURL() string {
	values := url.Values{}
	values.Add("q", p.q)
	values.Add("ie", p.ie)
	values.Add("oe", p.oe)
	values.Add("hl", p.hl)
    values.Add("tbm", p.tbm)
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "www.google.co.jp"
	u.Path = "search"
	u.RawQuery = values.Encode()
    return u.String()
}

func main() {
    var (
        imageFlag bool
        newsFlag bool
        shopFlag bool
    )
    
    f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
    f.BoolVar(&imageFlag, "i", false, "image flag")
    f.BoolVar(&newsFlag, "n", false, "new flag")
    f.BoolVar(&shopFlag, "s", false, "shop flag")
    
	f.Parse(os.Args[1:])
	args := f.Args()
    
    g := NewGgr()
    g.q = strings.Join(args, " ")
    if imageFlag {
        g.tbm = "isch"
    } else if newsFlag {
        g.tbm = "nws"
    } else if shopFlag {
        g.tbm = "shop"
    }
	webbrowser.Open(g.getSearchURL())
}
