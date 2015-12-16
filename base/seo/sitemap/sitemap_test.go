package sitemap_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	sm "github.com/itpkg/portal/base/seo/sitemap"
)

func TestSitemap(t *testing.T) {
	us := sm.UrlSet{Url: make([]*sm.Url, 0)}
	for i := 0; i < 10; i++ {
		us.Url = append(us.Url, &sm.Url{
			Loc:        fmt.Sprintf("http://aaa.com/%d", i),
			ChangeFreq: sm.Daily,
			LastMod:    sm.Time(time.Now()),
			Priority:   sm.Priority(0.67),
		})
	}
	if err := sm.Dump(os.Stdout, &us, false); err != nil {
		t.Errorf("bad in sitemap: %v", err)
	}
}
