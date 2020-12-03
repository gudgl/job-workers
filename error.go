package main

type JWError interface {
	error
	HandleError()
}
