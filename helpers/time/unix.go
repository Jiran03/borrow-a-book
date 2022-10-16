package time

import (
	"strconv"
	"time"
)

func NanoToTime(unix string) time.Time {
	i, _ := strconv.Atoi(unix)
	time := time.UnixMilli(int64(i))
	return time
}

func Timestamp() string {
	time := strconv.Itoa(int(time.Now().UnixNano() / 1000000))
	return time
}

func DayToNano(days int) string {
	days *= 86400000
	time := strconv.Itoa(int(time.Now().UnixNano()/1000000) + days)
	return time
}
