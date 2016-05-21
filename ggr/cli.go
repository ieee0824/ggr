package main

import (
	"flag"
	"os"

	"github.com/spiegel-im-spiegel/ggr"
	"github.com/toqueteos/webbrowser"
)

// sample url
// https://www.google.co.jp/search?q=test&ie=utf-8&oe=utf-8&hl=ja

func main() {
	var (
		imageFlag bool
		newsFlag  bool
		shopFlag  bool
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&imageFlag, "i", false, "image flag")
	f.BoolVar(&newsFlag, "n", false, "new flag")
	f.BoolVar(&shopFlag, "s", false, "shop flag")

	f.Parse(os.Args[1:])
	q := f.Args()

	t := ggr.TypeNormal
	if imageFlag {
		t = ggr.TypeImage
	} else if newsFlag {
		t = ggr.TypeNews
	} else if shopFlag {
		t = ggr.TypeShop
	}

	g := ggr.NewGgr(ggr.LangJa, t, q)
	webbrowser.Open(g.GetSearchURL())
}
