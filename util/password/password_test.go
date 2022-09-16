package password_test

import (
	"testing"

	"github.com/nanoyou/MaidNanaGo/util/password"
)

func TestPlain(t *testing.T) {
	passwordA := "abcdefg"
	ppA := password.NewPlainPassword(passwordA)
	t.Log(ppA.String())

	if !ppA.Validate(passwordA) {
		t.Error("相同密码校验不通过")
	}

	if ppA.Validate(passwordA + "A") {
		t.Error("不同密码校验通过")
	}
}

func TestSHA512(t *testing.T) {
	passwordA := "abcdefg"
	ppA := password.NewSHA512Password(passwordA)
	t.Log(ppA.String())

	if !ppA.Validate(passwordA) {
		t.Error("相同密码校验不通过")
	}

	if ppA.Validate(passwordA + "A") {
		t.Error("不同密码校验通过")
	}

}
