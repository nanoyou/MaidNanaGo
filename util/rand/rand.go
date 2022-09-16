package rand

import (
	"bytes"
	"math/rand"
	"time"
)

var table = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func init() {
	rand.Seed(time.Now().Unix())
}

// RandStr 随机字符串
func RandStr(length int) string {
	l := len(table)
	buffer := new(bytes.Buffer)
	for i := 0; i < length; i++ {
		buffer.WriteRune(table[rand.Intn(l)])
	}
	return buffer.String()
}
