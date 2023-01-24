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

// OSTypeEnum Operating system of the host.
type OSTypeEnum string

const (
	// OSTypeEnumWindows captures enum value "Windows"
	OSTypeEnumWindows OSTypeEnum = "Windows"
	// OSTypeEnumLinux captures enum value "Linux"
	OSTypeEnumLinux OSTypeEnum = "Linux"
	// OSTypeEnumESXi captures enum value "ESXi"
	OSTypeEnumESXi OSTypeEnum = "ESXi"
)

// InitiatorProtocolTypeEnum  Protocol type of the host initiator.
//   - iSCSI - An iSCSI initiator.
//   - FC - A Fibre Channel initiator.
type InitiatorProtocolTypeEnum string

const (
	// InitiatorProtocolTypeEnumISCSI captures enum value "iSCSI"
	InitiatorProtocolTypeEnumISCSI InitiatorProtocolTypeEnum = "iSCSI"
	// InitiatorProtocolTypeEnumNVME captures enum value "NVMe"
	InitiatorProtocolTypeEnumNVME InitiatorProtocolTypeEnum = "NVMe"
	// InitiatorProtocolTypeEnumFC captures enum value "FC"
	InitiatorProtocolTypeEnumFC InitiatorProtocolTypeEnum = "FC"
)

// ActiveSessionInstance active session instance
type ActiveSessionInstance struct {
	// Unique identifier of the appliance containing the session.
	ApplianceID string `json:"appliance_id,omitempty"`
	// Unique identifier of the bond the initiator is logged into.
	// Null if one of the following is non-null: veth_id, eth_port_id or fc_port_id.
	BondID string `json:"bond_id,omitempty"`
	// Unique identifier of the Ethernet port the initiator is logged into.
	// Null if one of the following is non-null: bond_id, veth_id or fc_port_id.
	EthPortID string `json:"eth_port_id,omitempty"`
	// Unique identifier of the FC port the initiator is logged into.
	// Null if one of the following is non-null: bond_id, veth_id or eth_port_id
	FcPortID string `json:"fc_port_id,omitempty"`
	// Unique identifier of node on the appliance on which active session is create.
	NodeID string `json:"node_id,omitempty"`
	// IQN or WWN of the target port that the initiator is logged into.
	PortName string `json:"port_name,omitempty"`
	// Unique identifier of the virtual Ethernet port the initiator is logged into.
	// Null if one of the following is non-null: bond, eth_port_id or fc_port_id.
	VethID string `json:"veth_id,omitempty"`
}

// InitiatorInstance initiator instance
type InitiatorInstance struct {
	// Array of active login session between an initiator and a target port.
	ActiveSessions []ActiveSessionInstance `json:"active_sessions"`
	// Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters.
	// This password is not queriable. CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapMutualPassword string `json:"chap_mutual_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapMutualUsername string `json:"chap_mutual_username,omitempty"`
	// Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters.
	// This password is not queriable. CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapSinglePassword string `json:"chap_single_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapSingleUsername string `json:"chap_single_username,omitempty"`
	// IQN name aka address.
	PortName string `json:"port_name,omitempty"`
	// port type
	PortType InitiatorProtocolTypeEnum `json:"port_type,omitempty"`
}

// InitiatorCreateModify initiator create modify
type InitiatorCreateModify struct {
	// Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters.
	// This password is not queriable. CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapMutualPassword *string `json:"chap_mutual_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapMutualUsername *string `json:"chap_mutual_username,omitempty"`
	// Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters.
	// This password is not queriable. CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapSinglePassword *string `json:"chap_single_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapSingleUsername *string `json:"chap_single_username,omitempty"`
	// IQN name aka address.
	// Required: true
	PortName *string `json:"port_name"`
	// port type
	// Required: true
	PortType *InitiatorProtocolTypeEnum `json:"port_type"`
}

// HostDelete request
type HostDelete struct {
	// Normally, this operation is not allowed on host types other than external.
	// This flag will override that error and allow the operation to continue.
	ForceInternal *bool `json:"force_internal,omitempty"`
}

