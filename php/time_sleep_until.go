package php

import (
	"time"
)

func TimeSleepUntil(timestamp int64) {
	t := time.Unix(timestamp, 0)
	sleepUntil(t)
}

func sleepUntil(t time.Time) {
	duration := t.Sub(time.Now())
	if duration > 0 {
		time.Sleep(duration)
	}
}
