/*
 *
 * Copyright © 2021-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	volumeGroupMockURL         = APIMockURL + volumeGroupURL
	volumeGroupSnapshotMockURL = APIMockURL + volumeGroupURL + "/test-id" + snapshotURL
)

var (
	volGroupSnapID  = "1966782b-60c9-40e2-a1ee-9b2b8f6b98e7"
	volGroupSnapID2 = "34380c29-2203-4490-aeb7-2853b9a85075"
)

func TestClientIMPL_CreateVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", volumeGroupMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := VolumeGroupCreate{
		Name:               "vg-test",
		Description:        "vg-test",
		ProtectionPolicyID: volID,
		VolumeIDs:          []string{volID2},
	}

	resp, err := C.CreateVolumeGroup(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_CreateVolumeGroupSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", volumeGroupSnapshotMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := VolumeGroupSnapshotCreate{
		Name:        "vgs-test",
		Description: "vgs-test",
	}

	resp, err := C.CreateVolumeGroupSnapshot(context.Background(), "test-id", &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_DeleteVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteVolumeGroup(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(200, respData))
	volumeGroup, err := C.GetVolumeGroup(context.Background(), volID)
	assert.Nil(t, err)
	assert.Equal(t, volID, volumeGroup.ID)
}

func TestClientIMPL_GetVolumeGroupByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", volumeGroupMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	volumeGroup, err := C.GetVolumeGroupByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, volumeGroup.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetVolumeGroupByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetVolumeGroupsByVolumeID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{ "volume_groups": [{"id": "%s", "name": "%s"}] }`, volID2, "volume-group")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetVolumeGroupsByVolumeID(context.Background(), volID)
	assert.Nil(t, err)
	assert.NotNil(t, resp.VolumeGroup)
	assert.NotEqual(t, len(resp.VolumeGroup), 0)
	assert.Equal(t, volID2, resp.VolumeGroup[0].ID)
}

func TestClientIMPL_GetVolumeGroups(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volID, volID2)
	httpmock.RegisterResponder("GET", volumeGroupMockURL,
		httpmock.NewStringResponder(200, respData))
	volumeGroups, err := C.GetVolumeGroups(context.Background())
	assert.Nil(t, err)
	assert.Len(t, volumeGroups, 2)
	assert.Equal(t, volID, volumeGroups[0].ID)
}

func TestClientIMPL_GetVolumeGroupSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volGroupSnapID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeGroupMockURL, volGroupSnapID),
		httpmock.NewStringResponder(200, respData))
	snapshot, err := C.GetVolumeGroupSnapshot(context.Background(), volGroupSnapID)
	assert.Nil(t, err)
	assert.Equal(t, volGroupSnapID, snapshot.ID)
}

func TestClientIMPL_GetVolumeGroupSnapshots(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volGroupSnapID, volGroupSnapID2)
	httpmock.RegisterResponder("GET", volumeGroupMockURL,
		httpmock.NewStringResponder(200, respData))
	snapshots, err := C.GetVolumeGroupSnapshots(context.Background())
	assert.Nil(t, err)
	assert.Len(t, snapshots, 2)
	assert.Equal(t, volGroupSnapID, snapshots[0].ID)
}

func TestClientIMPL_GetVolumeGroupSnapshotByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", volumeGroupMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volGroupSnapID)
	setResponder(respData)
	snapshot, err := C.GetVolumeGroupSnapshotByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volGroupSnapID, snapshot.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetVolumeGroupSnapshotByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_ModifyVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeGroupModify{
		ProtectionPolicyID: "new-id",
	}

	resp, err := C.ModifyVolumeGroup(context.Background(), &modifyParams, volID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_ModifyVolumeGroupSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprint(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeGroupSnapshotModify{
		Description:            "test description",
		Name:                   "test name",
		IsWriteOrderConsistent: false,
	}

	resp, err := C.ModifyVolumeGroupSnapshot(context.Background(), &modifyParams, volID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_RemoveMembersFromVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/remove_members", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	createReq := VolumeGroupMembers{
		VolumeIDs: []string{"id-1", "id-2"},
	}

	resp, err := C.RemoveMembersFromVolumeGroup(context.Background(), &createReq, volID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_UpdateVolumeGroupProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeGroupChangePolicy{
		ProtectionPolicyID: "id-1",
	}

	resp, err := C.UpdateVolumeGroupProtectionPolicy(context.Background(), volID, &modifyParams)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}
