package seo

/*
Sitemap: https://CHANGE-ME/sitemap.xml.gz
User-agent:*
Disallow:
*/

func Robots(txt string) (name string, body string) {
	return "robots.txt", txt
}
