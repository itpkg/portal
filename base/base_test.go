package base_test

import (
	"testing"

	"github.com/itpkg/portal/base"
)

func TestLogo(t *testing.T) {
	u := base.User{Email: "MyEmailAddress@example.com"}
	u.SetLogoByGravatar()
	v := "https://gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346.png"
	if u.Logo != v {
		t.Errorf("Wang %s, get %s", v, u.Logo)
	}

}
