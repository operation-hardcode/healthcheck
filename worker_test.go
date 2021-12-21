package healthcheck

import (
	"testing"
	"time"
)

func TestWorkerAddJob(t *testing.T) {
	w := NewWorker()

	w.AddJob(NewJob("1", time.Millisecond, func(j Job) {
		j.Cancel()
	}))

	if w.Size() != 1 {
		t.Error("The count of jobs must be equal to 1")
	}

	w.CancelJobs()

	if w.Size() != 0 {
		t.Error("The count of jobs must be equal to 0")
	}
}

func TestWorkerCancelJobById(t *testing.T) {
	w := NewWorker()

	job1 := NewJob("1", time.Millisecond, func(j Job) {})
	job2 := NewJob("2", time.Millisecond, func(j Job) {})

	w.AddJob(job1)
	w.AddJob(job2)

	if w.Size() != 2 {
		t.Error("The count of jobs must be equal to 2")
	}

	w.CancelJob("1")

	if w.Size() != 1 {
		t.Error("The count of jobs must be equal to 1")
	}

	w.CancelJobs()

	if w.Size() != 0 {
		t.Error("The count of jobs must be equal to 0")
	}
}
