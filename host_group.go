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

import (
	"context"
)

const (
	hostGroupURL = "host_group"
)

// AttachVolumeToHost attaches volume to hostGroup
func (c *ClientIMPL) AttachVolumeToHostGroup(
	ctx context.Context,
	hostGroupID string,
	attachParams *HostVolumeAttach) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: hostGroupURL,
			ID:       hostGroupID,
			Action:   "attach",
			Body:     attachParams},
		&resp)
	return resp, WrapErr(err)
}

// DetachVolumeFromHost detaches volume to hostGroup
func (c *ClientIMPL) DetachVolumeFromHostGroup(
	ctx context.Context,
	hostGroupID string,
	detachParams *HostVolumeDetach) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: hostGroupURL,
			ID:       hostGroupID,
			Action:   "detach",
			Body:     detachParams},
		&resp)
	return resp, WrapErr(err)
}
