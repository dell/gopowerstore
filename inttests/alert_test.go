/*
 *
 * Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package inttests

import (
	"context"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

func Test_GetAlerts(t *testing.T) {
	client := GetNewClient()

	// get alerts until there are no alerts left to get
	pageIndex := 0
	for {
		alerts, err := client.GetAlerts(context.Background(), gopowerstore.GetAlertsOpts{
			RequestPagination: gopowerstore.RequestPagination{
				PageSize:   1000,
				StartIndex: pageIndex,
			},
		})

		assert.NoError(t, err)
		assert.NotEqual(t, alerts, &gopowerstore.GetAlertsResponse{})

		if alerts.AlertsResponseMeta.Pagination.Next == 0 {
			break
		}
		pageIndex = alerts.AlertsResponseMeta.Pagination.Next
	}
}
