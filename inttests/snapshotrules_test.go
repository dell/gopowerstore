package inttests

import (
	"context"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

const TestSnapshotRulePrefix = "test_snapshotrule_"

func createSnapshotRule(t *testing.T) (string, string) {
	snapshotRuleName := TestSnapshotRulePrefix + randString(8)
	createParams := gopowerstore.SnapshotRuleCreate{
		Name:             snapshotRuleName,
		DesiredRetention: 8,
		Interval:         gopowerstore.SnapshotRuleIntervalEnumFour_Hours,
	}
	createResp, err := C.CreateSnapshotRule(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, snapshotRuleName
}

func deleteSnapshotRule(t *testing.T, id string) {
	_, err := C.DeleteSnapshotRule(
		context.Background(),
		&gopowerstore.SnapshotRuleDelete{
			DeleteSnaps: true,
		},
		id,
	)
	checkAPIErr(t, err)
}

func TestGetSnapshotRules(t *testing.T) {
	_, err := C.GetSnapshotRules(context.Background())
	checkAPIErr(t, err)
}

func TestGetSnapshotRule(t *testing.T) {
	snapshotRuleID, snapshotRuleName := createSnapshotRule(t)
	defer deleteSnapshotRule(t, snapshotRuleID)

	got, err := C.GetSnapshotRule(context.Background(), snapshotRuleID)
	checkAPIErr(t, err)

	assert.Equal(t, snapshotRuleID, got.ID)
	assert.Equal(t, snapshotRuleName, got.Name)
}

func TestModifySnapshotRule(t *testing.T) {
	snapshotRuleID, _ := createSnapshotRule(t)
	defer deleteSnapshotRule(t, snapshotRuleID)

	_, err := C.ModifySnapshotRule(context.Background(), &gopowerstore.SnapshotRuleCreate{DesiredRetention: 7}, snapshotRuleID)
	checkAPIErr(t, err)
	got, err := C.GetSnapshotRule(context.Background(), snapshotRuleID)
	checkAPIErr(t, err)
	assert.Equal(t, 7, got.DesiredRetention)
}
