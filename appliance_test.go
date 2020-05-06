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

const applianceMockURL = APIMockURL + applianceListCmaViewURL

func TestClientIMPL_GetApplianceListCMA(t *testing.T) {
	id := "A1"
	applianceType := ApplianceTypeEnumPowerStore
	applianceMode := ApplianceModeEnumUnified

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s", "name": "appliance-1", "ip_address": "127.0.0.1",
"appliance_type": "%s", "mode": "%s", "last_physical_total_space": 12077448036352, "last_physical_used_space": 1905262588}]`,
		id, applianceType, applianceMode)
	httpmock.RegisterResponder("GET", applianceMockURL,
		httpmock.NewStringResponder(200, respData))

	appliances, err := C.GetApplianceListCMA(context.Background())
	assert.Nil(t, err)
	assert.Len(t, appliances, 1)
	assert.Equal(t, id, appliances[0].ID)
	assert.Equal(t, applianceType, appliances[0].Type)
	assert.Equal(t, applianceMode, appliances[0].Mode)
}

func TestClientIMPL_GetApplianceListCMA_NotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[]`)
	httpmock.RegisterResponder("GET", applianceMockURL,
		httpmock.NewStringResponder(200, respData))

	_, err := C.GetApplianceListCMA(context.Background())
	assert.NotNil(t, err)
}

func TestClientIMPL_GetCapacity(t *testing.T) {
	totalSpace := 12077448036352
	usedSpace := 1905262588
	freeSpace := int64(totalSpace - usedSpace)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"last_physical_total_space": %d, "last_physical_used_space": %d}]`,
		totalSpace, usedSpace)
	httpmock.RegisterResponder("GET", applianceMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, freeSpace, resp)
}

func TestClientIMPL_GetCapacity_Zero(t *testing.T) {
	totalSpace := 1905262588
	usedSpace := 12077448036352
	var freeSpace int64

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"last_physical_total_space": %d, "last_physical_used_space": %d}]`,
		totalSpace, usedSpace)
	httpmock.RegisterResponder("GET", applianceMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, freeSpace, resp)
}

func TestClientIMPL_GetCapacity_NotFoundAppliance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[]`)
	httpmock.RegisterResponder("GET", APIMockURL+`/appliance_list_cma_view`,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), resp)
}
