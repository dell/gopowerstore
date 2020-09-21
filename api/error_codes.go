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

package api

const (
	// UnknownVolumeErrorCode - unknown volume error
	UnknownVolumeErrorCode = "0xE0A08001000E"
	// HostIsNotAttachedToVolumeErrorCode - host not attached to volume
	HostIsNotAttachedToVolumeErrorCode = "0xE0A01001001E"
	// InvalidInstance - instance not found
	InvalidInstance = "0xE04040020002"
	// VolumeNameAlreadyUseErrorCode - volume already exists
	VolumeNameAlreadyUseErrorCode = "0xE0A080010014"
	// SnapshotNameAlreadyUseErrorCode - snapshot already exists
	SnapshotNameAlreadyUseErrorCode = "0xE0A060010012"
	// FilesystemNameAlreadyUseErrorCode - fs already exists
	FilesystemNameAlreadyUseErrorCode = "0xE08010080009"
	// NoHostObjectFoundCode - no host object found in management db for specified ID
	NoHostObjectFoundCode = "0xE0A010010003"
	// BadRangeCode - invalid range was used in request
	BadRangeCode = "0xE04040010005"
	// VolumeAttachedToHost - volume attached to host
	VolumeAttachedToHost = "0xE0A080020001"
	// HostAlreadyPresentInNFSExport - host already have an access
	HostAlreadyPresentInNFSExport = "0xE080100F0008"
	// HostAlreadyRemovedFromNFSExport
	HostAlreadyRemovedFromNFSExport = "0xE080100F0009"
	// InstanceWasNotFound - Instance was not found on array
	InstanceWasNotFound = "0xE04040020009"
)
