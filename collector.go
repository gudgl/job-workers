package jw

// Collector represents someone who handles the errors
type Collector interface {
	// HandleResult represents how the Result should be handled
	HandleResult(r *Result)
}
