package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gudgl/job-workers"
)

const (
	keyTimeString = "time_string"
	layout        = "2006-01-02"
)

var (
	timeStrings = []string{
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12 13:23:45",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
	}
)

func main() {
	c := &collector{}
	client, err := jw.New(c, 0, 0)
	if err != nil {
		logrus.Fatal("Failed to create client.")
		return
	}

	client.Go()

	for _, ts := range timeStrings {
		client.SendJob(&job{
			timeString: ts,
		})
	}

	client.Wait()
}

type job struct {
	timeString string
}

func (j *job) Execute() *jw.Result {
	time, err := time.Parse(layout, j.timeString)
	if err != nil {
		return jw.WithError(err).WithKey(keyTimeString, j.timeString)
	}

	logrus.WithFields(logrus.Fields{
		"time_string": j.timeString,
		"time":        time,
	}).Info("Time parsed successfully")

	return nil
}

type collector struct {
}

func (c *collector) HandleResult(r *jw.Result) {
	logrus.WithFields(logrus.Fields{
		"time_string": r.GetValue(keyTimeString).(string),
	}).WithError(r.GetError()).Errorf("Failed to parse time")
}
