package jw

// Collector represents someone who handles the errors
type Collector interface {
	// HandleResult is used for handling the Result
	HandleResult(r *Result)
}
