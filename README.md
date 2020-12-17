# job-workers or shorter jw

-----

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
    err error
    keys map[string]interface{}
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
client, r := jw.NewClient(3, 3)
```

Next run the `Do` function from the client to create the workers and the collectors

```go
client.Do()
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

See the [examples](https://github.com/gudgl/job-workers/tree/main/examples) for more details

## Collaboration

Feel free to contribute and improve jw

## Licence

[MIT](https://github.com/gudgl/job-workers/blob/main/LICENSE)