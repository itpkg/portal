package cfg

import (
	"fmt"

	"github.com/itpkg/portal/base/utils"
)

type Http struct {
	Env     string `toml:"-"`
	Domain  string `toml:"domain"`
	Port    int    `toml:"port"`
	Secrets string `toml:"secret"`
}

func (p *Http) Home() string {
	if p.IsProduction() {
		return fmt.Sprintf("https://www.%s", p.Domain)
	} else {
		return "http://localhost:8080"
	}
}

func (p *Http) Key(begin, end int) ([]byte, error) {
	if b, e := utils.FromBase64(p.Secrets); e == nil {
		return b[begin:end], nil
	} else {
		return nil, e
	}

}

func (p *Http) IsProduction() bool {
	return p.Env == "production"
}
