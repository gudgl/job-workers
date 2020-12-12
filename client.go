package jw

import (
	"errors"
	"sync"
)

var (
	ErrNoWorkers    = errors.New("you must create at least one worker")
	ErrNoCollectors = errors.New("you must create at least one collector")
)

type Client interface {
	Do()
	SendJob(job Job)
	Wait()
}

type client struct {
	numberOfWorkers    int
	numberOfCollectors int
	jobs               chan Job
	errors            chan Error
	wgWorkers          sync.WaitGroup
	wgCollectors       sync.WaitGroup
}

func NewClient(
	numberOfWorkers,
	numberOfCollectors int,
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
		errors:            make(chan Error),
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
}

func (jw *client) Wait() {
	close(jw.jobs)
	jw.wgWorkers.Wait()

	close(jw.errors)
	jw.wgCollectors.Wait()
}

func (jw *client) SendJob(job Job) {
	if job != nil {
		jw.jobs <- job
	}
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
		err.HandleError()
	}
}
