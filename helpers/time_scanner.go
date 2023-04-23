package helpers

import "time"

type TimeScanner struct {
	time.Time
}

func (t *TimeScanner) Scan(value interface{}) error {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(value.([]byte)))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
