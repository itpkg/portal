package sitemap

import (
	"encoding/xml"
	"fmt"
	"time"
)

type TimeStr string
type PriorityStr string
type ChangeFreqStr string

const (
	Always  = ChangeFreqStr("always")
	Hourly  = ChangeFreqStr("hourly")
	Daily   = ChangeFreqStr("daily")
	Weekly  = ChangeFreqStr("weekly")
	Monthly = ChangeFreqStr("monthly")
	Yearly  = ChangeFreqStr("yearly")
	Never   = ChangeFreqStr("never")
)

type UrlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Url     []*Url   `xml:"url"`
}

type Url struct {
	Loc        string        `xml:"loc"`
	LastMod    TimeStr       `xml:"lastmod"`
	ChangeFreq ChangeFreqStr `xml:"changefreq"`
	Priority   PriorityStr   `xml:"priority"`
}

func Priority(p float32) PriorityStr {
	return PriorityStr(fmt.Sprintf("%.1f", p))
}

func Time(t time.Time) TimeStr {
	return TimeStr(t.Format("2006-01-02"))
}
