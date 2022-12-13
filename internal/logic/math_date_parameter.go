package logic

import (
	"git.ickey.com.cn/infra/utils/log"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const DATE_LAYOUT = "2006-01-02"

func CheckDate(s, e string) error {
	startTime, _err1 := time.Parse(DATE_LAYOUT, s)
	endTime, _err2 := time.Parse(DATE_LAYOUT, e)
	if _err1 != nil || _err2 != nil {
		_err := log.Error("error date parameters")
		return _err
	} else if startTime.After(endTime) {
		_err := log.Error("error date parameters")
		return _err
	} else if (endTime.Sub(startTime).Hours() / 24) > 31 {
		_err := log.Error("error date parameters")
		return _err
	} else {
		return nil
	}
}

func Local2UTC(s, e string) (sTime, eTime string, err error) {
	startTime, _ := time.Parse(DATE_LAYOUT, s)
	endTime, _ := time.Parse(DATE_LAYOUT, e)
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return "", "", err
	}
	startTime = startTime.In(loc)
	endTime = endTime.In(loc)
	return startTime.Format(DATE_LAYOUT), endTime.Format(DATE_LAYOUT), nil
}
