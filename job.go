package healthcheck

import (
	"context"
	"time"
)

type (
	JobId string

	handler func(Job)

	Job struct {
		id          JobId
		ctx         context.Context
		cancel      context.CancelFunc
		healthcheck time.Duration
		handler     handler
	}
)

func NewJob(id JobId, healthcheck time.Duration, h handler) *Job {
	ctx, cancel := context.WithCancel(context.Background())

	return &Job{
		id:          id,
		ctx:         ctx,
		cancel:      cancel,
		healthcheck: healthcheck,
		handler:     h,
	}
}

func (j *Job) Start() {
	go func() {
		ticker := time.NewTicker(j.healthcheck)

		for {
			select {
			case <-j.ctx.Done():
				return
			case <-ticker.C:
				j.handler(*j)
			default:
				//
			}
		}
	}()
}

func (j *Job) Cancel() {
	j.cancel()
}
