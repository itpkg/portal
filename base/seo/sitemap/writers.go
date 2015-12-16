package sitemap

import (
	"compress/gzip"
	"encoding/xml"
	"io"
)

func Dump(wrt io.Writer, us *UrlSet, zip bool) error {

	if zip {
		zip := gzip.NewWriter(wrt)
		defer zip.Close()
		wrt = zip
	}
	wrt.Write([]byte(xml.Header))
	en := xml.NewEncoder(wrt)
	return en.Encode(us)

}
