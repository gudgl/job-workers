package jw

// Job represents the task for the workers
type Job interface {
	// Execute represents how the Job should be done
	Execute() *Result
}