// HostCreate request
type HostCreate struct {
	// An optional description for the host. The description should not be more than 256 UTF-8
	// characters long and should not have any unprintable characters.
	Description *string `json:"description,omitempty"`
	// Normally, this operation is not allowed on host types other than external.
	// This flag will override that error and allow the operation to continue.
	ForceInternal *bool `json:"force_internal,omitempty"`
	// initiator
	Initiators *[]InitiatorCreateModify `json:"initiators"`
	// The host name. The name should not be more than 128 UTF-8 characters long
	// and should not have any unprintable characters.
	Name *string `json:"name"`
	// os type
	OsType *OSTypeEnum `json:"os_type"`
	// Metadata addition for Hosts on array with OE version 3.0 and above
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// UpdateInitiatorInHost update initiator in host
type UpdateInitiatorInHost struct {
	// Password for CHAP authentication. This value must be 12 to 64 printable UTF-8 characters.
	// CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapMutualPassword *string `json:"chap_mutual_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 printable UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapMutualUsername *string `json:"chap_mutual_username,omitempty"`
	// Password for CHAP authentication. This value must be 12 to 64 printable UTF-8 characters.
	// CHAP password is required when the cluster CHAP mode is mutual authentication.
	ChapSinglePassword *string `json:"chap_single_password,omitempty"`
	// Username for CHAP authentication. This value must be 1 to 64 printable UTF-8 characters.
	// CHAP username is required when the cluster CHAP mode is mutual authentication.
	ChapSingleUsername *string `json:"chap_single_username,omitempty"`
	// Initiator name
	PortName *string `json:"port_name,omitempty"`
}

// HostModify host modify
type HostModify struct {
	// The list of initiators to be added. CHAP username and password are optional.
	AddInitiators *[]InitiatorCreateModify `json:"add_initiators,omitempty"`
	// An optional description for the host.
	// The description should not be more than 256 UTF-8 characters long and should not have any unprintable characters.
	Description *string `json:"description,omitempty"`
	// Normally, this operation is not allowed on host types other than external.
	// This flag will override that error and allow the operation to continue.
	ForceInternal *bool `json:"force_internal,omitempty"`
	// Update list of existing initiators, identified by port_name, with new CHAP usernames and/or passwords.
	ModifyInitiators *[]UpdateInitiatorInHost `json:"modify_initiators,omitempty"`
	// The host name. The name should not be more than 128 UTF-8 characters long and should not have any unprintable characters.
	Name *string `json:"name,omitempty"`
	// The list of initiator port_names to be removed.
	RemoveInitiators *[]string `json:"remove_initiators,omitempty"`
}

// Host host instance
type Host struct {
	// A description for the host.
	Description string `json:"description,omitempty"`
	// Associated host group, if host is part of host group.
	HostGroupID string `json:"host_group_id,omitempty"`
	// Unique id of the host.
	ID string `json:"id,omitempty"`
	// initiators
	Initiators []InitiatorInstance `json:"host_initiators"`
	// The host name.
	Name string `json:"name,omitempty"`
	// os type
	OsType OSTypeEnum `json:"os_type,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (h *Host) Fields() []string {
	return []string{"id", "name", "description", "host_group_id",
		"os_type", "host_initiators"}
}

// HostVolumeMapping Details about a configured host or host group attached to a volume.
// The host or host group may not necessarily be connected.
type HostVolumeMapping struct {
	Volume struct {
		ApplianceID string `json:"appliance_id,omitempty"`
	} `json:"volume,omitempty"`

	// Unique identifier of a host group attached to a volume. The host_id and host_group_id cannot both be set.
	HostGroupID string `json:"host_group_id,omitempty"`
	// Unique identifier of a host attached to a volume. The host_id and host_group_id cannot both be set.
	HostID string `json:"host_id,omitempty"`
	// Unique identifier of a mapping between a host and a volume.
	ID string `json:"id,omitempty"`
	// Logical unit number for the host volume access.
	LogicalUnitNumber int64 `json:"logical_unit_number,omitempty"`
	// Unique identifier of the volume to which the host is attached.
	VolumeID string `json:"volume_id,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (h *HostVolumeMapping) Fields() []string {
	return []string{"volume(appliance_id)", "host_group_id", "host_id",
		"id", "logical_unit_number", "volume_id"}
}

// HostVolumeAttach Volume id and optional logical unit number for attaching to host.
type HostVolumeAttach struct {
	// Logical unit number for the volume, if desired.
	LogicalUnitNumber *int64 `json:"logical_unit_number,omitempty"`
	// Volume to attach.
	VolumeID *string `json:"volume_id"`
}

// HostVolumeDetach Volume id for detaching from host.
type HostVolumeDetach struct {
	// Volume to detach.
	VolumeID *string `json:"volume_id"`
}
