package seo

import (
	"io"

	"github.com/itpkg/portal/base/seo/sitemap"
)

type SitemapHandler func() []*sitemap.Url

var sitemapHandlers = make([]SitemapHandler, 0)

func RegisterSitemap(fns ...SitemapHandler) {
	sitemapHandlers = append(sitemapHandlers, fns...)
}

func Sitemap(wrt io.Writer) error {
	us := sitemap.UrlSet{Url: make([]*sitemap.Url, 0)}
	for _, f := range sitemapHandlers {
		items := f()
		us.Url = append(us.Url, items...)
	}
	return sitemap.Dump(wrt, &us, true)
}
