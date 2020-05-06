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

package mock

import (
	"context"
	"github.com/dell/gopowerstore"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	Client gopowerstore.Client
}

func TestMockProvideClientInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewMockClient(ctrl)
	ts := testStruct{c}
	assert.NotNil(t, ts.Client)
}

func TestMockClient_GetVolumes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewMockClient(ctrl)

	volID := "39bb1b5f-5624-490d-9ece-18f7b28a904e"

	c.EXPECT().GetVolumes(gomock.Any()).
		Return([]gopowerstore.Volume{{ID: volID}}, nil)

	volList, err := c.GetVolumes(context.Background())
	assert.Nil(t, err)
	assert.True(t, len(volList) == 1)
}

func TestNewMockClient_GetVolume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewMockClient(ctrl)
	volID := "39bb1b5f-5624-490d-9ece-18f7b28a904e"
	c.EXPECT().GetVolume(gomock.Any(), gomock.Eq(volID)).
		Return(gopowerstore.Volume{ID: volID}, nil)

	apiError := gopowerstore.NewAPIError()
	apiError.StatusCode = 404

	c.EXPECT().GetVolume(gomock.Any(), gomock.Not(volID)).
		Return(gopowerstore.Volume{}, *apiError)

	vol, err := c.GetVolume(context.Background(), volID)
	assert.Nil(t, err)
	assert.Equal(t, volID, vol.ID)

	vol, err = c.GetVolume(context.Background(), "other_random_id")
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(gopowerstore.APIError).StatusCode)
	assert.Empty(t, vol.ID)

}
