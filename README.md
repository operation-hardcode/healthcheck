[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

# healthcheck

A zero-dependencies and lightweight go library for job scheduling.

## Quick Start

```go
package main

import (
	"github.com/operation-hardcode/healthcheck"
	"time"
)

func main() {
	worker := healthcheck.NewWorker()

	worker.AddJob(healthcheck.NewJob("1", time.Minute, func(job healthcheck.Job) { 
		// work here every minute.
		if something {
			 job.Cancel()
		}
	}))

	<-worker.Work() // waiting for signals.
}
```

## Cancellation

```go
package main

import (
	"github.com/operation-hardcode/healthcheck"
	"time"
)

func main() {
	worker := healthcheck.NewWorker()

	worker.AddJob(healthcheck.NewJob("1", time.Hour, func(job healthcheck.Job) { 
		// work here.
	}))

	worker.AddJob(healthcheck.NewJob("2", time.Minute, func(job healthcheck.Job) {
		// work here.
	}))

	worker.CancelJob("1") // cancel concrete job by its id.

	worker.CancelJobs() // or cancel all jobs.
}
```