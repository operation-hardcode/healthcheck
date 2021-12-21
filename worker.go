package healthcheck

import (
	"os"
	"os/signal"
	"syscall"
)

type Worker struct {
	jobs map[JobId]*Job
}

func NewWorker() *Worker {
	return &Worker{jobs: map[JobId]*Job{}}
}

func (w *Worker) AddJob(j *Job) {
	w.jobs[j.id] = j

	j.Start()
}

func (w *Worker) CancelJob(jid JobId) {
	job, ok := w.jobs[jid]

	if ok {
		job.Cancel()

		delete(w.jobs, job.id)
	}
}

func (w *Worker) CancelJobs() {
	for _, job := range w.jobs {
		job.Cancel()

		delete(w.jobs, job.id)
	}
}

func (w *Worker) Size() int {
	return len(w.jobs)
}

func (w *Worker) Work() <-chan struct{} {
	quit := make(chan struct{})

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(quit)
		for {
			select {
			case <-signals:
				w.CancelJobs()

				return
			default:
				//
			}
		}
	}()

	return quit
}
