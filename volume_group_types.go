/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopowerstore

// VolumeGroupCreate create volume group request
type VolumeGroupCreate struct {
	// Unique name for the volume group.
	// The name should contain no special HTTP characters and no unprintable characters.
	// Although the case of the name provided is reserved, uniqueness check is case-insensitive,
	// so the same name in two different cases is not considered unique.
	Name string `json:"name"`
	// Description for the volume group. The description should not be more than 256
	// characters long and should not have any unprintable characters.
	Description string `json:"description,omitempty"`
	// Unique identifier of an optional protection policy to assign to the volume group.
	ProtectionPolicyID string `json:"protection_policy_id,omitempty"`
	//For a primary or a clone volume group, this property determines whether snapshot sets of the group will be write order consistent.
	IsWriteOrderConsistent bool `json:"is_write_order_consistent"`
	// A list of identifiers of existing volumes that should be added to the volume group.
	// All the volumes must be on the same Cyclone appliance and should not be part of another volume group.
	// If a list of volumes is not specified or if the specified list is empty, an
	// empty volume group of type Volume will be created.
	VolumeIds []string `json:"volume_ids,omitempty"`
}

// VolumeGroup details about a volume groups.
type VolumeGroup struct {
	// Unique identifier of the volume group.
	ID string `json:"id,omitempty"`
	// Name of the volume group.
	// This property supports case-insensitive filtering
	Name string `json:"name,omitempty"`
	// Description for the volume group.
	Description string `json:"description,omitempty"`
	// Unique identifier of the protection policy assigned to the volume.
	ProtectionPolicyID string `json:"protection_policy_id,omitempty"`
	//For a primary or a clone volume group, this property determines whether snapshot sets of the group will be write order consistent.
	IsWriteOrderConsistent bool `json:"is_write_order_consistent,omitempty"`
	// Volumes provides list of volumes associated to the volume group
	Volumes []Volume `json:"volume"`
	// ProtectionPolicy provides snapshot details of the volume or volumeGroup
	ProtectionPolicy ProtectionPolicy `json:"protection_policy"`
	// CreationTimeStamp provides volume group creation time
	CreationTimeStamp string `json:"creation_timestamp,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (v *VolumeGroup) Fields() []string {
	return []string{"id", "name", "description", "protection_policy_id", "creation_timestamp", "is_write_order_consistent"}
}

type VolumeGroups struct {
	VolumeGroup []VolumeGroup `json:"volume_group,omitempty"`
}

type VolumeGroupMembers struct {
	VolumeIds []string `json:"volume_ids"`
}

// VolumeGroupModify modifies existing Volume Group
type VolumeGroupModify struct {
	// empty to delete
	ProtectionPolicyId     string `json:"protection_policy_id"`
	Description            string `json:"description,omitempty"`
	Name                   string `json:"name,omitempty"`
	IsWriteOrderConsistent bool   `json:"is_write_order_consistent"`
}

type VolumeGroupChangePolicy struct {
	ProtectionPolicyID string `json:"protection_policy_id"`
}

// VolumeGroupSnapshotCreate create volume group snapshot request
type VolumeGroupSnapshotCreate struct {
	// Unique name for the volume group.
	Name string `json:"name"`
	// Optional description
	Description string `json:"description,omitempty"`
}
