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

import "errors"

type RPOEnum string
type RSStateEnum string

const (
	RpoFiveMinutes                         RPOEnum     = "Five_Minutes"
	RpoFifteenMinutes                      RPOEnum     = "Fifteen_Minutes"
	RpoThirtyMinutes                       RPOEnum     = "Thirty_Minutes"
	RpoOneHour                             RPOEnum     = "One_Hour"
	RpoSixHours                            RPOEnum     = "Six_Hours"
	RpoTwelveHours                         RPOEnum     = "Twelve_Hours"
	RpoOneDay                              RPOEnum     = "One_Day"
	RS_STATE_INITIALIZING                  RSStateEnum = "Initializing"
	RS_STATE_OK                            RSStateEnum = "OK"
	RS_STATE_SYNCHRONIZING                 RSStateEnum = "Synchronizing"
	RS_STATE_SYSTEM_PAUSED                 RSStateEnum = "System_Paused"
	RS_STATE_PAUSED                        RSStateEnum = "Paused"
	RS_STATE_PAUSED_FOR_MIGRATION          RSStateEnum = "Paused_For_Migration"
	RS_STATE_PAUSED_FOR_NDU                RSStateEnum = "Paused_For_NDU"
	RS_STATE_RESUMING                      RSStateEnum = "Resuming"
	RS_STATE_FAILING_OVER                  RSStateEnum = "Failing_Over"
	RS_STATE_FAILING_OVER_FOR_DR           RSStateEnum = "Failing_Over_For_DR"
	RS_STATE_FAILED_OVER                   RSStateEnum = "Failed_Over"
	RS_STATE_REPROTECTING                  RSStateEnum = "Reprotecting"
	RS_STATE_PARTIAL_CUTOVER_FOR_MIGRATION RSStateEnum = "Partial_Cutover_For_Migration"
	RS_STATE_ERROR                         RSStateEnum = "Error"
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
	RemoteSystemID     string             `json:"remote_system_id"`
	ProtectionPolicies []ProtectionPolicy `json:"policies"`
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

// failover params create failover request
type FailoverParams struct {
	// For DR failover.
	IsPlanned bool `json:"is_planned, omitempty"`
	// Reverse replication direction
	Reverse bool `json:"reverse,omitempty"`
	// Force for DR
	Force bool `json:"force,omitempty"`
}

type ProtectionPolicy struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	ReplicationRules []ReplicationRule `json:"replication_rules"`
	Volumes          []Volume          `json:"volume"`
	VolumeGroups     []VolumeGroup     `json:"volume_group"`
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
	ID               string      `json:"id,omitempty"`
	State            RSStateEnum `json:"state,omitempty"`
	Role             string      `json:"role,omitempty"`
	ResourceType     string      `json:"resource_type,omitempty"`
	LocalResourceId  string      `json:"local_resource_id,omitempty"`
	RemoteResourceId string      `json:"remote_resource_id,omitempty"`
	RemoteSystemId   string      `json:"remote_system_id,omitempty"` // todo: maybe name?

	StorageElementPairs []StorageElementPair `json:"storage_element_pairs,omitempty"`
}

func (r *ReplicationSession) Fields() []string {
	return []string{"id", "state", "role", "resource_type", "local_resource_id", "remote_resource_id", "remote_system_id", "storage_element_pairs"}
}
