package ggr

import (
	"net/url"
	"strings"
)

//Serach Types
const (
	TypeNormal = iota
	TypeImage
	TypeVideo
	TypeNews
	TypeShop
	TypeBook
	TypeBlog
	TypePatent
)

//Language
const (
	LangEn = iota
	LangJa
)

//Ggr - Make URL for Google search
type Ggr struct {
	host string //Host name
	q    string //Query keyword
	ie   string //character encoding for query
	oe   string //character encoding for result
	hl   string //language for display
	tbm  string //search type
}

//NewGgr returns Ggr instance.
func NewGgr(l int, t int, q []string) *Ggr {
	g := &Ggr{
		ie: "utf-8",
		oe: "utf-8",
	}
	g.Lang(l)
	g.SearchType(t)
	g.QueryStrings(q)
	return g
}

//Lang set language
func (p *Ggr) Lang(l int) {
	switch l {
	case LangJa:
		p.host = "www.google.co.jp"
		p.hl = "ja"
	default:
		p.host = "www.google.com"
		p.hl = "en"
	}
}

//QueryStrings set query keywords.
func (p *Ggr) QueryStrings(q []string) {
	if q != nil {
		p.q = strings.Join(q, " ")
	} else {
		p.q = ""
	}
}

//SearchType set query keywords.
func (p *Ggr) SearchType(t int) {
	switch t {
	case TypeImage:
		p.tbm = "isch"
	case TypeVideo:
		p.tbm = "vid"
	case TypeNews:
		p.tbm = "nws"
	case TypeShop:
		p.tbm = "shop"
	case TypeBook:
		p.tbm = "bks"
	case TypeBlog:
		p.tbm = "blg"
	case TypePatent:
		p.tbm = "pts"
	default:
		p.tbm = ""
	}
}

//GetSearchURL returns URL for Google search.
func (p Ggr) GetSearchURL() string {
	values := url.Values{}
	values.Add("q", p.q)
	values.Add("ie", p.ie)
	values.Add("oe", p.oe)
	values.Add("hl", p.hl)
	values.Add("tbm", p.tbm)
	u := url.URL{}
	u.Scheme = "https"
	u.Host = p.host
	u.Path = "search"
	u.RawQuery = values.Encode()
	return u.String()
}
