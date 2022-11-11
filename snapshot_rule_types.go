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

// SnapshotRuleIntervalEnum - Interval between snapshots taken by a snapshot rule.
type SnapshotRuleIntervalEnum string

// SnapshotRuleIntervalEnum known intervals
const (
	SnapshotRuleIntervalEnumFive_Minutes    SnapshotRuleIntervalEnum = "Five_Minutes"
	SnapshotRuleIntervalEnumFifteen_Minutes SnapshotRuleIntervalEnum = "Fifteen_Minutes"
	SnapshotRuleIntervalEnumThirty_Minutes  SnapshotRuleIntervalEnum = "Thirty_Minutes"
	SnapshotRuleIntervalEnumOne_Hour        SnapshotRuleIntervalEnum = "One_Hour"
	SnapshotRuleIntervalEnumTwo_Hours       SnapshotRuleIntervalEnum = "Two_Hours"
	SnapshotRuleIntervalEnumThree_Hours     SnapshotRuleIntervalEnum = "Three_Hours"
	SnapshotRuleIntervalEnumFour_Hours      SnapshotRuleIntervalEnum = "Four_Hours"
	SnapshotRuleIntervalEnumSix_Hours       SnapshotRuleIntervalEnum = "Six_Hours"
	SnapshotRuleIntervalEnumEight_Hours     SnapshotRuleIntervalEnum = "Eight_Hours"
	SnapshotRuleIntervalEnumTwelve_Hours    SnapshotRuleIntervalEnum = "Twelve_Hours"
	SnapshotRuleIntervalEnumOne_Day         SnapshotRuleIntervalEnum = "One_Day"
)

// TimeZoneEnum defines identifier for timezone
type TimeZoneEnum string

