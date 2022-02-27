package helper

import (
	"billing-cycle-ms-app/constant"
	"time"
)

func HelperDate(ds string) (time.Time, error) {
	t, err := time.Parse(constant.LayoutISO, ds)
	if err != nil {
		return t, err
	}
	return t, nil
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
