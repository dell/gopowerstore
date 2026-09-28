/*
 *
 * Copyright © 2020-2026 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	nasMockURL                  = nasURL
	fsMockURL                   = fsURL
	nfsMockServerURL            = nfsServerURL
	apiSoftwareInstalledMockURL = apiSoftwareInstalledURL
	jobsMockURL                 = jobsURL
)

var (
	nasID  = "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	fsID   = "3765da74-28a7-49db-a693-10cec1de91f8"
	fsID2  = "3765da74-28a7-49db-a693-10cec1de91f9"
	jobID  = "3765da74-28a7-49db-a693-10cec1de91f0"
	jobID2 = "3765da74-28a7-49db-a693-10cec1de91f1"
)

func TestClientIMPL_GetNASByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", nasMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nasID)
	setResponder(respData)
	nas, err := C.GetNASByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, nasID, nas.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetNASByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_ListFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, fsID, fsID2)
	httpmock.RegisterResponder("GET", fsMockURL,
		httpmock.NewStringResponder(200, respData))
	fileSystems, err := C.ListFS(context.Background())
	assert.Nil(t, err)
	assert.Len(t, fileSystems, 2)
	assert.Equal(t, fsID, fileSystems[0].ID)
}

func TestClientIMPL_GetInProgressJobsByFsName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, jobID, jobID2)
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", jobsMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	setResponder(respData)
	job, err := C.GetInProgressJobsByFsName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, jobID, job[0].ID)
	httpmock.Reset()
	setResponder = func(respData string) {
		httpmock.RegisterResponder("GET", jobsMockURL,
			httpmock.NewStringResponder(404, respData))
	}
	_, err = C.GetInProgressJobsByFsName(context.Background(), "test")
	assert.NotNil(t, err)
}

func TestClientIMPL_GetFSByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", fsMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, fsID)
	setResponder(respData)
	fs, err := C.GetFSByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, fsID, fs.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetFSByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, fsID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", fsMockURL, fsID),
		httpmock.NewStringResponder(200, respData))
	fs, err := C.GetFS(context.Background(), fsID)
	assert.Nil(t, err)
	assert.Equal(t, fsID, fs.ID)
}

func TestClientIMPL_CreateFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, fsID)
	httpmock.RegisterResponder("POST", fsMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := FsCreate{
		Description: "some description",
		Name:        "new-fs",
		NASServerID: "5e8d8e8e-671b-336f-db4e-cee0fbdc981e",
		Size:        3221225472,
	}

	fs, err := C.CreateFS(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, fsID, fs.ID)
}

func TestClientIMPL_CreateFS_WithPerformancePolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, fsID)
	httpmock.RegisterResponder("POST", fsMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := FsCreate{
		Description:         "fs with performance policy",
		Name:                "new-fs-perf",
		NASServerID:         "5e8d8e8e-671b-336f-db4e-cee0fbdc981e",
		Size:                3221225472,
		PerformancePolicyID: "perf-policy-123",
	}

	fs, err := C.CreateFS(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, fsID, fs.ID)
}

func TestClientIMPL_CloneFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, fsID)
	cloneURL := fmt.Sprintf("%s/%s/clone", fsMockURL, fsID)
	httpmock.RegisterResponder("POST", cloneURL,
		httpmock.NewStringResponder(201, respData))
	description := "some description"
	name := "clone-fs"
	cloneReq := FsClone{
		Description: &description,
		Name:        &name,
	}

	resp, err := C.CloneFS(context.Background(), &cloneReq, fsID)
	assert.Nil(t, err)
	assert.Equal(t, fsID, resp.ID)
}

func TestClientIMPL_DeleteFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", fsMockURL, fsID),
		httpmock.NewStringResponder(204, ""))
	resp, err := C.DeleteFS(context.Background(), fsID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_ModifyFS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", fsMockURL, fsID),
		httpmock.NewStringResponder(204, ""))
	desc := "New Description"
	resp, err := C.ModifyFS(context.Background(), &FSModify{
		Size:        3221225472 * 2,
		Description: &desc,
	}, fsID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestFSModify_PointerSemantics_OmitsNilFields(t *testing.T) {
	desc := "test fs description"
	mod := FSModify{
		Description: &desc,
	}
	data, err := json.Marshal(mod)
	assert.Nil(t, err)
	jsonStr := string(data)
	assert.Contains(t, jsonStr, `"description"`)
	assert.NotContains(t, jsonStr, `"protection_policy_id"`)
	assert.NotContains(t, jsonStr, `"performance_policy_id"`)
}

func TestFSModify_PointerSemantics_IncludesSetFields(t *testing.T) {
	desc := "new fs desc"
	protPolicy := "prot-policy-1"
	mod := FSModify{
		Description:        &desc,
		ProtectionPolicyID: &protPolicy,
	}
	data, err := json.Marshal(mod)
	assert.Nil(t, err)
	jsonStr := string(data)
	assert.Contains(t, jsonStr, `"description":"new fs desc"`)
	assert.Contains(t, jsonStr, `"protection_policy_id":"prot-policy-1"`)
}

func TestClientIMPL_CreateNAS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nasID)
	httpmock.RegisterResponder("POST", nasMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := NASCreate{
		Description: "some description",
		Name:        "new-nas",
	}

	nas, err := C.CreateNAS(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, nasID, nas.ID)
}

func TestClientIMPL_GetNASServers(t *testing.T) {
	id := "6721f30c-405b-8749-439d-ee23cab1d298"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, id)
	httpmock.RegisterResponder("GET", nasMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetNASServers(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, id, resp[0].ID)
}

func TestClientIMPL_GetNAS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nasID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", nasMockURL, nasID),
		httpmock.NewStringResponder(200, respData))
	nas, err := C.GetNAS(context.Background(), nasID)
	assert.Nil(t, err)
	assert.Equal(t, nasID, nas.ID)
}

func TestClientIMPL_ModifyNASByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	nasName := "test-nas"
	nasID := "1234abcd-5678-efgh-9012-ijklmnopqrst"

	httpmock.RegisterResponder("GET",
		`=~^nas_server\?name=eq\.test-nas.*`,
		httpmock.NewStringResponder(200, `[{"id": "1234abcd-5678-efgh-9012-ijklmnopqrst"}]`))

	httpmock.RegisterResponder("PATCH",
		fmt.Sprintf("%s/%s", nasMockURL, nasID),
		httpmock.NewStringResponder(200, ""))

	modifyReq := NASModify{
		ProtectionPolicyID: "new-policy-id",
	}

	err := C.ModifyNASByName(context.Background(), &modifyReq, nasName)

	assert.NoError(t, err)
}

func TestClientIMPL_DeleteNAS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", nasMockURL, nasID),
		httpmock.NewStringResponder(204, ""))
	resp, err := C.DeleteNAS(context.Background(), nasID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetNfsServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nasID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", nfsMockServerURL, nfsID),
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetNfsServer(context.Background(), nfsID)
	assert.Nil(t, err)
	assert.Equal(t, nasID, resp.ID)
}

func TestClientIMPL_CreateFsSnapshot(t *testing.T) {
	id := "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, id)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/snapshot", fsMockURL, id),
		httpmock.NewStringResponder(200, respData))
	resp, err := C.CreateFsSnapshot(context.Background(), &SnapshotFSCreate{}, id)
	assert.Nil(t, err)
	assert.Equal(t, id, resp.ID)
}

func TestClientIMPL_GetFsSnapshot(t *testing.T) {
	id := "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, id)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", fsMockURL, id),
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetFsSnapshot(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, id, resp.ID)
}

func TestClientIMPL_GetFsSnapshots(t *testing.T) {
	id := "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, id)
	httpmock.RegisterResponder("GET", fsMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetFsSnapshots(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, id, resp[0].ID)
}

func TestClientIMPL_GetFsSnapshotsByVolumeID(t *testing.T) {
	id := "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, id)
	httpmock.RegisterResponder("GET", fsMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetFsSnapshotsByVolumeID(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, id, resp[0].ID)
}

func TestClientIMPL_CreateFsFromSnapshot(t *testing.T) {
	id := "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
	name := "test"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, id)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/clone", fsMockURL, id),
		httpmock.NewStringResponder(200, respData))
	resp, err := C.CreateFsFromSnapshot(context.Background(), &FsClone{Name: &name}, id)
	assert.Nil(t, err)
	assert.Equal(t, id, resp.ID)
}

func TestClientIMPL_GetFsByFilter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, fsID)
	httpmock.RegisterResponder("GET", fsMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetFsByFilter(context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, fsID, resp[0].ID)
}

func Test_GetNASFields(t *testing.T) {
	fields := GetNASFields(3.7)
	assert.NotEmpty(t, fields)
	assert.NotContains(t, fields, "health_details")
	fields = GetNASFields(3.5)
	assert.NotEmpty(t, fields)
	assert.NotContains(t, fields, "health_details")
}

func Test_GetFSFields(t *testing.T) {
	fields40 := GetFSFields(4.0)
	assert.NotContains(t, fields40, "performance_policy_id")

	fields41 := GetFSFields(4.1)
	assert.Contains(t, fields41, "performance_policy_id")
}

func TestClientIMPL_GetFS_PerformancePolicySelect(t *testing.T) {
	testCases := []struct {
		name          string
		buildVersion  string
		shouldContain bool
	}{
		{"4.0 excludes performance_policy_id", "4.0.0.0", false},
		{"4.1 includes performance_policy_id", "4.1.0.0", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", "=~^/?software_installed.*",
				httpmock.NewStringResponder(200, fmt.Sprintf(`[{"id": "1", "is_cluster": true, "build_version": "%s"}]`, tc.buildVersion)))

			var selectValue string
			httpmock.RegisterResponder("GET", "=~^/?file_system/.*",
				func(req *http.Request) (*http.Response, error) {
					selectValue = req.URL.Query().Get("select")
					return httpmock.NewStringResponse(200, fmt.Sprintf(`{"id": "%s"}`, fsID)), nil
				})

			_, err := C.GetFS(context.Background(), fsID)
			assert.Nil(t, err)
			if tc.shouldContain {
				assert.Contains(t, selectValue, "performance_policy_id")
			} else {
				assert.NotContains(t, selectValue, "performance_policy_id")
			}
		})
	}
}

func Test_NASServersErr(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nasID)
	httpmock.RegisterResponder("GET", apiSoftwareInstalledMockURL,
		httpmock.NewStringResponder(404, respData))
	_, err := C.GetNASServers(context.Background())
	assert.NotNil(t, err)
}

func Test_NASServerByIdErr(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nasID)
	httpmock.RegisterResponder("GET", apiSoftwareInstalledMockURL,
		httpmock.NewStringResponder(404, respData))
	_, err := C.GetNAS(context.Background(), nasID)
	assert.NotNil(t, err)
}

func Test_NASServerByNameErr(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nasID)
	httpmock.RegisterResponder("GET", apiSoftwareInstalledMockURL,
		httpmock.NewStringResponder(404, respData))
	_, err := C.GetNASByName(context.Background(), "test")
	assert.NotNil(t, err)
}
