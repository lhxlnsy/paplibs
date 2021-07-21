package libs

import (
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
