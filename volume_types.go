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
	// Storage type. Valid values are:
	StorageType *StorageTypeEnum `json:"storage_type,omitempty"`
	MetaDataHeader
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
	Size int64 `json:"size"`
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
	// Storage type. Valid values are:
	StorageType StorageTypeEnum `json:"storage_type,omitempty"`
	// type
	Type VolumeTypeEnum `json:"type,omitempty"`
	// volume topology
	// World wide name of the volume.
	Wwn string `json:"wwn,omitempty"`

	ProtectionData ProtectionData `json:"protection_data,omitempty"`
}

// ProtectionData is a field that holds meta information about volume creation
type ProtectionData struct {
	SourceID string `json:"source_id"`
}

// Fields returns fields which must be requested to fill struct
func (v *Volume) Fields() []string {
	return []string{"description", "id", "name",
		"size", "state", "storage_type", "type", "wwn", "protection_data"}
}
