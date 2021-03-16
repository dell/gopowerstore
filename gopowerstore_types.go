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
	"github.com/dell/gopowerstore/api"
	"net/http"
)

const (
	// UnknownVolumeErrorCode indicates an unknown volume error
	UnknownVolumeErrorCode = api.UnknownVolumeErrorCode
	// VolumeNameAlreadyUseErrorCode indicates non unique volume name
	VolumeNameAlreadyUseErrorCode = api.VolumeNameAlreadyUseErrorCode
	// SnapshotNameAlreadyUseErrorCode indicates non unique snapshot name
	SnapshotNameAlreadyUseErrorCode = api.SnapshotNameAlreadyUseErrorCode
	// FilesystemNameAlreadyUseErrorCode indicates non unique fs name
	FilesystemNameAlreadyUseErrorCode = api.FilesystemNameAlreadyUseErrorCode
	// InvalidInstance - instance not found
	InvalidInstance = api.InvalidInstance
	// HostIsNotAttachedToVolumeErrorCode - host not attached to volume
	HostIsNotAttachedToVolumeErrorCode = api.HostIsNotAttachedToVolumeErrorCode
	// VolumeIsNotAttachedToHost - volume is not attached to host
	VolumeIsNotAttachedToHost = api.VolumeIsNotAttachedToHost
	// NoHostObjectFoundCode - no host object found in management db for specified ID
	NoHostObjectFoundCode = api.NoHostObjectFoundCode
	// BadRangeCode - invalid range was used in request
	BadRangeCode = api.BadRangeCode
	// VolumeAttachedToHost - volume attached to host
	VolumeAttachedToHost = api.VolumeAttachedToHost
	// InstanceWasNotFound - Instance was not found on array
	InstanceWasNotFound = api.InstanceWasNotFound
	// HostAlreadyPresentInNFSExport - host already have an access
	HostAlreadyPresentInNFSExport = api.HostAlreadyPresentInNFSExport
	// HostAlreadyRemovedFromNFSExport
	HostAlreadyRemovedFromNFSExport = api.HostAlreadyRemovedFromNFSExport
	//UnableToMatchHostVolume - Couldn't find any host volume matching volume id
	UnableToMatchHostVolume = api.UnableToMatchHostVolume
)

// RequestConfig represents options for request
type RequestConfig api.RequestConfig

// RenderRequestConfig returns internal struct with request config
func (rc RequestConfig) RenderRequestConfig() api.RequestConfig {
	return api.RequestConfig(rc)
}

// CreateResponse create response
type CreateResponse struct {
	// Unique identifier of the new instance created.
	ID string `json:"id,omitempty"`
}

// EmptyResponse is response without content
type EmptyResponse string

// APIError represents API error
type APIError struct {
	*api.ErrorMsg
}

// NewAPIError returns pointer to new APIError
func NewAPIError() *APIError {
	return &APIError{&api.ErrorMsg{}}
}

// WrapErr converts internal error type to public
func WrapErr(err error) error {
	errorMsg, ok := err.(*api.ErrorMsg)
	if ok {
		err = APIError{errorMsg}
	}
	return err
}

// VolumeIsNotExist returns true if API error indicate that volume is not exists
func (err *APIError) VolumeIsNotExist() bool {
	return (err.StatusCode == http.StatusNotFound || err.StatusCode == http.StatusUnprocessableEntity) &&
		(err.ErrorCode == UnknownVolumeErrorCode || err.ErrorCode == InvalidInstance ||
			err.ErrorCode == InstanceWasNotFound)
}

// VolumeNameIsAlreadyUse returns true if API error indicate that volume name is already in use
func (err *APIError) VolumeNameIsAlreadyUse() bool {
	return err.StatusCode == http.StatusUnprocessableEntity &&
		err.ErrorCode == VolumeNameAlreadyUseErrorCode
}

