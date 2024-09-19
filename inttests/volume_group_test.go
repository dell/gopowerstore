/*
 *
 * Copyright © 2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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

	g "github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VGName string = "csi_test_vg"
)

type VolumeGroupTestSuite struct {
	suite.Suite

	client  g.Client
	request g.VolumeGroup

	// Assurance that a volume group exists for the tests.
	assurance g.CreateResponse
}

func TestVolumeGroupSuite(t *testing.T) {
	suite.Run(t, new(VolumeGroupTestSuite))
}

func (s *VolumeGroupTestSuite) SetupSuite() {
	s.client = GetNewClient()

	var err error
	// Make sure a volume group exists on which we can run tests against.
	s.assurance, err = s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name:                   VGName,
		IsWriteOrderConsistent: false,
	})
	assert.NoError(s.T(), err)
}

func (s *VolumeGroupTestSuite) TearDownSuite() {
	// Delete the volume group we created for the test
	_, err := s.client.DeleteVolumeGroup(context.Background(), s.assurance.ID)
	assert.NoError(s.T(), err)
}

func (s *VolumeGroupTestSuite) SetupTest() {
}

func (s *VolumeGroupTestSuite) TearDownTest() {
}

// Returns true if one of the volume groups in vgs has an ID matching
// the ID provided by id and false if none of the volume groups have
// a matching ID.
func containsVolumeGroupID(vgs []g.VolumeGroup, id string) bool {
	for _, vg := range vgs {
		if vg.ID == id {
			return true
		}
	}
	return false
}

// Happy path test.
func (s *VolumeGroupTestSuite) TestGetVolumeGroups() {
	resp, err := s.client.GetVolumeGroups(context.Background())

	if assert.NoError(s.T(), err) {
		assert.True(s.T(), containsVolumeGroupID(resp, s.assurance.ID))
	}
}

type MetroVolumeGroupTestSuite struct {
	suite.Suite

	client        g.Client
	volumeGroupID string

	metro struct {
		config g.MetroConfig
	}
}

func TestMetroVolumeGroupSuite(t *testing.T) {
	suite.Run(t, new(MetroVolumeGroupTestSuite))
}

func (s *MetroVolumeGroupTestSuite) SetupSuite() {
	s.client = GetNewClient()

	// Get a remote system configured for metro replication
	remoteSystem := GetRemoteSystemForMetro(s.client, s.T())
	if remoteSystem.ID == "" {
		s.T().Skip("Could not get a remote system configured for metro. Skipping test suite...")
	}

	s.metro.config = g.MetroConfig{RemoteSystemID: remoteSystem.ID}

	// Create a volume group to run tests against
	resp, err := s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name:                   VGName,
		IsWriteOrderConsistent: false,
	})
	assert.NoError(s.T(), err)

	s.volumeGroupID = resp.ID
}

func (s *MetroVolumeGroupTestSuite) TearDownSuite() {
	// Delete the volume group when tests are finished
	_, err := s.client.DeleteVolumeGroup(context.Background(), s.volumeGroupID)
	if err != nil {
		s.T().Log("Unable to delete volume group. Please manually delete the volume group on the PowerStore array.")
	}
}

func (s *MetroVolumeGroupTestSuite) SetupTest() {
}

func (s *MetroVolumeGroupTestSuite) TearDownTest() {
}

func (s *MetroVolumeGroupTestSuite) TestConfigureMetroVolumeGroup() {
	// Should configure a metro volume group without errors.
	resp, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.volumeGroupID, &s.metro.config)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), resp)
}
