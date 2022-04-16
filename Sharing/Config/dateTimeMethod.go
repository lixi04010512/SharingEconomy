package config

import (
	"time"
)

//时间戳转日期
func UnixToTime(timestamp int) string {
	t :=time.Unix(int64(timestamp),0)
	return t.Format("2022-04-12 18:06:05")
}

//日期转时间戳
func DateToUnix(str string) int64 {
	templates := "2022-04-12 18:06:05"
	t,err :=time.ParseInLocation(templates,str,time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

//获取当前日期
func GetDate() string {
	templates := "2022-04-12 18:06:05"
	return time.Now().Format(templates)
}

//获取年月日
func GetDay() string {
	templates := "20220412"
	return time.Now().Format(templates)
}
