package main

import (
	"sync"
)

type JobWorkersConf struct {
	numberOfWorkers    int
	numberOfCollectors int
	jobs               chan Job
	errors             chan JWError
	wgWorkers          sync.WaitGroup
	wgCollectors       sync.WaitGroup
}

func NewJobWorkerConf(
	numberOfWorkers,
	numberOfCollectors int,
	job chan Job,
	errors chan JWError,
) JobWorkersConf {
	return JobWorkersConf{
		numberOfWorkers:    numberOfWorkers,
		numberOfCollectors: numberOfCollectors,
		jobs:               job,
		errors:             errors,
	}
}

func (jw *JobWorkersConf) Execute() {
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

func (jw *JobWorkersConf) startWorkers() {
	defer jw.wgWorkers.Done()

	for job := range jw.jobs {
		err := job.Do()
		if err != nil {
			jw.errors <- err
		}
	}
}

func (jw *JobWorkersConf) startCollectors() {
	defer jw.wgCollectors.Done()

	for err := range jw.errors {
		err.HandleError()
	}
}
