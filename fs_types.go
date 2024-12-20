/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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

// NASServerOperationalStatusEnum NAS lifecycle state.
type NASServerOperationalStatusEnum string

const (
	Stopped  NASServerOperationalStatusEnum = "Stopped"
	Starting NASServerOperationalStatusEnum = "Starting"
	Started  NASServerOperationalStatusEnum = "Started"
	Stopping NASServerOperationalStatusEnum = "Stopping"
	Failover NASServerOperationalStatusEnum = "Failover"
	Degraded NASServerOperationalStatusEnum = "Degraded"
	Unknown  NASServerOperationalStatusEnum = "Unknown"
)

type FileSystemTypeEnum string

const (
	FileSystemTypeEnumPrimary  FileSystemTypeEnum = "Primary"  // Normal file system or clone
	FileSystemTypeEnumSnapshot FileSystemTypeEnum = "Snapshot" // Snapshot of a file system
)

type FLRCreate struct {
	Mode             string `json:"mode,omitempty"`
	MinimumRetention string `json:"minimum_retention,omitempty"`
	DefaultRetention string `json:"default_retention,omitempty"`
	MaximumRetention string `json:"maximum_retention,omitempty"`
}

// FsCreate params for creating 'create fs' request
type FsCreate struct {
	Description                string      `json:"description,omitempty"`
	Name                       string      `json:"name,omitempty"`
	NASServerID                string      `json:"nas_server_id,omitempty"`
	Size                       int64       `json:"size_total,omitempty"`
	ConfigType                 string      `json:"config_type,omitempty"`
	AccessPolicy               string      `json:"access_policy,omitempty"`
	LockingPolicy              string      `json:"locking_policy,omitempty"`
	FolderRenamePolicy         string      `json:"folder_rename_policy,omitempty"`
	IsAsyncMTimeEnabled        *bool       `json:"is_async_MTime_enabled,omitempty"`
	ProtectionPolicyID         string      `json:"protection_policy_id,omitempty"`
	FileEventsPublishingMode   string      `json:"file_events_publishing_mode,omitempty"`
	HostIOSize                 string      `json:"host_io_size,omitempty"`
	FlrCreate                  interface{} `json:"flr_attributes,omitempty"`
	IsSmbSyncWritesEnabled     *bool       `json:"is_smb_sync_writes_enabled,omitempty"`
	IsSmbNoNotifyEnabled       *bool       `json:"is_smb_no_notify_enabled,omitempty"`
	IsSmbOpLocksEnabled        *bool       `json:"is_smb_op_locks_enabled,omitempty"`
	IsSmbNotifyOnAccessEnabled *bool       `json:"is_smb_notify_on_access_enabled,omitempty"`
	IsSmbNotifyOnWriteEnabled  *bool       `json:"is_smb_notify_on_write_enabled,omitempty"`
	SmbNotifyOnChangeDirDepth  int32       `json:"smb_notify_on_change_dir_depth,omitempty"`
	MetaDataHeader
}

type FlrAttributes struct {
	Mode string `json:"mode,omitempty"`
}

const (
	VMware8K  string = "VMware_8K"
	VMware16K string = "VMware_16K"
	VMware32K string = "VMware_32K"
	VMware64K string = "VMware_64K"
)

// MetaData returns the metadata headers.
func (fc *FsCreate) MetaData() http.Header {
	fc.once.Do(func() {
		fc.metadata = make(http.Header)
	})
	return fc.metadata
}

// FSModify modifies existing FS
type FSModify struct {
	// 	integer($int64)
	//minimum: 3221225472
	//maximum: 281474976710656
	//
	//Size, in bytes, presented to the host or end user. This can be used for both expand and shrink on a file system.
	Size                       int           `json:"size_total"`
	Description                string        `json:"description,omitempty"`
	AccessPolicy               string        `json:"access_policy,omitempty"`
	LockingPolicy              string        `json:"locking_policy,omitempty"`
	FolderRenamePolicy         string        `json:"folder_rename_policy,omitempty"`
	IsSmbSyncWritesEnabled     *bool         `json:"is_smb_sync_writes_enabled,omitempty"`
	IsSmbOpLocksEnabled        *bool         `json:"is_smb_op_locks_enabled,omitempty"`
	IsSmbNotifyOnAccessEnabled *bool         `json:"is_smb_notify_on_access_enabled,omitempty"`
	IsSmbNotifyOnWriteEnabled  *bool         `json:"is_smb_notify_on_write_enabled,omitempty"`
	SmbNotifyOnChangeDirDepth  int32         `json:"smb_notify_on_change_dir_depth,omitempty"`
	IsSmbNoNotifyEnabled       *bool         `json:"is_smb_no_notify_enabled,omitempty"`
	IsAsyncMtimeEnabled        *bool         `json:"is_async_MTime_enabled,omitempty"`
	ProtectionPolicyID         string        `json:"protection_policy_id,omitempty"`
	FileEventsPublishingMode   string        `json:"file_events_publishing_mode,omitempty"`
	FlrCreate                  FlrAttributes `json:"flr_attributes,omitempty"`
	ExpirationTimestamp        string        `json:"expiration_timestamp,omitempty"`
}

// NASCreate params for creating 'create nas' request
type NASCreate struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}

// SnapshotFSCreate params for creating 'create snapshot' request
type SnapshotFSCreate struct {
	// Unique name for the snapshot to be created.
	Name string `json:"name,omitempty"`
	// Description of the snapshot.
	Description string `json:"description,omitempty"`
	// Expiration timestamp of the snapshot.
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
	// Access type of the snapshot which can be 'Protocol' / 'Snapshot'
	AccessType string `json:"access_type,omitempty"`
}

