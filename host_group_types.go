/*
 *
 * Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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

// HostGroup hostgroup instance
type HostGroup struct {
	// A description for the hostgroup.
	Description string `json:"description,omitempty"`
	// Unique id of the hostgroup.
	ID string `json:"id,omitempty"`
	// The hostgroup name.
	Name string `json:"name,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (h *HostGroup) Fields() []string {
	return []string{"id", "name", "description"}
}
