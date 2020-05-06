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

// ApplianceTypeEnum Appliance types.
type ApplianceTypeEnum string

const (
	// ApplianceTypeEnumPowerStore - A storage appliance.
	ApplianceTypeEnumPowerStore ApplianceTypeEnum = "PowerStore"
	// ApplianceTypeEnumPowerStoreX - A storage and compute appliance.
	ApplianceTypeEnumPowerStoreX ApplianceTypeEnum = "PowerStoreX"
)

// ApplianceModeEnum Storage access modes supported by appliance.
type ApplianceModeEnum string

const (
	// ApplianceModeEnumUnified Both block and file storage are supported.
	ApplianceModeEnumUnified ApplianceModeEnum = "Unified"
	// ApplianceModeEnumBlock - Block storage only is supported.
	ApplianceModeEnumBlock ApplianceModeEnum = "Block"
)

// Appliance instance
type Appliance struct {
	// Unique identifier of the appliance.
	ID string `json:"id"`
	// Name of the appliance.
	Name string `json:"name"`
	// IP address value, in IPv4 or IPv6 format.
	IPAddress string `json:"ip_address"`
	// Appliance types.
	Type ApplianceTypeEnum `json:"appliance_type"`
	// Storage access modes supported by appliance.
	Mode ApplianceModeEnum `json:"mode"`
	//
	LastPhysicalTotalSpace int64 `json:"last_physical_total_space"`
	//
	LastPhysicalUsedSpace int64 `json:"last_physical_used_space"`
}

// Fields returns fields which must be requested to fill struct
func (h *Appliance) Fields() []string {
	return []string{"id", "name", "ip_address", "appliance_type",
		"mode", "last_physical_total_space", "last_physical_used_space"}
}
