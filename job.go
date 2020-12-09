package job_workers

type Job interface {
	Execute() error
}
