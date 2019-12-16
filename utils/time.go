package utils

import (
	"time"
)

type Time struct {
	time.Time
}

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	tt, err := time.Parse(timeFormat, string(b))
	if err != nil {
		return err
	}

	*t = Time{
		Time: tt,
	}

	return nil
}