// TimeZoneEnum known timezones
const (
	TimeZoneEnumEtc__GMT_plus_12        TimeZoneEnum = "Etc__GMT_plus_12"
	TimeZoneEnumUS__Samoa               TimeZoneEnum = "US__Samoa"
	TimeZoneEnumEtc__GMT_plus_11        TimeZoneEnum = "Etc__GMT_plus_11"
	TimeZoneEnumAmerica__Atka           TimeZoneEnum = "America__Atka"
	TimeZoneEnumUS__Hawaii              TimeZoneEnum = "US__Hawaii"
	TimeZoneEnumEtc__GMT_plus_10        TimeZoneEnum = "Etc__GMT_plus_10"
	TimeZoneEnumPacific__Marquesas      TimeZoneEnum = "Pacific__Marquesas"
	TimeZoneEnumUS__Alaska              TimeZoneEnum = "US__Alaska"
	TimeZoneEnumPacific__Gambier        TimeZoneEnum = "Pacific__Gambier"
	TimeZoneEnumEtc__GMT_plus_9         TimeZoneEnum = "Etc__GMT_plus_9"
	TimeZoneEnumPST8PDT                 TimeZoneEnum = "PST8PDT"
	TimeZoneEnumPacific__Pitcairn       TimeZoneEnum = "Pacific__Pitcairn"
	TimeZoneEnumUS__Pacific             TimeZoneEnum = "US__Pacific"
	TimeZoneEnumEtc__GMT_plus_8         TimeZoneEnum = "Etc__GMT_plus_8"
	TimeZoneEnumMexico__BajaSur         TimeZoneEnum = "Mexico__BajaSur"
	TimeZoneEnumAmerica__Boise          TimeZoneEnum = "America__Boise"
	TimeZoneEnumAmerica__Phoenix        TimeZoneEnum = "America__Phoenix"
	TimeZoneEnumMST7MDT                 TimeZoneEnum = "MST7MDT"
	TimeZoneEnumEtc__GMT_plus_7         TimeZoneEnum = "Etc__GMT_plus_7"
	TimeZoneEnumCST6CDT                 TimeZoneEnum = "CST6CDT"
	TimeZoneEnumAmerica__Chicago        TimeZoneEnum = "America__Chicago"
	TimeZoneEnumCanada__Saskatchewan    TimeZoneEnum = "Canada__Saskatchewan"
	TimeZoneEnumAmerica__Bahia_Banderas TimeZoneEnum = "America__Bahia_Banderas"
	TimeZoneEnumEtc__GMT_plus_6         TimeZoneEnum = "Etc__GMT_plus_6"
	TimeZoneEnumChile__EasterIsland     TimeZoneEnum = "Chile__EasterIsland"
	TimeZoneEnumAmerica__Bogota         TimeZoneEnum = "America__Bogota"
	TimeZoneEnumAmerica__New_York       TimeZoneEnum = "America__New_York"
	TimeZoneEnumEST5EDT                 TimeZoneEnum = "EST5EDT"
	TimeZoneEnumAmerica__Havana         TimeZoneEnum = "America__Havana"
	TimeZoneEnumEtc__GMT_plus_5         TimeZoneEnum = "Etc__GMT_plus_5"
	TimeZoneEnumAmerica__Caracas        TimeZoneEnum = "America__Caracas"
	TimeZoneEnumAmerica__Cuiaba         TimeZoneEnum = "America__Cuiaba"
	TimeZoneEnumAmerica__Santo_Domingo  TimeZoneEnum = "America__Santo_Domingo"
	TimeZoneEnumCanada__Atlantic        TimeZoneEnum = "Canada__Atlantic"
	TimeZoneEnumAmerica__Asuncion       TimeZoneEnum = "America__Asuncion"
	TimeZoneEnumEtc__GMT_plus_4         TimeZoneEnum = "Etc__GMT_plus_4"
	TimeZoneEnumCanada__Newfoundland    TimeZoneEnum = "Canada__Newfoundland"
	TimeZoneEnumChile__Continental      TimeZoneEnum = "Chile__Continental"
	TimeZoneEnumBrazil__East            TimeZoneEnum = "Brazil__East"
	TimeZoneEnumAmerica__Godthab        TimeZoneEnum = "America__Godthab"
	TimeZoneEnumAmerica__Miquelon       TimeZoneEnum = "America__Miquelon"
	TimeZoneEnumAmerica__Buenos_Aires   TimeZoneEnum = "America__Buenos_Aires"
	TimeZoneEnumEtc__GMT_plus_3         TimeZoneEnum = "Etc__GMT_plus_3"
	TimeZoneEnumAmerica__Noronha        TimeZoneEnum = "America__Noronha"
	TimeZoneEnumEtc__GMT_plus_2         TimeZoneEnum = "Etc__GMT_plus_2"
	TimeZoneEnumAmerica__Scoresbysund   TimeZoneEnum = "America__Scoresbysund"
	TimeZoneEnumAtlantic__Cape_Verde    TimeZoneEnum = "Atlantic__Cape_Verde"
	TimeZoneEnumEtc__GMT_plus_1         TimeZoneEnum = "Etc__GMT_plus_1"
	TimeZoneEnumUTC                     TimeZoneEnum = "UTC"
	TimeZoneEnumEurope__London          TimeZoneEnum = "Europe__London"
	TimeZoneEnumAfrica__Casablanca      TimeZoneEnum = "Africa__Casablanca"
	TimeZoneEnumAtlantic__Reykjavik     TimeZoneEnum = "Atlantic__Reykjavik"
	TimeZoneEnumAntarctica__Troll       TimeZoneEnum = "Antarctica__Troll"
	TimeZoneEnumEurope__Paris           TimeZoneEnum = "Europe__Paris"
	TimeZoneEnumEurope__Sarajevo        TimeZoneEnum = "Europe__Sarajevo"
	TimeZoneEnumEurope__Belgrade        TimeZoneEnum = "Europe__Belgrade"
	TimeZoneEnumEurope__Rome            TimeZoneEnum = "Europe__Rome"
	TimeZoneEnumAfrica__Tunis           TimeZoneEnum = "Africa__Tunis"
	TimeZoneEnumEtc__GMT_minus_1        TimeZoneEnum = "Etc__GMT_minus_1"
	TimeZoneEnumAsia__Gaza              TimeZoneEnum = "Asia__Gaza"
	TimeZoneEnumEurope__Bucharest       TimeZoneEnum = "Europe__Bucharest"
	TimeZoneEnumEurope__Helsinki        TimeZoneEnum = "Europe__Helsinki"
	TimeZoneEnumAsia__Beirut            TimeZoneEnum = "Asia__Beirut"
	TimeZoneEnumAfrica__Harare          TimeZoneEnum = "Africa__Harare"
	TimeZoneEnumAsia__Damascus          TimeZoneEnum = "Asia__Damascus"
	TimeZoneEnumAsia__Amman             TimeZoneEnum = "Asia__Amman"
	TimeZoneEnumEurope__Tiraspol        TimeZoneEnum = "Europe__Tiraspol"
	TimeZoneEnumAsia__Jerusalem         TimeZoneEnum = "Asia__Jerusalem"
	TimeZoneEnumEtc__GMT_minus_2        TimeZoneEnum = "Etc__GMT_minus_2"
	TimeZoneEnumAsia__Baghdad           TimeZoneEnum = "Asia__Baghdad"
	TimeZoneEnumAfrica__Asmera          TimeZoneEnum = "Africa__Asmera"
	TimeZoneEnumEtc__GMT_minus_3        TimeZoneEnum = "Etc__GMT_minus_3"
	TimeZoneEnumAsia__Tehran            TimeZoneEnum = "Asia__Tehran"
	TimeZoneEnumAsia__Baku              TimeZoneEnum = "Asia__Baku"
	TimeZoneEnumEtc__GMT_minus_4        TimeZoneEnum = "Etc__GMT_minus_4"
	TimeZoneEnumAsia__Kabul             TimeZoneEnum = "Asia__Kabul"
	TimeZoneEnumAsia__Karachi           TimeZoneEnum = "Asia__Karachi"
	TimeZoneEnumEtc__GMT_minus_5        TimeZoneEnum = "Etc__GMT_minus_5"
	TimeZoneEnumAsia__Kolkata           TimeZoneEnum = "Asia__Kolkata"
	TimeZoneEnumAsia__Katmandu          TimeZoneEnum = "Asia__Katmandu"
	TimeZoneEnumAsia__Almaty            TimeZoneEnum = "Asia__Almaty"
	TimeZoneEnumEtc__GMT_minus_6        TimeZoneEnum = "Etc__GMT_minus_6"
	TimeZoneEnumAsia__Rangoon           TimeZoneEnum = "Asia__Rangoon"
	TimeZoneEnumAsia__Hovd              TimeZoneEnum = "Asia__Hovd"
	TimeZoneEnumAsia__Bangkok           TimeZoneEnum = "Asia__Bangkok"
	TimeZoneEnumEtc__GMT_minus_7        TimeZoneEnum = "Etc__GMT_minus_7"
	TimeZoneEnumAsia__Hong_Kong         TimeZoneEnum = "Asia__Hong_Kong"
	TimeZoneEnumAsia__Brunei            TimeZoneEnum = "Asia__Brunei"
	TimeZoneEnumAsia__Singapore         TimeZoneEnum = "Asia__Singapore"
	TimeZoneEnumEtc__GMT_minus_8        TimeZoneEnum = "Etc__GMT_minus_8"
	TimeZoneEnumAsia__Pyongyang         TimeZoneEnum = "Asia__Pyongyang"
	TimeZoneEnumAustralia__Eucla        TimeZoneEnum = "Australia__Eucla"
	TimeZoneEnumAsia__Seoul             TimeZoneEnum = "Asia__Seoul"
	TimeZoneEnumEtc__GMT_minus_9        TimeZoneEnum = "Etc__GMT_minus_9"
	TimeZoneEnumAustralia__Darwin       TimeZoneEnum = "Australia__Darwin"
	TimeZoneEnumAustralia__Adelaide     TimeZoneEnum = "Australia__Adelaide"
	TimeZoneEnumAustralia__Sydney       TimeZoneEnum = "Australia__Sydney"
	TimeZoneEnumAustralia__Brisbane     TimeZoneEnum = "Australia__Brisbane"
	TimeZoneEnumAsia__Magadan           TimeZoneEnum = "Asia__Magadan"
	TimeZoneEnumEtc__GMT_minus_10       TimeZoneEnum = "Etc__GMT_minus_10"
	TimeZoneEnumAustralia__Lord_Howe    TimeZoneEnum = "Australia__Lord_Howe"
	TimeZoneEnumEtc__GMT_minus_11       TimeZoneEnum = "Etc__GMT_minus_11"
	TimeZoneEnumAsia__Kamchatka         TimeZoneEnum = "Asia__Kamchatka"
	TimeZoneEnumPacific__Fiji           TimeZoneEnum = "Pacific__Fiji"
	TimeZoneEnumAntarctica__South_Pole  TimeZoneEnum = "Antarctica__South_Pole"
	TimeZoneEnumEtc__GMT_minus_12       TimeZoneEnum = "Etc__GMT_minus_12"
	TimeZoneEnumPacific__Chatham        TimeZoneEnum = "Pacific__Chatham"
	TimeZoneEnumPacific__Tongatapu      TimeZoneEnum = "Pacific__Tongatapu"
	TimeZoneEnumPacific__Apia           TimeZoneEnum = "Pacific__Apia"
	TimeZoneEnumEtc__GMT_minus_13       TimeZoneEnum = "Etc__GMT_minus_13"
	TimeZoneEnumPacific__Kiritimati     TimeZoneEnum = "Pacific__Kiritimati"
	TimeZoneEnumEtc__GMT_minus_14       TimeZoneEnum = "Etc__GMT_minus_14"
)

