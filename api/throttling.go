/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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

type TimeoutSemaphore struct {
	timeout   time.Duration
	semaphore chan struct{}
	logger    Logger
}

func NewTimeoutSemaphore(timeout, rateLimit int, logger Logger) *TimeoutSemaphore {
	log := logger

	if log == nil {
		log = &defaultLogger{}
	}

	return &TimeoutSemaphore{
		timeout:   time.Duration(timeout) * time.Second,
		semaphore: make(chan struct{}, rateLimit),
		logger:    log,
	}
}

func (ts *TimeoutSemaphore) Acquire(ctx context.Context) error {
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

func (ts *TimeoutSemaphore) Release(ctx context.Context) {
	<-ts.semaphore
	ts.logger.Debug(ctx, "release a lock")
}

func (ts *TimeoutSemaphore) SetLogger(logger Logger) TimeoutSemaphoreInterface {
	ts.logger = logger
	return ts
}
