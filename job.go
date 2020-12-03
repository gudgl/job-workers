package main

type Job interface {
	Do() JWError
}
