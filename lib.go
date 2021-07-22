package libs

import (
	"os"
	"path/filepath"
	"time"
	"unsafe"
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

func BytesToBits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	return dst
}

func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func IntToBits(num int64) []int {
	return BytesToBits(IntToByteArray(num))
}
