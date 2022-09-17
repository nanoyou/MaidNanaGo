package bytes_test

import (
	"testing"

	"github.com/nanoyou/MaidNanaGo/util/bytes"
)

func TestBytes(t *testing.T) {
	var number int64 = 941974425
	arr := bytes.Int64ToBytes(number)
	if bytes.BytesToInt64(arr) != number {
		t.Error("不相等")
	}
}
