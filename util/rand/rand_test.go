package rand_test

import (
	"testing"

	"github.com/nanoyou/MaidNanaGo/util/rand"
)

func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if rand.RandStr(16) == rand.RandStr(16) {
			t.Error("随机字符串相同")
		}
	}
}
