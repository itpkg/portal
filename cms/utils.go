package cms

import (
	"bufio"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func Md2Hm(input []byte) []byte {
	return bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon(input))
}

func FirstLine(file string) string {
	f, e := os.Open(file)
	if e != nil {
		return ""
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	return s.Text()
}
