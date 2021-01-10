# job-workers or shorter jw

-----
[![Go Reference](https://pkg.go.dev/badge/github.com/gudgl/job-workers.svg)](https://pkg.go.dev/github.com/gudgl/job-workers) \
jw is a library for asynchronous execution of tasks and handling results from them

## Install

```textmate
go get github.com/gudgl/job-workers
```

## Usage

First you need to implement both interfaces `Job` and `Collector`

```go
type Job interface {
    Execute() *Result
}
```

```go
type Collector interface {
    HandleResult(r *Result)
}
```

The `Execute` method is what workers should do, how to do the job. On the other hand `HandlerResult` is how the `Result`
should be handled from the collectors.

```go
type Keys map[string]interface{}

type Result struct {
    Error error
    Keys map[string]interface{}
}
```

This is the `Result` with a list of methods:
- `WithError(err error) *Result`
- `WithKeys(keys Keys) *Result`
- `WithKey(key string, value interface{}) *Result`
- `GetValue(key string) interface{}`
- `GetError() error`

To create new `Result` use one of:
- `NewResult() *Result` - create empty `Result`
- `WithError(err error) *Result` - create `Result` with `error`
- `WithKeys(keys Keys) *Result` - create `Result` with given `keys`
- `WithKey(key string, value interface{}) *Result` - create `Result` with the given pair of `key` `value`

Second you should create the client. It should be created with at least one worker, or it returns error from
type `ErrNoWorkers`, and with at least one collector, or it returns error from type `ErrNoCollectors`

```go
client, r := jw.NewClient(collector, 3, 3)
```

Next run the `Go` function from the client to start the workers and the collectors

```go
client.Go()
```

Then give/send them jobs one by one

```go
for _, job := range jobs{
    client.SendJob(job)
}
```

or send them all at once

```go
client.SendJobs(jobs)
```

Last just wait for workers and collectors to finish

```go
client.Wait()
```

That's all you need to do to get it up running

## Examples

```go
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
```

## Collaboration

If you would like to contribute, head on to the [issues](https://github.com/gudgl/job-workers/issues/new) page for tasks that need help.

## Licence

[MIT](https://github.com/gudgl/job-workers/blob/main/LICENSE)