/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	nasMockURL = APIMockURL + nasURL
	fsMockURL  = APIMockURL + fsURL
)

var nasID = "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
var fsID = "3765da74-28a7-49db-a693-10cec1de91f8"

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
	resp, err := C.ModifyFS(context.Background(), &FSModify{
		Size:        3221225472 * 2,
		Description: "New Description",
	}, fsID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
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

func TestClientIMPL_DeleteNAS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", nasMockURL, nasID),
		httpmock.NewStringResponder(204, ""))
	resp, err := C.DeleteNAS(context.Background(), nasID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}
