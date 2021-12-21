package healthcheck

import (
	"context"
	"time"
)

type (
	JobId string

	handler func(job)

	job struct {
		id          JobId
		ctx         context.Context
		cancel      context.CancelFunc
		healthcheck time.Duration
		handler     handler
	}
)

func NewJob(id JobId, healthcheck time.Duration, h handler) *job {
	ctx, cancel := context.WithCancel(context.Background())

	return &job{
		id:          id,
		ctx:         ctx,
		cancel:      cancel,
		healthcheck: healthcheck,
		handler:     h,
	}
}

func (j *job) Start() {
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

func (j *job) Cancel() {
	j.cancel()
}
