package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	f := func(sec int, ctx context.Context, ts TimeoutSemaphoreInterface) error {
		if err := ts.Acquire(ctx); err != nil {
			return err
		}
		time.Sleep(time.Duration(sec) * time.Second)
		ts.Release(ctx)

		return nil
	}

	// long running function
	ts := NewTimeoutSemaphore(1, 1, &defaultLogger{})
	go f(3, context.Background(), ts)
	// wait for run long function
	time.Sleep(1 * time.Second)
	err := f(1, context.Background(), ts)
	assert.NotNil(t, err)

	// fast running function
	ts = NewTimeoutSemaphore(3, 1, &defaultLogger{})
	go f(1, context.Background(), ts)
	err = f(2, context.Background(), ts)
	assert.Nil(t, err)
}
