package util

import (
	"errors"
	"math"
	"strconv"
	"time"
)

// 获取当前时间
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 指定格式当前时间
func NowFormat(format string) string {
	return time.Now().Format(format)
}

// 获取X天前
func XBeforeDay(days int) string {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -days) //若要获取x天前的时间
	return oldTime.Format("2006-01-02 15:04:05")
}

func XBeforeDayFormat(days int, format string) string {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -days) //若要获取x天前的时间
	return oldTime.Format(format)
}

// 获取相差时间
func GetMinDiffer(startTime, endTime string) int64 {
	var minute int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		minute = diff / 60
		return minute
	} else {
		return minute
	}
}

// 获取相差时间
func GetMinAndSecDiffer(startTime, endTime string) (int64, int64) {
	var minute int64
	var second int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		minute = diff / 60
		second = diff % 60
		return minute, second
	} else {
		return minute, second
	}
}

func IsToday(t string) bool {
	format := "2006-01-02 15:04:05"
	sqlUpdatedAt, _ := time.ParseInLocation(format, t, time.Local)

	tNow := time.Now()
	tZero := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, tNow.Location())
	tTw := time.Date(tNow.Year(), tNow.Month(), tNow.Day()+1, 0, 0, 0, 0, tNow.Location())

	subZ := sqlUpdatedAt.Sub(tZero)
	subT := sqlUpdatedAt.Sub(tTw)
	if subZ >= 0 && subT < 0 {
		return true
	} else {
		return false
	}
}

func IsYesterday(t string) bool {
	format := "2006-01-02 15:04:05"
	sqlUpdatedAt, _ := time.ParseInLocation(format, t, time.Local)

	tNow := time.Now()
	tZero := time.Date(tNow.Year(), tNow.Month(), tNow.Day()-1, 0, 0, 0, 0, tNow.Location())
	tTw := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, tNow.Location())

	subZ := sqlUpdatedAt.Sub(tZero)
	subT := sqlUpdatedAt.Sub(tTw)
	if subZ >= 0 && subT < 0 {
		return true
	} else {
		return false
	}
}

func IsThisYear(t string) bool {
	format := "2006-01-02 15:04:05"
	sqlUpdatedAt, _ := time.ParseInLocation(format, t, time.Local)

	tNow := time.Now()
	tZero := time.Date(tNow.Year(), 1, 1, 0, 0, 0, 0, tNow.Location())
	tTw := time.Date(tNow.Year()+1, 1, 1, 0, 0, 0, 0, tNow.Location())

	subZ := sqlUpdatedAt.Sub(tZero)
	subT := sqlUpdatedAt.Sub(tTw)
	if subZ >= 0 && subT < 0 {
		return true
	} else {
		return false
	}
}

func IsBeforeYear(t string) bool {
	format := "2006-01-02 15:04:05"
	sqlUpdatedAt, _ := time.ParseInLocation(format, t, time.Local)

	tNow := time.Now()
	tZero := time.Date(tNow.Year(), 1, 1, 0, 0, 0, 0, tNow.Location())

	subZ := sqlUpdatedAt.Sub(tZero)
	if subZ < 0 {
		return true
	} else {
		return false
	}
}

// utc时间转标准时间
func UtcTimeToStandardTime(utcTime string) (string, error) {
	if utcTime == "" {
		return "", errors.New("param is empty")
	}

	t, _ := time.Parse(time.RFC3339, utcTime)

	return t.Local().Format("2006-01-02 15:04:05"), nil
}

// 秒级时间转HH:MM:SS
func SecondsToDate(seconds int) string {
	hour  := int(math.Floor(float64(seconds/ 3600)))
	minute := int(math.Floor(float64(seconds % 3600 / 60)))
	second := int(math.Floor(float64(seconds % 3600 % 60)))

	return addZero(hour) + ":" + addZero(minute) + ":" + addZero(second)
}
// 补0
func addZero(num int) string {
	var result string
	if num < 10 {
		result = "0" + strconv.FormatInt(int64(num), 10)
	} else {
		result = strconv.FormatInt(int64(num), 10)
	}
	return result
}
