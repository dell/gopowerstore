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

const applianceListCmaViewURL = "appliance_list_cma_view"

// GetApplianceListCMA return a list of Appliance
func (c *ClientIMPL) GetApplianceListCMA(ctx context.Context) (resp []Appliance, err error) {
	client := c.APIClient()
	var appliance Appliance
	qp := client.QueryParams().Select(appliance.Fields()...)
	_, err = client.Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    applianceListCmaViewURL,
			QueryParams: qp},
		&resp)
	err = WrapErr(err)
	if err != nil {
		return
	}
	if len(resp) == 0 {
		return resp, errors.New("can't get appliance list")
	}
	return
}

// GetCapacity return capacity of first appliance
func (c *ClientIMPL) GetCapacity(ctx context.Context) (int64, error) {
	var resp []Appliance
	client := c.APIClient()
	qp := client.QueryParams().Select("last_physical_total_space", "last_physical_used_space")
	_, err := client.Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    applianceListCmaViewURL,
			QueryParams: qp},
		&resp)
	err = WrapErr(err)
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, errors.New("can't get appliance list")
	}
	freeSpace := resp[0].LastPhysicalTotalSpace - resp[0].LastPhysicalUsedSpace
	if freeSpace < 0 {
		return 0, nil
	}
	return freeSpace, nil
}
