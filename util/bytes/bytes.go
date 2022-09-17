package bytes

import "encoding/binary"

func Int64ToBytes(num int64) []byte {
	return binary.LittleEndian.AppendUint64([]byte{}, uint64(num))
}

func BytesToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
