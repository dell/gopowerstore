package gopowerstore

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	volumeGroupMockURL = APIMockURL + volumeGroupURL
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
		VolumeIds:          []string{volID2},
	}

	resp, err := C.CreateVolumeGroup(context.Background(), &createReq)
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
	respData := fmt.Sprintf(`{ "volume_group": [{"id": "%s", "name": "%s"}] }`, volID2, "volume-group")
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetVolumeGroupsByVolumeID(context.Background(), volID)
	assert.Nil(t, err)
	assert.NotNil(t, resp.VolumeGroup)
	assert.NotEqual(t, len(resp.VolumeGroup), 0)
	assert.Equal(t, volID2, resp.VolumeGroup[0].ID)
}

func TestClientIMPL_ModifyVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeGroupModify{
		ProtectionPolicyId: "new-id",
	}

	resp, err := C.ModifyVolumeGroup(context.Background(), &modifyParams, volID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_RemoveMembersFromVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/remove_members", volumeGroupMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	createReq := VolumeGroupRemoveMember{
		VolumeIds: []string{"id-1", "id-2"},
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
