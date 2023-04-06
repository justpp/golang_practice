package timer

import "time"

const (
	Ymd    = "2006-01-02"
	YmdHi  = "2006-01-02 15:04"
	YmdHis = "2006-01-02 15:04:05"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTime.Add(duration), nil
}
