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

import "context"

const (
	limitURL      = "limit"
	maxVolumeSize = "Max_Volume_Size"
)

func (c *ClientIMPL) callGetLimit(ctx context.Context) (resp map[string]int64, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "GET",
			Endpoint: limitURL,
		},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) GetMaxVolumeSize(ctx context.Context) (int64, error) {
	resp, err := c.callGetLimit(ctx)
	limit, ok := resp[maxVolumeSize]
	if !ok {
		limit = -1
	}
	return limit, err
}
