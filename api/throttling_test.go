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

//nolint:revive
package api

import (
	"context"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	tests := []struct {
		name     string
		ctx      func() (context.Context, context.CancelFunc)
		ts       func() TimeoutSemaphoreInterface
		holdTime time.Duration
		timeout  time.Duration
		wantErr  bool
	}{
		{
			name: "successfully rate limits",
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 200*time.Millisecond)
			},
			ts: func() TimeoutSemaphoreInterface {
				return NewTimeoutSemaphore(300*time.Millisecond, 1, &defaultLogger{})
			},
			// should hold the semaphore for less than the ctx timeout
			holdTime: 100 * time.Millisecond,
			wantErr:  false,
		},
		{
			name: "second call times out",
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 101*time.Millisecond)
			},
			ts: func() TimeoutSemaphoreInterface {
				return NewTimeoutSemaphore(100*time.Millisecond, 1, &defaultLogger{})
			},
			// hold the semaphore for longer than the context timeout
			holdTime: 200 * time.Millisecond,
			wantErr:  true,
		},
		{
			name: "context is canceled",
			ctx: func() (context.Context, context.CancelFunc) {
				// cancel the context to trigger <-ctx.Done() condition
				ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
				cancel()
				// an empty cancel func since we've already called it
				return ctx, func() {}
			},
			ts: func() TimeoutSemaphoreInterface {
				return NewTimeoutSemaphore(100*time.Millisecond, 1, &defaultLogger{})
			},
			holdTime: 1 * time.Millisecond,
			wantErr:  true,
		},
		{
			name: "context timeout is shorter than default timeout",
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 100*time.Millisecond)
			},
			ts: func() TimeoutSemaphoreInterface {
				return NewTimeoutSemaphore(200*time.Millisecond, 1, &defaultLogger{})
			},
			holdTime: 10 * time.Millisecond,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := tt.ts()

			// fill the semaphore
			err := ts.Acquire(context.Background())
			if err != nil {
				t.Errorf("failed to acquire semaphore: %v", err)
			}

			// hold the semaphore for some time before releasing it
			// so that the next acquisition attempt is blocked
			go func() {
				defer ts.Release(context.Background())
				time.Sleep(tt.holdTime)
			}()

			ctx, cancel := tt.ctx()
			defer cancel()

			// try to acquire the semaphore
			err = ts.Acquire(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Acquire() returned an unexpected error: %v", err)
			}
		})
	}
}
