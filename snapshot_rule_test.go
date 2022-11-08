package gopowerstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	snapshotRuleMockURL = APIMockURL + snapshotRuleURL
)

var (
	snapshotRuleID  = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
	snapshotRuleID2 = "3765da74-28a7-49db-a693-10cec1de91f8"
)

func TestClientIMPL_GetSnapshotRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, snapshotRuleID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", snapshotRuleMockURL, snapshotRuleID),
		httpmock.NewStringResponder(200, respData))
	snapshotRule, err := C.GetSnapshotRule(context.Background(), snapshotRuleID)
	assert.Nil(t, err)
	assert.Equal(t, snapshotRuleID, snapshotRule.ID)
}

func TestClientIMPL_GetSnapshotRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, snapshotRuleID, snapshotRuleID2)
	httpmock.RegisterResponder("GET", snapshotRuleMockURL,
		httpmock.NewStringResponder(200, respData))
	snapshotRules, err := C.GetSnapshotRules(context.Background())
	assert.Nil(t, err)
	assert.Len(t, snapshotRules, 2)
	assert.Equal(t, snapshotRuleID, snapshotRules[0].ID)
}

func TestClientIMPL_CreateSnapshotRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, snapshotRuleID)
	httpmock.RegisterResponder("POST", snapshotRuleMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := SnapshotRuleCreate{
		Name:             "test_snapshotrule",
		DesiredRetention: 8,
		Interval:         SnapshotRuleIntervalEnumFour_Hours,
	}

	resp, err := C.CreateSnapshotRule(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, snapshotRuleID, resp.ID)
}

func TestClientIMPL_ModifySnapshotRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", snapshotRuleMockURL, snapshotRuleID),
		httpmock.NewStringResponder(201, ""))

	modifyParams := SnapshotRuleCreate{
		TimeZone: TimeZoneEnumUS__Pacific,
	}

	resp, err := C.ModifySnapshotRule(context.Background(), &modifyParams, snapshotRuleID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_DeleteSnapshotRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", snapshotRuleMockURL, snapshotRuleID),
		httpmock.NewStringResponder(204, ""))
	deleteReq := SnapshotRuleDelete{
		DeleteSnaps: true,
	}
	resp, err := C.DeleteSnapshotRule(context.Background(), &deleteReq, snapshotRuleID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}
