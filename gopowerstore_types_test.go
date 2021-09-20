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

package gopowerstore

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAPIError_VolumeNameIsAlreadyUse(t *testing.T) {
	apiError := NewAPIError()
	assert.False(t, apiError.VolumeNameIsAlreadyUse())
	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.ErrorCode = VolumeNameAlreadyUseErrorCode
	assert.True(t, apiError.VolumeNameIsAlreadyUse())
}

func TestAPIError_VolumeIsNotExist(t *testing.T) {
	apiError := NewAPIError()
	assert.False(t, apiError.NotFound())
	apiError.StatusCode = http.StatusNotFound
	apiError.ErrorCode = UnknownVolumeErrorCode
	assert.True(t, apiError.NotFound())
}

func TestAPIError_HostIsNotAttachedToVolume(t *testing.T) {
	apiError := NewHostIsNotAttachedToVolume()
	assert.True(t, apiError.HostIsNotAttachedToVolume())
}

func TestAPIError_VolumeAttachedToHost(t *testing.T) {
	apiError := NewVolumeAttachedToHostError()
	assert.True(t, apiError.VolumeAttachedToHost())
}