// FsClone request for cloning snapshot/fs
type FsClone struct {
	// Unique name for the fs to be created.
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	MetaDataHeader
}

// MetaData returns the metadata headers.
func (fc *FsClone) MetaData() http.Header {
	fc.once.Do(func() {
		fc.metadata = make(http.Header)
	})
	return fc.metadata
}

// Details about the FileSystem
type FileSystem struct {
	// File system id
	ID string `json:"id,omitempty"`
	// File system name
	Name string `json:"name,omitempty"`
	// File system description
	Description string `json:"description,omitempty"`
	// Id of the NAS Server on which the file system is mounted
	NasServerID string `json:"nas_server_id,omitempty"`
	// Type of filesystem: normal or snapshot
	FilesystemType FileSystemTypeEnum `json:"filesystem_type,omitempty"`
	// Size, in bytes, presented to the host or end user
	SizeTotal int64 `json:"size_total,omitempty"`
	// Size used, in bytes, for the data and metadata of the file system
	SizeUsed int64 `json:"size_used,omitempty"`
	// Id of a parent filesystem
	ParentID string `json:"parent_id,omitempty"`
	// Indicates the file system type.
	ConfigType string `json:"config_type,omitempty"`
	// File system security access policies.
	AccessPolicy string `json:"access_policy,omitempty"`
	// [ Native, UNIX, Windows ]
	LockingPolicy string `json:"locking_policy,omitempty"`
	// File system folder rename policies for the file system with multiprotocol access enabled.
	FolderRenamePolicy string `json:"folder_rename_policy,omitempty"`
	// Indicates whether asynchronous MTIME is enabled on the file system
	IsAsyncMTimeEnabled bool `json:"is_async_MTime_enabled,omitempty"`
	// Unique identifier of the protection policy
	ProtectionPolicyID string `json:"protection_policy_id,omitempty"`
	// State of the event notification services for all file systems
	FileEventsPublishingMode string `json:"file_events_publishing_mode,omitempty"`
	// Typical size of writes
	HostIOSize string `json:"host_io_size,omitempty"`
	// Flr attributes
	FlrCreate FlrAttributes `json:"flr_attributes,omitempty"`
	// Indicates whether the synchronous writes option is enabled
	IsSmbSyncWritesEnabled bool `json:"is_smb_sync_writes_enabled,omitempty"`
	// Indicates whether notifications of changes to a directory file structure are enabled.
	IsSmbNoNotifyEnabled bool `json:"is_smb_no_notify_enabled,omitempty"`
	// Indicates whether opportunistic file locking is enabled on the file system.
	IsSmbOpLocksEnabled bool `json:"is_smb_op_locks_enabled,omitempty"`
	// Indicates whether file access notifications are enabled on the file system
	IsSmbNotifyOnAccessEnabled bool `json:"is_smb_notify_on_access_enabled,omitempty"`
	// Indicates whether file writes notifications are enabled on the file system.
	IsSmbNotifyOnWriteEnabled bool `json:"is_smb_notify_on_write_enabled,omitempty"`
	// Lowest directory level to which the enabled notifications apply
	SmbNotifyOnChangeDirDepth int32 `json:"smb_notify_on_change_dir_depth,omitempty"`
	// Expiration timestamp in unix timestamp
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
	// Access type of the file system
	AccessType string `json:"access_type,omitempty"`
}

// NFS server instance in NAS server
type NFSServerInstance struct {
	// Unique identifier for NFS server
	ID string `json:"id"`
	// IsNFSv4Enabled is set to true if nfsv4 is enabled on NAS server
	IsNFSv4Enabled bool `json:"is_nfsv4_enabled,omitempty"`
}

// Details about the NAS.
type NAS struct {
	// Unique identifier of the NAS server.
	ID string `json:"id,omitempty"`
	// Description of the NAS server
	Description string `json:"description,omitempty"`
	// Name of the NAS server
	Name string `json:"name,omitempty"`
	// CurrentNodeId represents on which node the nas server is present
	CurrentNodeID string `json:"current_node_id,omitempty"`
	// NAS server operational status: [ Stopped, Starting, Started, Stopping, Failover, Degraded, Unknown ]
	OperationalStatus NASServerOperationalStatusEnum `json:"operational_status,omitempty"`
	// IPv4 file interface id nas server currently uses
	CurrentPreferredIPv4InterfaceID string `json:"current_preferred_IPv4_interface_id"`
	// NfsServers define NFS server instance if nfs exports are present
	NfsServers []NFSServerInstance `json:"nfs_servers"`
}

// Fields returns fields which must be requested to fill struct
func (n *NAS) Fields() []string {
	return []string{"description", "id", "name", "current_node_id", "operational_status", "current_preferred_IPv4_interface_id", "nfs_servers"}
}

// Fields returns fields which must be requested to fill struct
func (n *FileSystem) Fields() []string {
	return []string{"description", "id", "name", "nas_server_id", "filesystem_type", "size_total", "size_used", "parent_id", "expiration_timestamp", "access_type", "config_type", "access_policy", "locking_policy", "folder_rename_policy", "is_async_MTime_enabled", "protection_policy_id", "file_events_publishing_mode", "host_io_size", "flr_attributes", "is_smb_sync_writes_enabled", "is_smb_no_notify_enabled", "is_smb_op_locks_enabled", "is_smb_notify_on_access_enabled", "is_smb_notify_on_write_enabled", "smb_notify_on_change_dir_depth"}
}

func (n *NFSServerInstance) Fields() []string {
	return []string{"id", "is_nfsv4_enabled"}
}
