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

// MetricsRequest parameters to make metrics request
type MetricsRequest struct {
	Entity   string `json:"entity"`
	EntityID string `json:"entity_id"`
	Interval string `json:"interval"`
}

// ApplianceMetrics is returned by space_metrics_by_appliance metrics request
type ApplianceMetrics struct {
	// Unique identifier of the appliance.
	ApplianceID string `json:"appliance_id"`
	// Total amount of space
	PhysicalTotal int64 `json:"physical_total"`
	// Amount of space currently used
	PhysicalUsed int64 `json:"physical_used"`
}

// Fields returns fields which must be requested to fill struct
func (h *ApplianceMetrics) Fields() []string {
	return []string{"appliance_id", "physical_total", "physical_used"}
}
