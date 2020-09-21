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

import (
	"context"
	"errors"
)

const metricsURL = "metrics"

// GetCapacity return capacity of first appliance
func (c *ClientIMPL) GetCapacity(ctx context.Context) (int64, error) {
	var resp []ApplianceMetrics
	client := c.APIClient()
	qp := client.QueryParams().Select("physical_total", "physical_used")
	_, err := client.Query(
		ctx,
		RequestConfig{
			Method:      "POST",
			Endpoint:    metricsURL,
			Action:      "generate",
			QueryParams: qp,
			Body:        &MetricsRequest{
				Entity:   "space_metrics_by_appliance",
				EntityID: "A1",
				Interval: "Five_Mins",
			},
		},
		&resp)
	err = WrapErr(err)
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, errors.New("can't get appliance list")
	}
	freeSpace := resp[0].PhysicalTotal - resp[0].PhysicalUsed
	if freeSpace < 0 {
		return 0, nil
	}
	return freeSpace, nil
}
