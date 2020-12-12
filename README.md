# job-workers or shorter jw

-----

jw is a library for asynchronous execution of tasks and handling results from them

## Install
```textmate
go get github.com/gudgl/job-workers
```

## Usage
First you need to implement both interfaces `Job` and `Error`
```go
type Job interface {
    Execute() Error
}
```
```go
type Error interface {
    HandleError()
}
```
The `Execute` method is what workers should do, how to do the job. 
On the other hand `HandlerError` is how the error should be handled from the collectors.

Second you should create the client. It should be created with at least one worker, 
or it returns error from type `ErrNoWorkers`, and with at least one collector, or it 
returns error from type `ErrNoCollectors`
```go
client, err := jw.NewClient(3, 3)
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
That's it

See the [examples](https://github.com/gudgl/job-workers/tree/main/examples) for more details

## Collaboration
Feel free to contribute and improve jw

## Licence
[MIT](https://github.com/gudgl/job-workers/blob/main/LICENSE)