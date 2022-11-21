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

// RemoteSystem details about a remote system
type RemoteSystem struct {
	// Unique identifier of the remote system instance.
	ID string `json:"id,omitempty"`
	// User-specified name of the remote system instance.
	// This property supports case-insensitive filtering
	Name string `json:"name,omitempty"`
	// User-specified description of the remote system instance.
	Description string `json:"description,omitempty"`
	// Serial number of the remote system instance
	SerialNumber string `json:"serial_number,omitempty"`
	// Management IP address of the remote system instance
	ManagementAddress string `json:"management_address,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (r *RemoteSystem) Fields() []string {
	return []string{"id", "name", "description", "serial_number", "management_address"}
}

// Cluster details about the cluster
type Cluster struct {
	// Unique identifier of the cluster.
	ID string `json:"id,omitempty"`
	// User-specified name of the cluster
	Name string `json:"name,omitempty"`
	// Management IP address of the remote system instance
	ManagementAddress string `json:"management_address,omitempty"`
	// Current state of the cluster
	State string `json:"state,omitempty"`
	// NVMe Subsystem NQN for cluster
	NVMeNQN string `json:"nvm_subsystem_nqn,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (r *Cluster) Fields() []string {
	return []string{"id", "name", "management_address", "state"}
}
