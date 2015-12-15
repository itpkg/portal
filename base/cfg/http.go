package cfg

import (
	"github.com/itpkg/portal/base/utils"
)

type Http struct {
	Domain  string `toml:"domain"`
	Port    int    `toml:"port"`
	Secrets string `toml:"secret"`
}

func (p *Http) Key(begin, end int) ([]byte, error) {
	if b, e := utils.String2Bytes(p.Secrets); e == nil {
		return b[begin:end], nil
	} else {
		return nil, e
	}

}
