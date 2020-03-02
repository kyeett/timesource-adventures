package main

import (
	"fmt"
	"strings"
	"time"
)

type TimeSource interface {
	Now() time.Time
}

var _ TimeSource = (*SmarterDummy)(nil)

type SmarterDummy struct {
	DummyTime time.Time
}

var _ TimeSource = (*SmarterDummy)(nil)

// Now returns the time value this instance contains.
func (f *SmarterDummy) Now() time.Time {
	return f.DummyTime
}

func (f *SmarterDummy) AddDate(years int, months int, days int) {
	f.DummyTime = f.DummyTime.AddDate(years, months, days)
}

//https://docs.killbill.io/0.16/faq.html
//"http://127.0.0.1:8080/1.0/kb/test/clock?requestedDate=2015-12-14T23:02:15.000Z"
//"http://127.0.0.1:8080/1.0/kb/test/clock?days=10"

func main() {
	var t TimeSource
	now := time.Now()

	dummyTimeSource := &SmarterDummy{
		DummyTime: now,
	}
	t = dummyTimeSource

	fmt.Println(t.Now())
	time.Sleep(10 * time.Millisecond)
	fmt.Println(t.Now())

	updateTimeSource(dummyTimeSource, "requestedDate=2015-12-14T23:02:15.000Z")

	dummyTimeSource.AddDate(0, -1, 0)
	fmt.Println(t.Now())
}

func updateTimeSource(t TimeSource, request string) error {
	d := t.(*SmarterDummy)

	splat := strings.Split(request, "=")

	t2, err := time.Parse(time.RFC3339Nano, splat[1])
	if err != nil {
		return err
	}

	d.DummyTime = t2
	fmt.Println(t.Now())
	return nil
}