// DaysOfWeekEnum - days of week
type DaysOfWeekEnum string

// DaysOfWeekEnum - known days of week
const (
	DaysOfWeekEnumMonday    DaysOfWeekEnum = "Monday"
	DaysOfWeekEnumTuesday   DaysOfWeekEnum = "Tuesday"
	DaysOfWeekEnumWednesday DaysOfWeekEnum = "Wednesday"
	DaysOfWeekEnumThursday  DaysOfWeekEnum = "Thursday"
	DaysOfWeekEnumFriday    DaysOfWeekEnum = "Friday"
	DaysOfWeekEnumSaturday  DaysOfWeekEnum = "Saturday"
	DaysOfWeekEnumSunday    DaysOfWeekEnum = "Sunday"
)

// NASAccessTypeEnums - NAS filesystem snapshot access method
type NASAccessTypeEnum string

const (
	// NASAccessTypeEnumSnapshot - NAS filesystem snapshot access method - snapshot
	// the files within the snapshot may be access directly from the production file system in the .snapshot subdirectory of each directory.
	NASAccessTypeEnumSnapshot NASAccessTypeEnum = "Snapshot"

	// NASAccessTypeEnumProtocol - NAS filesystem snapshot access method - protocol
	// the entire file system snapshot may be shared and mounted on a client like any other file system, except that it is readonly.
	NASAccessTypeEnumProtocol NASAccessTypeEnum = "Protocol"
)

