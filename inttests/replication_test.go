package inttests

import (
	"context"
	"fmt"
	"github.com/dell/gopowerstore"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type ReplicationTestSuite struct {
	suite.Suite
	remoteSystem    string
	remoteSystemMIP string
	remoteClient    gopowerstore.Client
	pp              gopowerstore.CreateResponse
	vg              gopowerstore.CreateResponse
	rr              gopowerstore.CreateResponse
	vol             gopowerstore.CreateResponse
}

func (suite *ReplicationTestSuite) SetupSuite() {
	suite.remoteSystem, suite.remoteSystemMIP = getRemoteSystem(suite.T())
	err := godotenv.Load("GOPOWERSTORE_TEST.env")
	if err != nil {
		return
	}
	user := os.Getenv("GOPOWERSTORE_USERNAME")
	pass := os.Getenv("GOPOWERSTORE_PASSWORD")

	clientOptions := &gopowerstore.ClientOptions{}
	clientOptions.SetInsecure(true)
	client, err := gopowerstore.NewClientWithArgs("https://"+suite.remoteSystemMIP+"/api/rest", user, pass, clientOptions)
	if err != nil {
		return
	}
	suite.remoteClient = client
}

func (suite *ReplicationTestSuite) TearDownSuite() {
	C.ModifyVolumeGroup(context.Background(), &gopowerstore.VolumeGroupModify{ProtectionPolicyId: ""}, suite.vg.ID)
	C.RemoveMembersFromVolumeGroup(context.Background(), &gopowerstore.VolumeGroupRemoveMember{VolumeIds: []string{suite.vol.ID}}, suite.vg.ID)
	C.ModifyVolume(context.Background(), &gopowerstore.VolumeModify{ProtectionPolicyID: ""}, suite.vol.ID)
	C.DeleteProtectionPolicy(context.Background(), suite.pp.ID)
	C.DeleteReplicationRule(context.Background(), suite.rr.ID)
	C.DeleteVolumeGroup(context.Background(), suite.vg.ID)
	vgid, err := suite.remoteClient.GetVolumeGroupByName(context.Background(), "intcsi-vgtst")
	if err != nil {
		logrus.Info(err)
	}
	suite.remoteClient.DeleteVolumeGroup(context.Background(), vgid.ID)
	C.DeleteVolume(context.Background(), nil, suite.vol.ID)
}
func getRemoteSystem(t *testing.T) (string, string) {
	resp, err := C.GetAllRemoteSystems(context.Background())
	skipTestOnError(t, err)
	if len(resp) == 0 {
		t.Skip("Skipping test as there are no remote systems configured on array.")
	}
	return resp[0].ID, resp[0].ManagementAddress
}

func (suite *ReplicationTestSuite) TestReplication() {
	t := suite.T()
	remoteSystem := suite.remoteSystem
	rs, err := C.GetRemoteSystem(context.Background(), remoteSystem)
	assert.NoError(t, err)
	assert.Equal(t, rs.ID, remoteSystem)

	suite.rr, err = C.CreateReplicationRule(context.Background(), &gopowerstore.ReplicationRuleCreate{
		Name:           "intcsi-ruletst",
		Rpo:            gopowerstore.RpoFifteenMinutes,
		RemoteSystemID: rs.ID,
	})
	assert.NoError(t, err)

	suite.pp, err = C.CreateProtectionPolicy(context.Background(), &gopowerstore.ProtectionPolicyCreate{
		Name:               "intcsi-pptst",
		ReplicationRuleIds: []string{suite.rr.ID},
	})
	assert.NoError(t, err)
	suite.vg, err = C.CreateVolumeGroup(context.Background(), &gopowerstore.VolumeGroupCreate{
		Name:               "intcsi-vgtst",
		ProtectionPolicyID: suite.pp.ID,
	})
	assert.NoError(t, err)

	volName := "intcsi-voltst"
	tpe := gopowerstore.StorageTypeEnumBlock
	size := int64(1048576)
	suite.vol, err = C.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{
		Name:          &volName,
		Size:          &size,
		StorageType:   &tpe,
		VolumeGroupID: suite.vg.ID,
	})
	assert.NoError(t, err)
	volId := suite.vol.ID
	_, err = C.GetVolumeGroupsByVolumeID(context.Background(), volId)
	assert.NoError(t, err)

	for tout := 0; tout < 30; tout += 1 {
		_, err = C.GetReplicationSessionByLocalResourceID(context.Background(), suite.vg.ID)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
		fmt.Printf("Retrying.")
	}

	assert.NoError(t, err)

}

func TestGetCluster(t *testing.T) {
	resp, err := C.GetCluster(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp.ID, "0")
}
func TestReplicationSuite(t *testing.T) {
	suite.Run(t, new(ReplicationTestSuite))
}
