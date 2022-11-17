/*
 *
 * Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	k8sClusterMockURL = APIMockURL + k8sClusterURL
)

var k8sClusterID = "5e8d7b7b-671b-336f-db4e-cee0fbdc981e"

func TestClientIMPL_RegisterK8sCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, k8sClusterID)
	httpmock.RegisterResponder("POST", k8sClusterMockURL,
		httpmock.NewStringResponder(201, respData))
	registerReq := K8sCluster{
		Name:      "test-cluster",
		IPAddress: "1.1.1.1",
		Port:      8080,
		Token:     "test-token",
	}

	cl, err := C.RegisterK8sCluster(context.Background(), &registerReq)
	assert.Nil(t, err)
	assert.Equal(t, k8sClusterID, cl.ID)
}
