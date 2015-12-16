package seo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/itpkg/portal/base/seo"
	"golang.org/x/tools/blog/atom"
)

func TestRss(t *testing.T) {
	seo.RegisterRssHandler(func(string) []*atom.Entry {
		items := make([]*atom.Entry, 0)
		for i := 0; i < 10; i++ {
			items = append(items, &atom.Entry{
				ID: fmt.Sprintf("article_%d", i),
				Link: []atom.Link{
					{
						Href: fmt.Sprintf("http://aaa.com/%d", i),
					},
				},
				Title: fmt.Sprintf("Title %d", i),
				Summary: &atom.Text{
					Body: fmt.Sprintf("Summary %d", i),
				},
				Content: &atom.Text{
					Body: fmt.Sprintf("Content %d", i),
				},
			})
		}
		return items
	})
	if err := seo.Rss(os.Stdout, "zh-CN", "ttttttt", "http://aaa.com", "whoami", "whoami@gmail.com"); err != nil {
		t.Errorf("bad in rss: %v", err)
	}
}

func TestSeo(t *testing.T) {
	bn, bb := seo.BaiduVerify("bbbbbbb")
	t.Logf("Baidu verify file: %s \n %s", bn, bb)
	gn, gb := seo.GoogleVerify("ggggg")
	t.Logf("Google verify file: %s \n %s", gn, gb)
	rn, rb := seo.Robots("sitemap: http://aaa.com/sitemap.xml.gz")
	t.Logf("robots.txt: %s \n %s", rn, rb)
}
