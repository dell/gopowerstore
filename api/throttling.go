package api

import (
	"context"
	"time"
)

// TimeoutSemaphoreInterface gives ability to limit rate of requests to PowerStore API
type TimeoutSemaphoreInterface interface {
	Acquire(ctx context.Context) error
	Release(ctx context.Context)
	SetLogger(logger Logger) TimeoutSemaphoreInterface
}

type TimeoutSemaphoreError struct {
	msg string
}

func (e *TimeoutSemaphoreError) Error() string {
	return e.msg
}

type timeoutSemaphore struct {
	timeout   time.Duration
	semaphore chan struct{}
	logger    Logger
}

func NewTimeoutSemaphore(timeout, rateLimit int, logger Logger) *timeoutSemaphore {
	log := logger

	if log == nil {
		log = &defaultLogger{}
	}

	return &timeoutSemaphore{
		timeout:   time.Duration(timeout) * time.Second,
		semaphore: make(chan struct{}, rateLimit),
		logger:    log,
	}
}

func (ts *timeoutSemaphore) Acquire(ctx context.Context) error {
	var cancelFunc func()
	ctx, cancelFunc = context.WithTimeout(ctx, ts.timeout)
	defer cancelFunc()
	for {
		select {
		case ts.semaphore <- struct{}{}:
			ts.logger.Debug(ctx, "acquire a lock")
			return nil
		case <-ctx.Done():
			msg := "lock is acquire failed, timeout expired"
			ts.logger.Error(ctx, msg)
			return &TimeoutSemaphoreError{msg}
		}
	}
}

func (ts *timeoutSemaphore) Release(ctx context.Context) {
	<-ts.semaphore
	ts.logger.Debug(ctx, "release a lock")
}


func (ts *timeoutSemaphore) SetLogger(logger Logger) TimeoutSemaphoreInterface {
	ts.logger = logger
	return ts
}