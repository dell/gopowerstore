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
	"net/http"
)

// VolumeStateEnum Volume life cycle states.
type VolumeStateEnum string

const (
	// VolumeStateEnumReady - Volume is operating normally
	VolumeStateEnumReady VolumeStateEnum = "Ready"
	// VolumeStateEnumInitializing - Volume is starting but not yet ready for use
	VolumeStateEnumInitializing VolumeStateEnum = "Initializing"
	// VolumeStateEnumOffline - Volume is not available
	VolumeStateEnumOffline VolumeStateEnum = "Offline"
	// VolumeStateEnumDestroying - Volume is being deleted. No new operations are allowed
	VolumeStateEnumDestroying VolumeStateEnum = "Destroying"
)

// VolumeTypeEnum Type of volume.
type VolumeTypeEnum string

const (
	// VolumeTypeEnumPrimary - A base object.
	VolumeTypeEnumPrimary VolumeTypeEnum = "Primary"
	// VolumeTypeEnumClone - A read-write object that shares storage with the object from which it is sourced.
	VolumeTypeEnumClone VolumeTypeEnum = "Clone"
	// VolumeTypeEnumSnapshot - A read-only object created from a volume or clone.
	VolumeTypeEnumSnapshot VolumeTypeEnum = "Snapshot"
)

// StorageCreatorTypeEnum Creator type of the storage resource.
type StorageCreatorTypeEnum string

const (
	// StorageCreatorTypeEnumUser - A resource created by a user
	StorageCreatorTypeEnumUser StorageCreatorTypeEnum = "User"
	// StorageCreatorTypeEnumSystem - A resource created by the replication engine.
	StorageCreatorTypeEnumSystem StorageCreatorTypeEnum = "System"
	// StorageCreatorTypeEnumScheduler - A resource created by the snapshot scheduler
	StorageCreatorTypeEnumScheduler StorageCreatorTypeEnum = "Scheduler"
)

// StorageTypeEnum Possible types of storage for a volume.
type StorageTypeEnum string

const (
	// StorageTypeEnumBlock - Typical storage type that is displayed for all system management.
	StorageTypeEnumBlock StorageTypeEnum = "Block"
	// StorageTypeEnumFile - Volume internal to an SD-NAS file_system or nas_server object. Not manageable by the external user
	StorageTypeEnumFile StorageTypeEnum = "File"
)

// VolumeCreate create volume request
type VolumeCreate struct {
	// Unique name for the volume to be created.
	// This value must contain 128 or fewer printable Unicode characters.
	Name *string `json:"name"`
	// Optional sector size, in bytes. Only 512-byte and 4096-byte sectors are supported.
	SectorSize *int64 `json:"sector_size,omitempty"`
	// Size of the volume to be created, in bytes. Minimum volume size is 1MB.
	// Maximum volume size is 256TB. Size must be a multiple of 8192.
	Size *int64 `json:"size"`
	// Volume group to add the volume to. If not specified, the volume is not added to a volume group.
	VolumeGroupID string `json:"volume_group_id,omitempty"`

	MetaDataHeader
}

// VolumeComputeDifferences compute snapshot differences in a volume request
type VolumeComputeDifferences struct {
	// Unique identifier of the snapshot used to determine the differences from the current snapshot.
	// If not specified, returns all allocated extents of the current snapshot.
	// The base snapshot must be from the same base volume as the snapshot being compared with.
	BaseSnapshotID *string `json:"base_snapshot_id"`
	// The position of the first logical byte to be used in the comparison.
	// If not specified, the comparison starts at the beginning of the snapshot.
	// The offset must be a multiple of the chunk_size. For best performance, use a multiple of 4K bytes.
	Offset *int64 `json:"offset"`
	// Length of the comparison scan segment in bytes. length / chunk_size is the number of chunks,
	// with each chunk represented as a bit in the bitmap returned in the response. The number of chunks
	// must be divisible by 8 so that the returned bitmap is a byte array. The length and chunk_size
	// must be chosen so that there are no more than 32K chunks, resulting in a returned byte array
	// bitmap of at most 4K bytes. The length starting from the offset must not exceed the size of
	// the snapshot. The length must be a multiple of the chunk_size.
	Length *int64 `json:"length"`
	// Granularity of the chunk in bytes. Must be a power of 2 so that each bit in the returned
	// bitmap represents a chunk sized range of bytes.
	ChunkSize *int64 `json:"chunk_size"`
}

// VolumeComputeDifferencesResponse compute snapshot differences in a volume response
type VolumeComputeDifferencesResponse struct {
	// Base64-encoded bitmap with bits set for chunks that are either:
	// Allocated and nonzero when base_snapshot_id not specified, or
	// Unshared with the base snapshot when a base_snapshot_id is specified
	ChunkBitmap *string `json:"chunk_bitmap"`
	// Recommended offset to be used for the next compute_differences invocation
	// A value of -1 will be returned if the end of the object has been reached
	// while scanning for differences or allocations
	NextOffset *int64 `json:"next_offset"`
}

// MetaData returns the metadata headers.
func (vc *VolumeCreate) MetaData() http.Header {
	vc.once.Do(func() {
		vc.metadata = make(http.Header)
	})
	return vc.metadata
}

// VolumeModify modify volume request
type VolumeModify struct {
	// Unique identifier of the volume instance.
	Name string `json:"name,omitempty"`
	//  Size of the volume in bytes. Minimum volume size is 1MB. Maximum volume size is 256TB.
	//  Size must be a multiple of 8192.
	Size int64 `json:"size,omitempty"`
	// Unique identifier of the protection policy assigned to the volume.
	ProtectionPolicyID string `json:"protection_policy_id"`
}

// VolumeClone request for cloning snapshot/volume
type VolumeClone struct {
	// Unique name for the volume to be created.
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	MetaDataHeader
}

// MetaData returns the metadata headers.
func (vc *VolumeClone) MetaData() http.Header {
	vc.once.Do(func() {
		vc.metadata = make(http.Header)
	})
	return vc.metadata
}

// SnapshotCreate params for creating 'create snapshot' request
type SnapshotCreate struct {
	// Unique name for the snapshot to be created.
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// VolumeDelete body for VolumeDelete request
type VolumeDelete struct {
	ForceInternal *bool `json:"force_internal,omitempty"`
}

// Volume Details about a volume, including snapshots and clones of volumes.
type Volume struct {
	Description string `json:"description,omitempty"`
	// Unique identifier of the volume instance.
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	//  Size of the volume in bytes. Minimum volume size is 1MB. Maximum volume size is 256TB.
	//  Size must be a multiple of 8192.
	Size int64 `json:"size,omitempty"`
	// state
	State VolumeStateEnum `json:"state,omitempty"`
	// type
	Type VolumeTypeEnum `json:"type,omitempty"`
	// volume topology
	// World wide name of the volume.
	Wwn string `json:"wwn,omitempty"`

	// ApplianceID - Placeholder for appliance ID where the volume resides
	ApplianceID string `json:"appliance_id,omitempty"`

	ProtectionData ProtectionData `json:"protection_data,omitempty"`
}

// ProtectionData is a field that holds meta information about volume creation
type ProtectionData struct {
	SourceID string `json:"source_id"`
}

// Fields returns fields which must be requested to fill struct
func (v *Volume) Fields() []string {
	return []string{"description", "id", "name",
		"size", "state", "type", "wwn", "appliance_id", "protection_data"}
}
