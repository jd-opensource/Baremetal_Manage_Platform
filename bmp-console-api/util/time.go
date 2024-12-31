package util

import (
	"time"
)

//ParseTimeToDate 解析time格式为date格式
func ParseTimeToDate(str, timeLayout, dateLayout string) (time.Time, error) {
	var res time.Time

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return res, err
	}
	res, err = time.Parse(dateLayout, t.Format(dateLayout))
	if err != nil {
		return res, err
	}
	return res, nil
}

//GetFirstDayOfMonth 获取某个月份第一天
func GetFirstDayOfMonth(month, monthLayout string) (time.Time, error) {
	var res time.Time

	t, err := time.Parse(monthLayout, month)
	if err != nil {
		return res, err
	}
	return t.AddDate(0, 0, -t.Day()+1), nil
}

//GetLastDayOfMonth 获取某个月份最后一天
func GetLastDayOfMonth(month, monthLayout string) (time.Time, error) {
	var res time.Time

	t, err := time.Parse(monthLayout, month)
	if err != nil {
		return res, err
	}
	return t.AddDate(0, 1, -t.Day()), nil
}

//GetDateSubDays 获取某两个时间的差值
func GetDateSubDays(dateStart, dateEnd time.Time) int {
	return int(dateEnd.Sub(dateStart).Hours()/24) + 1
}

//GetMonthDays 获取某个月份总天数
func GetMonthDays(month, monthLayout string) (int, error) {
	dateStart, err := GetFirstDayOfMonth(month, monthLayout)
	if err != nil {
		return 0, err
	}
	dateEnd, err := GetLastDayOfMonth(month, monthLayout)
	if err != nil {
		return 0, err
	}

	return GetDateSubDays(dateStart, dateEnd), nil
}

//GetMonthCount 获取某两个月份之间的个数,包含这2个月份
//ex: 2021-07和2021-09之间包含3个月
func GetMonthCount(startMonth, endMonth, monthLayout string) (int, error) {
	startMonthTime, err := time.Parse(monthLayout, startMonth)
	if err != nil {
		return 0, err
	}
	endMonthTime, err := time.Parse(monthLayout, endMonth)
	if err != nil {
		return 0, err
	}

	startYear := startMonthTime.Year()
	endYear := endMonthTime.Year()
	startM := int(startMonthTime.Month())
	endM := int(endMonthTime.Month())
	yearInterval := endYear - startYear

	if endM <= startM {
		yearInterval--
	}

	monthInterval := (endM + 12 - startM) % 12
	return yearInterval*12 + monthInterval + 1, nil
}

//GetMonthSub 获取某两个月份之间的月份，包含这2个月份
func GetMonthSub(startMonth, endMonth, monthLayout string) ([]string, error) {
	months, err := GetMonthCount(startMonth, endMonth, monthLayout)
	if err != nil {
		return nil, err
	}
	startMonthTime, err := time.Parse(monthLayout, startMonth)
	if err != nil {
		return nil, err
	}
	monthArr := make([]string, 0)
	for i := 0; i < months; i++ {
		monthArr = append(monthArr, startMonthTime.AddDate(0, i, 0).Format(monthLayout))
	}
	return monthArr, nil
}
