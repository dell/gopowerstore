/*
 *
 * Copyright 2020 Dell EMC Corporation
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

// FsCreate params for creating 'create fs' request
type FsCreate struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	NASServerID string `json:"nas_server_id"`
	Size        int64  `json:"size_total"`
	MetaDataHeader
}

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
	Size        int    `json:"size_total"`
	Description string `json:"description,omitempty"`
}

// NASCreate params for creating 'create nas' request
type NASCreate struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}

// SnapshotFSCreate params for creating 'create snapshot' request
type SnapshotFSCreate struct {
	// Unique name for the snapshot to be created.
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
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
	ParentId string `json:"parent_id,omitempty"`
}

// Details about the NAS.
type NAS struct {
	// Unique identifier of the NAS server.
	ID string `json:"id,omitempty"`
	// Description of the NAS server
	Description string `json:"description,omitempty"`
	// Name of the NAS server
	Name string `json:"name,omitempty"`
	// NAS server operational status: [ Stopped, Starting, Started, Stopping, Failover, Degraded, Unknown ]
	OperationalStatus NASServerOperationalStatusEnum `json:"operational_status,omitempty"`
	// IPv4 file interface id nas server currently uses
	CurrentPreferredIPv4InterfaceId string `json:"current_preferred_IPv4_interface_id"`
}

// Fields returns fields which must be requested to fill struct
func (n *NAS) Fields() []string {
	return []string{"description", "id", "name", "operational_status", "current_preferred_IPv4_interface_id"}
}

// Fields returns fields which must be requested to fill struct
func (n *FileSystem) Fields() []string {
	return []string{"description", "id", "name", "nas_server_id", "filesystem_type", "size_total", "size_used", "parent_id"}
}
