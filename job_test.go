package healthcheck

import (
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	count := 0

	j := NewJob("1", time.Millisecond, func(j Job) {
		count++
		j.Cancel()
	})

	if j.id != "1" {
		t.Error("Expected 1, got ", j.id)
	}

	j.Start()
	<-j.ctx.Done()

	if count != 1 {
		t.Error("The job need to call handler at least once.")
	}
}
