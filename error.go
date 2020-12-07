package main

type Collector interface {
	HandleError(err error)
}