// PolicyManagedByEnum - defines entities who manage the instance
type PolicyManagedByEnum string

const (
	// PolicyManagedByEnumUser - instance is managed by the end user
	PolicyManagedByEnumUser PolicyManagedByEnum = "User"
	// PolicyManagedByEnumMetro - instance is managed by the peer system where the policy was assigned, in a Metro Cluster configuration
	PolicyManagedByEnumMetro PolicyManagedByEnum = "Metro"
	// PolicyManagedByEnumReplication - destination instance is managed by the source system in a Replication configuration
	PolicyManagedByEnumReplication PolicyManagedByEnum = "Replication"
	// PolicyManagedByEnumVMware_vSphere - instance is managed by the system through VMware vSphere/vCenter
	PolicyManagedByEnumVMware_vSphere PolicyManagedByEnum = "VMware_vSphere"
)

// SnapshotRuleCreate create snapshot rule request
type SnapshotRuleCreate struct {
	// Name of the snapshot rule
	// minLength: 1
	// maxLength: 128
	Name string `json:"name,omitempty"`

	// Interval between snapshots taken by a snapshot rule
	Interval SnapshotRuleIntervalEnum `json:"interval,omitempty"`

	// Time of the day to take a daily snapshot, with format "hh:mm" using a 24 hour clock
	// Either the interval parameter or the time_of_day parameter will be set, but not both.
	TimeOfDay string `json:"time_of_day,omitempty"`

	// Time zone identifier for applying the time zone to the time_of_day for a snapshot rule, including any DST effects if applicable
	// Applies only when a time_of_day is specified in the snapshot rule. Defaults to UTC if not specified.
	// Was added in version 2.0.0.0
	TimeZone TimeZoneEnum `json:"timezone,omitempty"`

	// Days of the week when the snapshot rule should be applied.
	// Days are determined based on the UTC time zone, unless the time_of_day and timezone properties are set.
	DaysOfWeek []DaysOfWeekEnum `json:"days_of_week,omitempty"`

	// Desired snapshot retention period in hours. The system will retain snapshots for this time period.
	// minimum: 0
	// maximum: 8760
	DesiredRetention int32 `json:"desired_retention,omitempty"`

	// NAS filesystem snapshot access method.
	// setting is ignored for volume, virtual_volume, and volume_group snapshots
	NASAccessType NASAccessTypeEnum `json:"nas_access_type,omitempty"`

	// Indicates whether this snapshot rule can be modified.
	// default: false
	IsReadOnly bool `json:"is_read_only,omitempty"`
}

