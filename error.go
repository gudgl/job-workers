package job_workers

type Collector interface {
	HandleError(err error)
}
