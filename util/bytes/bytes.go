package bytes

import "encoding/binary"

func Int64ToBytes(num int64) []byte {
	r := []byte{}
	binary.LittleEndian.PutUint64(r, uint64(num))
	return r
}

func BytesToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
