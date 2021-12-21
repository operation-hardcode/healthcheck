package healthcheck

import (
	"os"
	"os/signal"
	"syscall"
)

type worker struct {
	jobs map[JobId]*job
}

func NewWorker() *worker {
	return &worker{jobs: map[JobId]*job{}}
}

func (w *worker) AddJob(j *job) {
	w.jobs[j.id] = j

	j.Start()
}

func (w *worker) CancelJob(jid JobId) {
	job, ok := w.jobs[jid]

	if ok {
		delete(w.jobs, job.id)

		job.Cancel()
	}
}

func (w *worker) CancelJobs() {
	for _, job := range w.jobs {
		job.Cancel()

		delete(w.jobs, job.id)
	}
}

func (w *worker) Size() int {
	return len(w.jobs)
}

func (w *worker) Work() <-chan struct{} {
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
