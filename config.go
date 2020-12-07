package main

import (
	"errors"
	"sync"
)

var (
	ErrNoWorkers    = errors.New("must be at least one worker")
	ErrNoCollectors = errors.New("must be at least one collector")
)

type Client interface {
	Do()
	SendJob(job Job)
}

type client struct {
	numberOfWorkers    int
	numberOfCollectors int
	jobs               chan Job
	errors             chan error
	wgWorkers          sync.WaitGroup
	wgCollectors       sync.WaitGroup
	collector          Collector
}

func NewClient(
	numberOfWorkers,
	numberOfCollectors int,
	collector Collector,
) (Client, error) {
	if numberOfWorkers < 0 {
		return nil, ErrNoWorkers
	}

	if numberOfCollectors < 0 {
		return nil, ErrNoCollectors
	}

	return &client{
		numberOfWorkers:    numberOfWorkers,
		numberOfCollectors: numberOfCollectors,
		jobs:               make(chan Job),
		errors:             make(chan error),
		collector:          collector,
	}, nil
}

func (jw *client) Do() {
	for i := 0; i < jw.numberOfWorkers; i++ {
		jw.wgWorkers.Add(1)

		go jw.startWorkers()
	}

	for i := 0; i < jw.numberOfCollectors; i++ {
		jw.wgCollectors.Add(1)

		go jw.startCollectors()
	}

	jw.wgWorkers.Wait()
	close(jw.jobs)

	jw.wgCollectors.Wait()
	close(jw.errors)
}

func (jw *client) SendJob(job Job) {
	jw.jobs <- job
}

func (jw *client) startWorkers() {
	defer jw.wgWorkers.Done()

	for job := range jw.jobs {
		err := job.Execute()
		if err != nil {
			jw.errors <- err
		}
	}
}

func (jw *client) startCollectors() {
	defer jw.wgCollectors.Done()

	for err := range jw.errors {
		jw.collector.HandleError(err)
	}
}
