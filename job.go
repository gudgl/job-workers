package jw

// Job represents the task for the workers
type Job interface {
	Execute() *Result
}
