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
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	hostGroupMockURL = APIMockURL + hostGroupURL
)

var hostGroupID = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
var hostGroupID2 = "3765da74-28a7-49db-a693-10cec1de91f8"

func TestClientIMPL_AttachVolumeToHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/attach", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(204, ""))
	attach := HostVolumeAttach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	attach.VolumeID = &id
	_, err := C.AttachVolumeToHostGroup(context.Background(), hostGroupID, &attach)
	assert.Nil(t, err)
}

func TestClientIMPL_DetachVolumeFromHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/detach", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(204, ""))
	detach := HostVolumeDetach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	detach.VolumeID = &id
	_, err := C.DetachVolumeFromHostGroup(context.Background(), hostGroupID, &detach)
	assert.Nil(t, err)
}

func TestClientIMPL_GetHostGroupByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", hostGroupMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, hostGroupID)
	setResponder(respData)
	hostGroup, err := C.GetHostGroupByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, hostGroupID, hostGroup.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetHostByName(context.Background(), "test")
	assert.NotNil(t, err)
}
