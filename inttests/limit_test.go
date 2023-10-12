/*
 *
 * Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxVolumeSize(t *testing.T) {
	customHeaders := C.GetCustomHTTPHeaders()
	if customHeaders == nil {
		customHeaders = make(http.Header)
	}
	customHeaders.Add("DELL-VISIBILITY", "internal")
	C.SetCustomHTTPHeaders(customHeaders)

	limit, err := C.GetMaxVolumeSize(context.Background())

	// reset custom header
	customHeaders.Del("DELL-VISIBILITY")
	C.SetCustomHTTPHeaders(customHeaders)

	checkAPIErr(t, err)
	assert.Positive(t, limit)
}

func TestGetMaxVolumeSizeEndpointNotFound(t *testing.T) {
	limit, err := C.GetMaxVolumeSize(context.Background())

	assert.Equal(t, "The REST endpoint [GET /api/rest/limit?select=id%2Climit] cannot be found.", err.Error())
	assert.Negative(t, limit)
}
