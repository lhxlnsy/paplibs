package libs

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"time"
)

const (
	TIMEFORMAT = "2006-01-02 15:04:05:000"
)

func CWD() string {
	if os.Getenv("TRAVIS_BUILD_DIR") != "" {
		return os.Getenv("TRAVIS_BUILD_DIR")
	}

	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

func StrToTime(timestr string) time.Time {
	timestamp, err := time.Parse(TIMEFORMAT, timestr)
	if err != nil {
		return time.Time{}
	}
	return timestamp
}

func Int64ToBytes(num int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(num))
	return buf
}
func BytesToBits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 7; i > 0; i-- {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	return dst
}

func IntToBits(num int64) []int {
	return BytesToBits(Int64ToBytes(num))
}
