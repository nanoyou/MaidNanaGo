package request_test

import (
	"testing"

	"github.com/nanoyou/MaidNanaGo/controller/request"
	"github.com/nanoyou/MaidNanaGo/validator"
)

var v = validator.Get()

func assertStruct(t *testing.T, s any) {
	if err := v.Struct(s); err != nil {
		t.Error(s)
	}
}
func assertStructFail(t *testing.T, s any) {
	if err := v.Struct(s); err == nil {
		t.Error(s)
	}
}
func TestRegisterRequest(t *testing.T) {
	req := request.RegisterRequest{}
	// 空
	assertStructFail(t, req)
	req.Password = "123456789"
	req.VerificationCode = 114514
	// 不足6位
	req.Username = "aaaaa"
	assertStructFail(t, req)
	// 6位
	req.Username = "aaaaaa"
	assertStruct(t, req)
	// 20位
	req.Username = "01234567890123456789"
	assertStruct(t, req)
	// 超过20位
	req.Username = "012345678901234567890"
	assertStructFail(t, req)
	// 允许的字符
	req.Username = "abc123ABC-_"
	assertStruct(t, req)
	// 不允许的字符
	req.Username = "啊啊啊aaaaaa"
	assertStructFail(t, req)
	req.Username = "__12345a."
	assertStructFail(t, req)
	req.Username = "__12345a*"
	assertStructFail(t, req)
}
