package interfaces

import (
	"time"
)

type CustomDateTime time.Time

type CustomDate time.Time

func (c *CustomDateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	*c = CustomDateTime(t)
	return nil
}

func (c CustomDateTime) ToTime() time.Time {

	t := time.Time(c)
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	*c = CustomDate(t)
	return nil
}

func (c CustomDate) ToTime() time.Time {
	t := time.Time(c)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
