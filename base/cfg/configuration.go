package cfg

import (
	"github.com/BurntSushi/toml"
)

func Load(file, env string) (interface{}, error) {
	items := make(map[string]interface{})
	if _, err := toml.DecodeFile(file, items); err == nil {
		return items[env], nil
	} else {
		return nil, err
	}
}
