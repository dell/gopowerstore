/*
 *
 * Copyright © 2020-2025 Dell Inc. or its subsidiaries. All Rights Reserved.
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

 type NASHealthStatus struct {
    ID            string        `json:"id"`
    Name          string        `json:"name"`
    HealthDetails HealthDetails `json:"health_details"`
}

type HealthDetails struct {
    State              string `json:"state"`
    InfoAckedCount     int    `json:"info_acked_count"`
    MajorAckedCount    int    `json:"major_acked_count"`
    MinorAckedCount    int    `json:"minor_acked_count"`
    InfoUnackedCount   int    `json:"info_unacked_count"`
    MajorUnackedCount  int    `json:"major_unacked_count"`
    MinorUnackedCount  int    `json:"minor_unacked_count"`
    CriticalAckedCount int    `json:"critical_acked_count"`
    CriticalUnackedCount int  `json:"critical_unacked_count"`
}