// SnapshotNameIsAlreadyUse returns true if API error indicate that snapshot name is already in use
func (err *APIError) SnapshotNameIsAlreadyUse() bool {
	return err.StatusCode == http.StatusBadRequest &&
		err.ErrorCode == SnapshotNameAlreadyUseErrorCode
}

// FSNameIsAlreadyUse returns true if API error indicate that fs name is already in use
func (err *APIError) FSNameIsAlreadyUse() bool {
	return (err.StatusCode == http.StatusBadRequest || err.StatusCode == http.StatusUnprocessableEntity) &&
		err.ErrorCode == FilesystemNameAlreadyUseErrorCode
}

// HostIsNotAttachedToVolume returns true if API error indicate that host is not attached to volume
func (err *APIError) HostIsNotAttachedToVolume() bool {
	return err.StatusCode == http.StatusBadRequest &&
		err.ErrorCode == HostIsNotAttachedToVolumeErrorCode
}

// VolumeIsNotAttachedToHost returns true if API error indicate that volume is not attached to host
func (err *APIError) VolumeIsNotAttachedToHost() bool {
	return (err.StatusCode == http.StatusBadRequest &&
		err.ErrorCode == VolumeIsNotAttachedToHost ) || (err.StatusCode == http.StatusBadRequest &&
		err.ErrorCode == UnableToMatchHostVolume)
}

// HostIsNotExist returns true if API error indicate that host is not exists
func (err *APIError) HostIsNotExist() bool {
	return (err.StatusCode == http.StatusNotFound || err.StatusCode == http.StatusBadRequest) &&
		(err.ErrorCode == InvalidInstance || err.ErrorCode == NoHostObjectFoundCode)
}

// BadRange returns true if API error indicate that request was submitted with invalid range
func (err *APIError) BadRange() bool {
	return err.StatusCode == http.StatusRequestedRangeNotSatisfiable || err.ErrorCode == BadRangeCode
}

// VolumeAttachedToHost returns true if API error indicate that operation can't be complete because
// volume is attached to host
func (err *APIError) VolumeAttachedToHost() bool {
	return err.StatusCode == http.StatusUnprocessableEntity || err.ErrorCode == VolumeAttachedToHost
}

// HostAlreadyRemovedFromNFSExport returns true if API error indicate that operation can't be complete because
// host ip already removed from nfs export access
func (err *APIError) HostAlreadyRemovedFromNFSExport() bool {
	return err.StatusCode == http.StatusUnprocessableEntity || err.ErrorCode == HostAlreadyRemovedFromNFSExport
}

// HostAlreadyPresentInNFSExport returns true if API error indicate that operation can't be complete because
// host ip already present in nfs export access
func (err *APIError) HostAlreadyPresentInNFSExport() bool {
	return err.StatusCode == http.StatusUnprocessableEntity || err.ErrorCode == HostAlreadyPresentInNFSExport
}

// NewVolumeIsNotExistError returns new VolumeIsNotExistError
func NewVolumeIsNotExistError() APIError {
	return notExistError()
}

// NewHostIsNotExistError returns new HostIsNotExistError
func NewHostIsNotExistError() APIError {
	return notExistError()
}

// NewHostIsNotAttachedToVolume returns new HostIsNotAttachedToVolume error
func NewHostIsNotAttachedToVolume() APIError {
	apiError := APIError{&api.ErrorMsg{}}
	apiError.ErrorCode = HostIsNotAttachedToVolumeErrorCode
	apiError.StatusCode = http.StatusBadRequest
	return apiError
}

// NewVolumeAttachedToHostError returns new VolumeAttachedToHost error
func NewVolumeAttachedToHostError() APIError {
	apiError := APIError{&api.ErrorMsg{}}
	apiError.ErrorCode = VolumeAttachedToHost
	apiError.StatusCode = http.StatusUnprocessableEntity
	return apiError
}

func notExistError() APIError {
	apiError := APIError{&api.ErrorMsg{}}
	apiError.ErrorCode = InvalidInstance
	apiError.StatusCode = http.StatusNotFound
	return apiError
}
