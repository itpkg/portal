package seo

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/tools/blog/atom"
)

type RssHandler func(lang string) []*atom.Entry

var rssHandlers = make([]RssHandler, 0)

func RegisterRssHandler(fns ...RssHandler) {
	rssHandlers = append(rssHandlers, fns...)
}

func Rss(wrt io.Writer, lang, title, home, username, email string) error {
	feed := atom.Feed{
		Title: title,
		ID:    uuid.New(),
		Link: []atom.Link{
			{
				Href: home,
			},
		},
		Updated: atom.Time(time.Now()),
		Author: &atom.Person{
			Name:  username,
			Email: email,
		},
		Entry: make([]*atom.Entry, 0),
	}
	for _, fn := range rssHandlers {
		items := fn(lang)
		feed.Entry = append(feed.Entry, items...)
	}

	if _, err := wrt.Write([]byte(xml.Header)); err != nil {
		return err
	}
	end := xml.NewEncoder(wrt)
	return end.Encode(feed)
}
