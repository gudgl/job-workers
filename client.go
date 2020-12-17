package jw

import (
	"errors"
	"sync"
)

var (
	ErrNoWorkers     = errors.New("invalid number of worker")
	ErrNoCollectors  = errors.New("invalid number of collector")
	ErrNilCollectors = errors.New("collector is required")
)

// Client orchestrates the workers and collectors
type Client interface {
	Init()
	SendJob(job Job)
	SendJobs(job ...Job)
	Wait()
}

// client implements Client interface
type client struct {
	numberOfWorkers    int
	numberOfCollectors int
	jobs               chan Job
	results            chan *Result
	wgWorkers          sync.WaitGroup
	wgCollectors       sync.WaitGroup
	collector          Collector
}

// New creates new jw client
// NOTE: there must be a collector and at least one worker and at least oen collector
func New(
	collector Collector,
	numberOfWorkers,
	numberOfCollectors int,
) (Client, error) {
	if numberOfWorkers < 0 {
		return nil, ErrNoWorkers
	}

	if numberOfCollectors < 0 {
		return nil, ErrNoCollectors
	}

	if collector == nil {
		return nil, ErrNilCollectors
	}

	return &client{
		collector:          collector,
		numberOfWorkers:    numberOfWorkers,
		numberOfCollectors: numberOfCollectors,
		jobs:               make(chan Job),
		results:            make(chan *Result),
	}, nil
}

// Init creates the workers and the collectors
func (jw *client) Init() {
	for i := 0; i < jw.numberOfWorkers; i++ {
		jw.wgWorkers.Add(1)

		go jw.startWorkers()
	}

	for i := 0; i < jw.numberOfCollectors; i++ {
		jw.wgCollectors.Add(1)

		go jw.startCollectors()
	}
}

// Wait closes the channels and waits for all workers and collectors to finish
func (jw *client) Wait() {
	close(jw.jobs)
	jw.wgWorkers.Wait()

	close(jw.results)
	jw.wgCollectors.Wait()
}

// SendJob sends a job to the jobs channel
func (jw *client) SendJob(job Job) {
	if job != nil {
		jw.jobs <- job
	}
}

// SendJobs sends all given jobs to the job channel
func (jw *client) SendJobs(jobs ...Job) {
	for _, job := range jobs {
		jw.jobs <- job
	}
}

// startWorkers assigning the worker to listen to the jobs channel
func (jw *client) startWorkers() {
	defer jw.wgWorkers.Done()

	for job := range jw.jobs {
		r := job.Execute()
		if r != nil {
			jw.results <- r
		}
	}
}

// startCollectors assigning the collectors to listen to the errors channel
func (jw *client) startCollectors() {
	defer jw.wgCollectors.Done()

	for r := range jw.results {
		jw.collector.HandleResult(r)
	}
}
