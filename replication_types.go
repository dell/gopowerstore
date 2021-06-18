package gopowerstore

import "errors"

type RPOEnum string

const (
	RpoFiveMinutes    RPOEnum = "Five_Minutes"
	RpoFifteenMinutes RPOEnum = "Fifteen_Minutes"
	RpoThirtyMinutes  RPOEnum = "Thirty_Minutes"
	RpoOneHour        RPOEnum = "One_Hour"
	RpoSixHours       RPOEnum = "Six_Hours"
	RpoTwelveHours    RPOEnum = "Twelve_Hours"
	RpoOneDay         RPOEnum = "One_Day"
)

func (rpo RPOEnum) IsValid() error {
	switch rpo {
	case RpoFiveMinutes, RpoFifteenMinutes, RpoThirtyMinutes, RpoOneHour, RpoSixHours, RpoTwelveHours, RpoOneDay:
		return nil
	}
	return errors.New("invalid rpo type")
}

// ReplicationRuleCreate create replication rule request
type ReplicationRuleCreate struct {
	// Name of the replication rule.
	Name string `json:"name"`
	// Recovery point objective (RPO), which is the acceptable amount of data, measured in units of time, that may be lost in case of a failure.
	Rpo RPOEnum `json:"rpo"`
	// Unique identifier of the remote system to which this rule will replicate the associated resources
	RemoteSystemID string `json:"remote_system_id"`
}

type ReplicationRule struct {
	// ID of replication rule
	ID string `json:"id"`
	// Name of replication rule
	Name string `json:"name"`
	// Rpo (Recovery point objective), which is the acceptable amount of data, measured in units of time, that may be lost in case of a failure.
	Rpo RPOEnum `json:"rpo"`
	// RemoteSystemID - unique identifier of the remote system to which this rule will replicate the associated resources.
	RemoteSystemID string `json:"remote_system_id"`
}

func (rule *ReplicationRule) Fields() []string {
	return []string{"id", "name", "rpo", "remote_system_id"}
}

// ProtectionPolicyCreate create protection policy request
type ProtectionPolicyCreate struct {
	// Policy name.
	Name string `json:"name"`
	// Policy description.
	Description string `json:"description,omitempty"`
	// IDs of replication rules
	ReplicationRuleIds []string `json:"replication_rule_ids,omitempty"`
	// IDs of snapshot rules
	SnapshotRuleIds []string `json:"snapshot_rule_ids,omitempty"`
}

type ProtectionPolicy struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	ReplicationRules []ReplicationRule `json:"replication_rules"`
}

func (policy *ProtectionPolicy) Fields() []string {
	return []string{"id", "name", "replication_rules"}
}

type StorageElementPair struct {
	LocalStorageElementId  string `json:"local_storage_element_id,omitempty"`
	RemoteStorageElementId string `json:"remote_storage_element_id,omitempty"`
	StorageElementType     string `json:"storage_element_type,omitempty"`
	ReplicationShadowId    string `json:"replication_shadow_id,omitempty"`
}

type ReplicationSession struct {
	ID               string `json:"id,omitempty"`
	State            string `json:"state,omitempty"`
	Role             string `json:"role,omitempty"`
	ResourceType     string `json:"resource_type,omitempty"`
	LocalResourceId  string `json:"local_resource_id,omitempty"`
	RemoteResourceId string `json:"remote_resource_id,omitempty"`
	RemoteSystemId   string `json:"remote_system_id,omitempty"` // todo: maybe name?

	StorageElementPairs []StorageElementPair `json:"storage_element_pairs,omitempty"`
}

func (r *ReplicationSession) Fields() []string {
	return []string{"id", "state", "role", "resource_type", "local_resource_id", "remote_resource_id", "remote_system_id", "storage_element_pairs"}
}
