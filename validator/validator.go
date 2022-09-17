package validator

import "github.com/go-playground/validator/v10"

var v = validator.New()

func init() {
	// URL 安全的名字
	v.RegisterValidation("urlsafename", func(fl validator.FieldLevel) bool {
		str := fl.Field().String()
		for _, ch := range str {
			if ch >= 'a' && ch <= 'z' {
				continue
			}
			if ch >= 'A' && ch <= 'Z' {
				continue
			}
			if ch >= '0' && ch <= '9' {
				continue
			}
			if ch == '-' || ch == '_' {
				continue
			}
			return false
		}
		return true
	})
}

func Get() *validator.Validate {
	return v
}
