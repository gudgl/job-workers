package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gudgl/job-workers"
)

const layout = "2006-01-02"

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
	client, err := jw.NewClient(10, 10)
	if err != nil {
		logrus.Fatal("Failed to create client.")
		return
	}

	client.Do()

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

func (j *job) Execute() jw.Error {
	time, err := time.Parse(layout, j.timeString)
	if err != nil {
		return &e{
			err,
			j.timeString,
		}
	}

	logrus.WithFields(logrus.Fields{
		"time_string": j.timeString,
		"time":        time,
	}).Info("Time parsed successfully")

	return nil
}

type e struct {
	err        error
	failedTime string
}

func (e *e) HandleError() {
	logrus.WithFields(logrus.Fields{
		"time_string": e.failedTime,
	}).WithError(e.err).Errorf("Failed to parse time")
}