// SnapshotRuleDelete body for SnapshotRuleDelete request
type SnapshotRuleDelete struct {
	// Specify whether all snapshots previously created by this snapshot rule should also be deleted when this rule is removed.
	// default false
	DeleteSnaps bool `json:"delete_snaps,omitempty"`
}

// SnapshotRule Details about a snapshot rule
type SnapshotRule struct {

	// Unique identifier of the snapshot rule
	ID string `json:"id,omitempty"`

	// Snapshot rule name.
	// This property supports case-insensitive filtering.
	Name string `json:"name,omitempty"`

	// Interval between snapshots taken by a snapshot rule.
	Interval SnapshotRuleIntervalEnum `json:"interval,omitempty"`

	// Time of the day to take a daily snapshot, with format "hh:mm" using a 24 hour clock
	// Either the interval parameter or the time_of_day parameter will be set, but not both.
	TimeOfDay string `json:"time_of_day,omitempty"`

	// Time zone identifier for applying the time zone to the time_of_day for a snapshot rule, including any DST effects if applicable
	// Applies only when a time_of_day is specified in the snapshot rule. Defaults to UTC if not specified.
	// Was added in version 2.0.0.0
	TimeZone TimeZoneEnum `json:"timezone,omitempty"`

	// Days of the week when the snapshot rule should be applied.
	// Days are determined based on the UTC time zone, unless the time_of_day and timezone properties are set.
	DaysOfWeek []DaysOfWeekEnum `json:"days_of_week,omitempty"`

	// Desired snapshot retention period in hours. The system will retain snapshots for this time period.
	// minimum: 0
	// maximum: 8760
	DesiredRetention int32 `json:"desired_retention,omitempty"`

	// Indicates whether this is a replica of a snapshot rule on a remote system
	// that is the source of a replication session replicating a storage resource to the local system.
	// defalut : false
	IsReplica bool `json:"is_replica,omitempty"`

	// NAS filesystem snapshot access method.
	// setting is ignored for volume, virtual_volume, and volume_group snapshots
	NASAccessType NASAccessTypeEnum `json:"nas_access_type,omitempty"`

	// Indicates whether this snapshot rule can be modified.
	// default: false
	IsReadOnly bool `json:"is_read_only,omitempty"`

	// entity that owns and manages this instance
	ManagedBy PolicyManagedByEnum `json:"managed_by,omitempty"`

	// 	Unique identifier of the managing entity based on the value of the managed_by property, as shown below:
	//         User - Empty
	//         Metro - Unique identifier of the remote system where the policy was assigned.
	//         Replication - Unique identifier of the source remote system.
	//         VMware_vSphere - Unique identifier of the owning VMware vSphere/vCenter.
	ManagedById string `json:"managed_by_id,omitempty"`

	// Localized message string corresponding to interval
	Interval_l10n string `json:"interval_l10n,omitempty"`

	// Localized message string corresponding to timezone
	Timezone_l10n string `json:"timezone_l10n,omitempty"`

	// Localized message array corresponding to days_of_week
	DaysOfWeek_l10n []string `json:"days_of_week_l10n,omitempty"`

	// Localized message string corresponding to nas_access_type
	NASAccessType_l10n string `json:"nas_access_type_l10n,omitempty"`

	ManagedNy_l10n string `json:"managed_by_l10n,omitempty"`

	// todo Policies []PolicyInstance `json:"policies,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (s *SnapshotRule) Fields() []string {
	return []string{"id", "name",
		"interval", "time_of_day", "timezone", "days_of_week", "desired_retention",
		"is_replica", "nas_access_type", "is_read_only",
		"managed_by", "managed_by_id",
		"interval_l10n", "timezone_l10n", "days_of_week_l10n", "nas_access_type_l10n", "managed_by_l10n",
	}
}
