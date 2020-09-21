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

const metricsMockURL = APIMockURL + metricsURL


func TestClientIMPL_GetCapacity(t *testing.T) {
	totalSpace := 12077448036352
	usedSpace := 1905262588
	freeSpace := int64(totalSpace - usedSpace)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"physical_total": %d, "physical_used": %d}]`,
		totalSpace, usedSpace)
	httpmock.RegisterResponder("POST", metricsMockURL + "/generate",
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
	respData := fmt.Sprintf(`[{"physical_total": %d, "physical_used": %d}]`,
		totalSpace, usedSpace)
	httpmock.RegisterResponder("POST", metricsMockURL + "/generate",
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, freeSpace, resp)
}
