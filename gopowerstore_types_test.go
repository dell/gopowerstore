/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_VolumeNameIsAlreadyUse(t *testing.T) {
	apiError := NewAPIError()
	assert.False(t, apiError.VolumeNameIsAlreadyUse())
	apiError.StatusCode = http.StatusUnprocessableEntity
	assert.True(t, apiError.VolumeNameIsAlreadyUse())
}

func TestAPIError_VolumeIsNotExist(t *testing.T) {
	apiError := NewAPIError()
	assert.False(t, apiError.NotFound())
	apiError.StatusCode = http.StatusNotFound
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

func TestAPIError_VolumeDetachedFromHost(t *testing.T) {
	apiError := NewVolumeAttachedToHostError()
	assert.True(t, apiError.VolumeDetachedFromHost())
}

func TestAPIError_ReplicationSessionAlreadyCreated(t *testing.T) {
	apiError := NewAPIError()
	apiError.StatusCode = http.StatusUnprocessableEntity
	assert.False(t, apiError.ReplicationSessionAlreadyCreated())

	apiError.StatusCode = http.StatusBadRequest
	assert.True(t, apiError.ReplicationSessionAlreadyCreated())
}

func TestAPIError_VolumeAlreadyRemovedFromVolumeGroup(t *testing.T) {
	apiError := NewAPIError()
	apiError.StatusCode = http.StatusBadRequest
	assert.False(t, apiError.VolumeAlreadyRemovedFromVolumeGroup())

	apiError.StatusCode = http.StatusUnprocessableEntity
	assert.False(t, apiError.VolumeAlreadyRemovedFromVolumeGroup())

	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.Message = "One or more volumes to be removed are not part of the volume group"
	assert.True(t, apiError.VolumeAlreadyRemovedFromVolumeGroup())
}

func TestAPIError_FSCreationLimitReached(t *testing.T) {
	apiError := NewAPIError()

	// Case 1: Wrong status code
	apiError.StatusCode = http.StatusBadRequest
	apiError.Message = "The limit of 125 file systems for the NAS server has been reached"
	assert.False(t, apiError.FSCreationLimitReached())

	// Case 2: Correct status, but unrelated message
	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.Message = "some unrelated error message"
	assert.False(t, apiError.FSCreationLimitReached())

	// Case 3: Correct status and matching message
	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.Message = "The limit of 125 file systems for the NAS server has been reached"
	assert.True(t, apiError.FSCreationLimitReached())

	// Case 4: Correct status and message contains code
	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.Message = "error code 0xE08010080451: The limit of 125 file systems for the NAS server"
	assert.True(t, apiError.FSCreationLimitReached())

	// Case 5: Leading/trailing whitespace (optional robustness)
	apiError.StatusCode = http.StatusUnprocessableEntity
	apiError.Message = "   The limit of 125 file systems for the NAS server   "
	assert.True(t, apiError.FSCreationLimitReached())
}
