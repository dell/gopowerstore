/*
 *
 * Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package inttests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancelFunc()
	_, err := C.GetVolumes(ctx)
	assert.NotNil(t, err)
}

func TestTraceID(t *testing.T) {
	_, err := C.GetVolumes(C.SetTraceID(context.Background(), "reqid-1"))
	assert.Nil(t, err)
}

func TestCustomLogger(t *testing.T) {
	C.SetLogger(&customLogger{})
	_, err := C.GetVolumes(C.SetTraceID(context.Background(), "reqid-1"))
	assert.Nil(t, err)
}

type customLogger struct{}

func (cl *customLogger) Info(ctx context.Context, format string, args ...interface{}) {
	log.Printf("INFO:"+format, args...)
}

func (cl *customLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	log.Printf("DEBUG:"+format, args...)
}

func (cl *customLogger) Error(ctx context.Context, format string, args ...interface{}) {
	log.Printf("ERROR:"+format, args...)
}
