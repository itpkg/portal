package cfg

import (
	"github.com/itpkg/portal/base/utils"
)

func Load(file, env string) (interface{}, error) {
	items := make(map[string]interface{})
	if err := utils.FromToml(file, items); err == nil {
		return items[env], nil
	} else {
		return nil, err
	}
}
