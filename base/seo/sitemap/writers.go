package sitemap

import (
	"compress/gzip"
	"encoding/xml"
	"io"
)

func Dump(wrt io.Writer, us *UrlSet, zip bool) error {
	wrt.Write([]byte(xml.Header))
	if zip {
		zip := gzip.NewWriter(wrt)
		defer zip.Close()
		wrt = zip
	}
	en := xml.NewEncoder(wrt)
	return en.Encode(us)
}
