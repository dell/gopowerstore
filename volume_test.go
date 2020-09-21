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
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	volumeMockURL = APIMockURL + volumeURL
)

var volID = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
var volID2 = "3765da74-28a7-49db-a693-10cec1de91f8"

func TestClientIMPL_GetVolumes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volID, volID2)
	httpmock.RegisterResponder("GET", volumeMockURL,
		httpmock.NewStringResponder(200, respData))
	vols, err := C.GetVolumes(context.Background())
	assert.Nil(t, err)
	assert.Len(t, vols, 2)
	assert.Equal(t, volID, vols[0].ID)
}

func TestClientIMPL_GetVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(200, respData))
	vol, err := C.GetVolume(context.Background(), volID)
	assert.Nil(t, err)
	assert.Equal(t, volID, vol.ID)
}

func TestClientIMPL_GetVolumeByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", volumeMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	vol, err := C.GetVolumeByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, vol.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetVolumeByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.VolumeIsNotExist())
}

func TestClientIMPL_GetSnapshotsByVolumeID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{
		"description":"",
		"id":"%s",
		"name":"rpo.VOLUME_esa51_volume_test1.2019-12-06T12:35:21Z 183173616",
		"size":10737418240,
		"state":"Ready",
		"storage_type":"Block",
		"type":"Snapshot",
		"wwn":null,
		"protection_data":{
			"family_id": "52ebb13c-16a0-4466-9319-a182b58b1c39",
			"parent_id": "52ebb13c-16a0-4466-9319-a182b58b1c39",
			"source_id": "%s", 
			"creator_type": "System", 
			"copy_signature": "e79bbc7f-da33-4f24-8eaa-0e3ef7533153",
			"source_timestamp": "2019-12-06T12:35:21.873907+00:00",
			"creator_type_l10n": "System",
			"is_app_consistent": null,
			"created_by_rule_id": null,
			"created_by_rule_name": null,
			"expiration_timestamp": null
		}
	}]`, volID2, volID)
	httpmock.RegisterResponder("GET", volumeMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetSnapshotsByVolumeID(context.Background(), volID2)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, volID2, resp[0].ID)
}

func TestClientIMPL_CreateVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", volumeMockURL,
		httpmock.NewStringResponder(201, respData))
	name := "test_vol"
	size := int64(11111111)
	createReq := VolumeCreate{}
	createReq.Name = &name
	createReq.Size = &size

	resp, err := C.CreateVolume(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_CreateSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID2)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/snapshot", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))
	name := "test_vol"
	desc := "desc"
	createReq := SnapshotCreate{}
	createReq.Name = &name
	createReq.Description = &desc

	resp, err := C.CreateSnapshot(context.Background(), &createReq, volID)
	assert.Nil(t, err)
	assert.Equal(t, volID2, resp.ID)
}

func TestClientIMPL_CreateVolumeFromSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID2)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/clone", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	name := "new_volume_from_snap"
	createParams := VolumeClone{}
	createParams.Name = &name
	resp, err := C.CreateVolumeFromSnapshot(context.Background(), &createParams, volID)
	assert.Nil(t, err)
	assert.Equal(t, volID2, resp.ID)
}

func TestClientIMPL_ModifyVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeModify{
		Name: "newname",
		Size: 8192 * 99,
	}

	resp, err := C.ModifyVolume(context.Background(), &modifyParams, volID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_DeleteSnapshot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(204, ""))
	force := true
	deleteReq := VolumeDelete{}
	deleteReq.ForceInternal = &force
	resp, err := C.DeleteSnapshot(context.Background(), &deleteReq, volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_DeleteVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(204, ""))
	force := true
	deleteReq := VolumeDelete{}
	deleteReq.ForceInternal = &force
	resp, err := C.DeleteVolume(context.Background(), &deleteReq, volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}
