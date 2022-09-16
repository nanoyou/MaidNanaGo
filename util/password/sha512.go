package password

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/nanoyou/MaidNanaGo/util/rand"
)

type SHA512Password struct {
	password string
	salt     string
}

func NewSHA512Password(password string) *SHA512Password {
	r := &SHA512Password{}
	r.salt = rand.RandStr(32)
	r.password = hashSHA512(password, r.salt)
	return r
}

func (sp *SHA512Password) Validate(password string) bool {
	return hashSHA512(password, sp.salt) == sp.password
}

func (sp *SHA512Password) String() string {
	return "SHA-512:" + sp.salt + ":" + sp.password
}

func hashSHA512(password string, salt string) string {
	hash := sha512.New()
	result := hash.Sum([]byte(password + salt))
	return hex.EncodeToString(result)
}
