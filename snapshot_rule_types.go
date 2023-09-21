/*
 *
 * Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	SnapshotRuleIntervalEnumFiveMinutes    SnapshotRuleIntervalEnum = "Five_Minutes"
	SnapshotRuleIntervalEnumFifteenMinutes SnapshotRuleIntervalEnum = "Fifteen_Minutes"
	SnapshotRuleIntervalEnumThirtyMinutes  SnapshotRuleIntervalEnum = "Thirty_Minutes"
	SnapshotRuleIntervalEnumOneHour        SnapshotRuleIntervalEnum = "One_Hour"
	SnapshotRuleIntervalEnumTwoHours       SnapshotRuleIntervalEnum = "Two_Hours"
	SnapshotRuleIntervalEnumThreeHours     SnapshotRuleIntervalEnum = "Three_Hours"
	SnapshotRuleIntervalEnumFourHours      SnapshotRuleIntervalEnum = "Four_Hours"
	SnapshotRuleIntervalEnumSixHours       SnapshotRuleIntervalEnum = "Six_Hours"
	SnapshotRuleIntervalEnumEightHours     SnapshotRuleIntervalEnum = "Eight_Hours"
	SnapshotRuleIntervalEnumTwelveHours    SnapshotRuleIntervalEnum = "Twelve_Hours"
	SnapshotRuleIntervalEnumOneDay         SnapshotRuleIntervalEnum = "One_Day"
)

// TimeZoneEnum defines identifier for timezone
type TimeZoneEnum string

// TimeZoneEnum known timezones
const (
	TimeZoneEnumEtcGMTPlus12         TimeZoneEnum = "EtcGMTPlus12"
	TimeZoneEnumUSSamoa              TimeZoneEnum = "USSamoa"
	TimeZoneEnumEtcGMTPlus11         TimeZoneEnum = "EtcGMTPlus11"
	TimeZoneEnumAmericaAtka          TimeZoneEnum = "AmericaAtka"
	TimeZoneEnumUSHawaii             TimeZoneEnum = "USHawaii"
	TimeZoneEnumEtcGMTPlus10         TimeZoneEnum = "EtcGMTPlus10"
	TimeZoneEnumPacificMarquesas     TimeZoneEnum = "PacificMarquesas"
	TimeZoneEnumUSAlaska             TimeZoneEnum = "USAlaska"
	TimeZoneEnumPacificGambier       TimeZoneEnum = "PacificGambier"
	TimeZoneEnumEtcGMTPlus9          TimeZoneEnum = "EtcGMTPlus9"
	TimeZoneEnumPST8PDT              TimeZoneEnum = "PST8PDT"
	TimeZoneEnumPacificPitcairn      TimeZoneEnum = "PacificPitcairn"
	TimeZoneEnumUSPacific            TimeZoneEnum = "USPacific"
	TimeZoneEnumEtcGMTPlus8          TimeZoneEnum = "EtcGMTPlus8"
	TimeZoneEnumMexicoBajaSur        TimeZoneEnum = "MexicoBajaSur"
	TimeZoneEnumAmericaBoise         TimeZoneEnum = "AmericaBoise"
	TimeZoneEnumAmericaPhoenix       TimeZoneEnum = "AmericaPhoenix"
	TimeZoneEnumMST7MDT              TimeZoneEnum = "MST7MDT"
	TimeZoneEnumEtcGMTPlus7          TimeZoneEnum = "EtcGMTPlus7"
	TimeZoneEnumCST6CDT              TimeZoneEnum = "CST6CDT"
	TimeZoneEnumAmericaChicago       TimeZoneEnum = "AmericaChicago"
	TimeZoneEnumCanadaSaskatchewan   TimeZoneEnum = "CanadaSaskatchewan"
	TimeZoneEnumAmericaBahiaBanderas TimeZoneEnum = "AmericaBahia_Banderas"
	TimeZoneEnumEtcGMTPlus6          TimeZoneEnum = "EtcGMTPlus6"
	TimeZoneEnumChileEasterIsland    TimeZoneEnum = "ChileEasterIsland"
	TimeZoneEnumAmericaBogota        TimeZoneEnum = "AmericaBogota"
	TimeZoneEnumAmericaNewYork       TimeZoneEnum = "AmericaNewYork"
	TimeZoneEnumEST5EDT              TimeZoneEnum = "EST5EDT"
	TimeZoneEnumAmericaHavana        TimeZoneEnum = "AmericaHavana"
	TimeZoneEnumEtcGMTPlus5          TimeZoneEnum = "EtcGMTPlus5"
	TimeZoneEnumAmericaCaracas       TimeZoneEnum = "AmericaCaracas"
	TimeZoneEnumAmericaCuiaba        TimeZoneEnum = "AmericaCuiaba"
	TimeZoneEnumAmericaSantoDomingo  TimeZoneEnum = "AmericaSantoDomingo"
	TimeZoneEnumCanadaAtlantic       TimeZoneEnum = "CanadaAtlantic"
	TimeZoneEnumAmericaAsuncion      TimeZoneEnum = "AmericaAsuncion"
	TimeZoneEnumEtcGMTPlus4          TimeZoneEnum = "EtcGMTPlus4"
	TimeZoneEnumCanadaNewFoundLand   TimeZoneEnum = "CanadaNewfoundland"
	TimeZoneEnumChileContinental     TimeZoneEnum = "ChileContinental"
	TimeZoneEnumBrazilEast           TimeZoneEnum = "BrazilEast"
	TimeZoneEnumAmericaGodthab       TimeZoneEnum = "AmericaGodthab"
	TimeZoneEnumAmericaMiquelon      TimeZoneEnum = "AmericaMiquelon"
	TimeZoneEnumAmericaBuenosAires   TimeZoneEnum = "AmericaBuenosAires"
	TimeZoneEnumEtcGMTPlus3          TimeZoneEnum = "EtcGMTPlus3"
	TimeZoneEnumAmericaNoronha       TimeZoneEnum = "AmericaNoronha"
	TimeZoneEnumEtcGMTPlus2          TimeZoneEnum = "EtcGMTPlus2"
	TimeZoneEnumAmericaScoresbysund  TimeZoneEnum = "AmericaScoresbysund"
	TimeZoneEnumAtlanticCapeVerde    TimeZoneEnum = "AtlanticCapeVerde"
	TimeZoneEnumEtcGMTPlus1          TimeZoneEnum = "EtcGMTPlus1"
	TimeZoneEnumUTC                  TimeZoneEnum = "UTC"
	TimeZoneEnumEuropeLondon         TimeZoneEnum = "EuropeLondon"
	TimeZoneEnumAfricaCasablanca     TimeZoneEnum = "AfricaCasablanca"
	TimeZoneEnumAtlanticReykjavik    TimeZoneEnum = "AtlanticReykjavik"
	TimeZoneEnumAntarcticaTroll      TimeZoneEnum = "AntarcticaTroll"
	TimeZoneEnumEuropeParis          TimeZoneEnum = "EuropeParis"
	TimeZoneEnumEuropeSarajevo       TimeZoneEnum = "EuropeSarajevo"
	TimeZoneEnumEuropeBelgrade       TimeZoneEnum = "EuropeBelgrade"
	TimeZoneEnumEuropeRome           TimeZoneEnum = "EuropeRome"
	TimeZoneEnumAfricaTunis          TimeZoneEnum = "AfricaTunis"
	TimeZoneEnumEtcGMTMinus1         TimeZoneEnum = "EtcGMTMinus1"
	TimeZoneEnumAsiaGaza             TimeZoneEnum = "AsiaGaza"
	TimeZoneEnumEuropeBucharest      TimeZoneEnum = "EuropeBucharest"
	TimeZoneEnumEuropeHelsinki       TimeZoneEnum = "EuropeHelsinki"
	TimeZoneEnumAsiaBeirut           TimeZoneEnum = "AsiaBeirut"
	TimeZoneEnumAfricaHarare         TimeZoneEnum = "AfricaHarare"
	TimeZoneEnumAsiaDamascus         TimeZoneEnum = "AsiaDamascus"
	TimeZoneEnumAsiaAmman            TimeZoneEnum = "AsiaAmman"
	TimeZoneEnumEuropeTiraspol       TimeZoneEnum = "EuropeTiraspol"
	TimeZoneEnumAsiaJerusalem        TimeZoneEnum = "AsiaJerusalem"
	TimeZoneEnumEtcGMTMinus2         TimeZoneEnum = "EtcGMTMinus2"
	TimeZoneEnumAsiaBaghdad          TimeZoneEnum = "AsiaBaghdad"
	TimeZoneEnumAfricaAsmera         TimeZoneEnum = "AfricaAsmera"
	TimeZoneEnumEtcGMTMinus3         TimeZoneEnum = "EtcGMTMinus3"
	TimeZoneEnumAsiaTehran           TimeZoneEnum = "AsiaTehran"
	TimeZoneEnumAsiaBaku             TimeZoneEnum = "AsiaBaku"
	TimeZoneEnumEtcGMTMinus4         TimeZoneEnum = "EtcGMTMinus4"
	TimeZoneEnumAsiaKabul            TimeZoneEnum = "AsiaKabul"
	TimeZoneEnumAsiaKarachi          TimeZoneEnum = "AsiaKarachi"
	TimeZoneEnumEtcGMTMinus5         TimeZoneEnum = "EtcGMTMinus5"
	TimeZoneEnumAsiaKolkata          TimeZoneEnum = "AsiaKolkata"
	TimeZoneEnumAsiaKatmandu         TimeZoneEnum = "AsiaKatmandu"
	TimeZoneEnumAsiaAlmaty           TimeZoneEnum = "AsiaAlmaty"
	TimeZoneEnumEtcGMTMinus6         TimeZoneEnum = "EtcGMTMinus6"
	TimeZoneEnumAsiaRangoon          TimeZoneEnum = "AsiaRangoon"
	TimeZoneEnumAsiaHovd             TimeZoneEnum = "AsiaHovd"
	TimeZoneEnumAsiaBangkok          TimeZoneEnum = "AsiaBangkok"
	TimeZoneEnumEtcGMTMinus7         TimeZoneEnum = "EtcGMTMinus7"
	TimeZoneEnumAsiaHongKong         TimeZoneEnum = "AsiaHongKong"
	TimeZoneEnumAsiaBrunei           TimeZoneEnum = "AsiaBrunei"
	TimeZoneEnumAsiaSingapore        TimeZoneEnum = "AsiaSingapore"
	TimeZoneEnumEtcGMTMinus8         TimeZoneEnum = "EtcGMTMinus8"
	TimeZoneEnumAsiaPyongyang        TimeZoneEnum = "AsiaPyongyang"
	TimeZoneEnumAustraliaEucla       TimeZoneEnum = "AustraliaEucla"
	TimeZoneEnumAsiaSeoul            TimeZoneEnum = "AsiaSeoul"
	TimeZoneEnumEtcGMTMinus9         TimeZoneEnum = "EtcGMTMinus9"
	TimeZoneEnumAustraliaDarwin      TimeZoneEnum = "AustraliaDarwin"
	TimeZoneEnumAustraliaAdelaide    TimeZoneEnum = "AustraliaAdelaide"
	TimeZoneEnumAustraliaSydney      TimeZoneEnum = "AustraliaSydney"
	TimeZoneEnumAustraliaBrisbane    TimeZoneEnum = "AustraliaBrisbane"
	TimeZoneEnumAsiaMagadan          TimeZoneEnum = "AsiaMagadan"
	TimeZoneEnumEtcGMTMinus10        TimeZoneEnum = "EtcGMTMinus10"
	TimeZoneEnumAustraliaLordHowe    TimeZoneEnum = "AustraliaLordHowe"
	TimeZoneEnumEtcGMTMinus11        TimeZoneEnum = "EtcGMTMinus11"
	TimeZoneEnumAsiaKamchatka        TimeZoneEnum = "AsiaKamchatka"
	TimeZoneEnumPacificFiji          TimeZoneEnum = "PacificFiji"
	TimeZoneEnumAntarcticaSouthPole  TimeZoneEnum = "AntarcticaSouthPole"
	TimeZoneEnumEtcGMTMinus12        TimeZoneEnum = "EtcGMTMinus12"
	TimeZoneEnumPacificChatham       TimeZoneEnum = "PacificChatham"
	TimeZoneEnumPacificTongatapu     TimeZoneEnum = "PacificTongatapu"
	TimeZoneEnumPacificApia          TimeZoneEnum = "PacificApia"
	TimeZoneEnumEtcGMTMinus13        TimeZoneEnum = "EtcGMTMinus13"
	TimeZoneEnumPacificKiritimati    TimeZoneEnum = "PacificKiritimati"
	TimeZoneEnumEtcGMTMinus14        TimeZoneEnum = "EtcGMTMinus14"
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
	PolicyManagedByEnumVMwareVSphere PolicyManagedByEnum = "VMware_vSphere"
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
	ManagedByID string `json:"managed_by_id,omitempty"`

	// Localized message string corresponding to interval
	IntervalL10n string `json:"interval_l10n,omitempty"`

	// Localized message string corresponding to timezone
	TimezoneL10n string `json:"timezone_l10n,omitempty"`

	// Localized message array corresponding to days_of_week
	DaysOfWeekL10n []string `json:"days_of_week_l10n,omitempty"`

	// Localized message string corresponding to nas_access_type
	NASAccessTypeL10n string `json:"nas_access_type_l10n,omitempty"`

	ManagedNyL10n string `json:"managed_by_l10n,omitempty"`

	Policies []ProtectionPolicy `json:"policies,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (s *SnapshotRule) Fields() []string {
	return []string{
		"id", "name",
		"interval", "time_of_day", "timezone", "days_of_week", "desired_retention",
		"is_replica", "nas_access_type", "is_read_only",
		"managed_by", "managed_by_id",
		"interval_l10n", "timezone_l10n", "days_of_week_l10n", "nas_access_type_l10n", "managed_by_l10n", "policies",
	}
}
