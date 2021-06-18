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
}

// Fields returns fields which must be requested to fill struct
func (v *VolumeGroup) Fields() []string {
	return []string{"id", "name", "description", "protection_policy_id"}
}

type VolumeGroups struct {
	VolumeGroup []VolumeGroup `json:"volume_group,omitempty"`
}

type VolumeGroupRemoveMember struct {
	VolumeIds []string `json:"volume_ids"`
}

// VolumeGroupModify modifies existing Volume Group
type VolumeGroupModify struct {
	// empty to delete
	ProtectionPolicyId string `json:"protection_policy_id"`
	Description        string `json:"description,omitempty"`
}

type VolumeGroupChangePolicy struct {
	ProtectionPolicyID string `json:"protection_policy_id"`
}
