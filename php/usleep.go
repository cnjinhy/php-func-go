package php

import "time"

func Usleep(microseconds int) {
	time.Sleep(time.Duration(microseconds) * time.Microsecond)
}
