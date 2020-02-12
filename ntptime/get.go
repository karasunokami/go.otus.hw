package ntptime

import (
	"github.com/beevik/ntp"
	"time"
)

func Get() (time.Time, time.Time, error) {
	ntpTime, err := ntp.Time("0.ru.pool.ntptime.org")
	now := time.Now()

	if err != nil {
		return now, ntpTime, err
	}

	return now, ntpTime, err
}
