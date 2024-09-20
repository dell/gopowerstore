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
	"fmt"
	"net/http"
	"testing"

	g "github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VGPrefix string = "test_vg_"
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
		Name: VGPrefix + randString(8),
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

/*
/	////////////////////////////
/	/ METRO VOLUME GROUP TESTS /
/ 	////////////////////////////
*/
type MetroVolumeGroupTestSuite struct {
	suite.Suite

	client g.Client
	vg     g.VolumeGroup

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
}

func (s *MetroVolumeGroupTestSuite) TearDownSuite() {
}

func (s *MetroVolumeGroupTestSuite) SetupTest() {
	// create a unique vg name for each test run
	s.vg.Name = VGPrefix + randString(8)

	// Create a volume group to run tests against
	resp, err := s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name: s.vg.Name,
	})
	assert.NoError(s.T(), err)

	s.vg.ID = resp.ID

	// create a volume to add to the vg to make it a valid vg we can test with
	volID, volName := CreateVol(s.T())
	s.vg.Volumes = append(s.vg.Volumes, g.Volume{ID: volID, Name: volName})
}

func (s *MetroVolumeGroupTestSuite) TearDownTest() {
	// Delete all the volumes in the volume group
	err := s.deleteVolumesInVG()
	if err != nil {
		s.T().Logf("%s. Please delete from PowerStore when tests complete.", err.Error())
	}

	// Delete the volume group from the previous test.
	_, err = s.client.DeleteVolumeGroup(context.Background(), s.vg.ID)
	if err != nil {
		// 404 status means it was already deleted.
		// warn about other errors encountered while deleting
		if err.(g.APIError).StatusCode != http.StatusNotFound {
			s.T().Logf("Unable to delete test volume group %s. Please delete from PowerStore when tests complete. err: %s", s.vg.Name, err.Error())
		}
	}

	// Sanitize for next test.
	s.vg.Name = ""
	s.vg.ID = ""
}

func (s *MetroVolumeGroupTestSuite) deleteVolumesInVG() error {
	for _, vol := range s.vg.Volumes {
		_, err := s.client.DeleteVolume(context.Background(), nil, vol.ID)
		if err != nil {
			// 404 status means it was already deleted.
			// warn about other errors encountered while deleting
			if err.(g.APIError).StatusCode != http.StatusNotFound {
				return fmt.Errorf("Error encountered deleting vol: %s. Err: %s", vol.Name, err)
			}
		}
	}
	return nil
}

// Should configure a metro volume group without errors.
func (s *MetroVolumeGroupTestSuite) TestConfigureMetroVolumeGroup() {
	resp, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.ID, &s.metro.config)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), resp)
}

// Try to configure metro on a volume group without any volumes in it.
func (s *MetroVolumeGroupTestSuite) TestConfigMetroVGOnEmptyVG() {
	// delete all the volumes from the volume group
	err := s.deleteVolumesInVG()
	assert.NoError(s.T(), err)

	// Attempt to configure metro on an empty volume group
	_, err = s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.ID, &s.metro.config)

	assert.Equal(s.T(), http.StatusUnprocessableEntity, err.(g.APIError).StatusCode)
}

// Try to configure metro on a non-existent volume group.
func (s *MetroVolumeGroupTestSuite) TestMetroVGNonExistantVG() {
	// Delete that volume group, retaining the volume group ID.
	_, err := s.client.DeleteVolumeGroup(context.Background(), s.vg.ID)
	assert.NoError(s.T(), err)

	// Try to configure metro volume group using the deleted vg ID.
	_, err = s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.ID, &s.metro.config)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), http.StatusNotFound, err.(g.APIError).StatusCode)
}

// Execute ConfigureMetroVolume with a bad request body.
func (s *MetroVolumeGroupTestSuite) TestMetroVGBadRequest() {
	// Pass an emtpy configuration body with the request
	_, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.ID, nil)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), http.StatusBadRequest, err.(g.APIError).StatusCode)
}
