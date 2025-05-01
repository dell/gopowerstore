/*
 *
 * Copyright © 2021-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"math/rand"
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
	Timeout   time.Duration
	Delay     time.Duration
	Semaphore chan struct{}
	Logger    Logger
}

func NewTimeoutSemaphore(timeout int64, rateLimit int, delay int, logger Logger) *TimeoutSemaphore {
	log := logger

	if log == nil {
		log = &defaultLogger{}
	}

	return &TimeoutSemaphore{
		Timeout:   time.Duration(timeout) * time.Second,
		Semaphore: make(chan struct{}, rateLimit),
		Delay:     time.Duration(delay) * time.Millisecond,
		Logger:    log,
	}
}

func (ts *TimeoutSemaphore) Acquire(ctx context.Context) error {
	acquireCtx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	ts.Logger.Info(ctx, "Number of elements in queue/sem channel %d", len(ts.Semaphore))
	for {
		select {
		case ts.Semaphore <- struct{}{}:
			ts.Logger.Info(ctx, "Acquired lock in throttling")
			s := rand.Intn(100)
			ts.Logger.Info(ctx, "Sleeping for %d millis before moving ahead", s)
			time.Sleep(time.Duration(s))
			return nil
		case <-ctx.Done():
			msg := "main context expired"
			ts.Logger.Error(ctx, msg)
			return &TimeoutSemaphoreError{msg}
		case <-acquireCtx.Done():
			msg := "acquire context expired"
			ts.Logger.Error(ctx, msg)
			return &TimeoutSemaphoreError{msg}
		}
	}
}

func (ts *TimeoutSemaphore) Release(ctx context.Context) {
	<-ts.Semaphore
	ts.Logger.Debug(ctx, "release a lock")
}

func (ts *TimeoutSemaphore) SetLogger(logger Logger) TimeoutSemaphoreInterface {
	ts.Logger = logger
	return ts
}
