/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

const (
	TestVolumePrefix       = "test_vol_"
	DefaultVolSize   int64 = 1048576
	DefaultChunkSize int64 = 1048576
)

func createVol(t *testing.T) (string, string) {
	volName := TestVolumePrefix + randString(8)
	createParams := gopowerstore.VolumeCreate{}
	createParams.Name = &volName
	size := DefaultVolSize
	createParams.Size = &size
	createResp, err := C.CreateVolume(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, volName
}

func deleteVol(t *testing.T, id string) {
	_, err := C.DeleteVolume(context.Background(), nil, id)
	checkAPIErr(t, err)
}

func createSnap(volID string, t *testing.T, volName string) gopowerstore.CreateResponse {
	return createSnapWithSuffix(volID, t, volName, "snapshot")
}

func createSnapWithSuffix(volID string, t *testing.T, volName string, snapshotSuffix string) gopowerstore.CreateResponse {
	volume, err := C.GetVolume(context.Background(), volID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, volume.Name)
	assert.Equal(t, volName, volume.Name)
	snapName := volName + snapshotSuffix
	snapDesc := "just a description"
	snap, snapCreateErr := C.CreateSnapshot(context.Background(), &gopowerstore.SnapshotCreate{
		Name:        &snapName,
		Description: &snapDesc,
	}, volID)
	checkAPIErr(t, snapCreateErr)
	return snap
}

func TestModifyVolume(t *testing.T) {
	volID, _ := createVol(t)
	defer deleteVol(t, volID)

	_, err := C.ModifyVolume(context.Background(), &gopowerstore.VolumeModify{Size: DefaultVolSize * 2, Name: "rename"}, volID)
	checkAPIErr(t, err)
	gotVol, err := C.GetVolume(context.Background(), volID)
	checkAPIErr(t, err)
	assert.Equal(t, DefaultVolSize*2, gotVol.Size)
	assert.Equal(t, "rename", gotVol.Name)
}

func TestGetSnapshotsByVolumeID(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)

	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)

	snapList, err := C.GetSnapshotsByVolumeID(context.Background(), volID)
	checkAPIErr(t, err)

	assert.Equal(t, 1, len(snapList))
	assert.Equal(t, snap.ID, snapList[0].ID)
}

func TestGetSnapshot(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)

	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)

	got, err := C.GetSnapshot(context.Background(), snap.ID)
	checkAPIErr(t, err)

	assert.Equal(t, snap.ID, got.ID)
}

func TestGetSnapshots(t *testing.T) {
	_, err := C.GetSnapshots(context.Background())
	checkAPIErr(t, err)
}

func TestGetNonExistingSnapshot(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)

	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)
	_, err := C.DeleteSnapshot(context.Background(), nil, snap.ID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, snap.ID)

	got, err := C.GetSnapshot(context.Background(), snap.ID)
	assert.Error(t, err)
	assert.Empty(t, got)
}

func TestCreateSnapshot(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)
	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)
}

func TestDeleteSnapshot(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)
	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)
	_, err := C.DeleteSnapshot(context.Background(), nil, snap.ID)
	checkAPIErr(t, err)
}

func TestCreateVolumeFromSnapshot(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)
	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)

	name := "new_volume_from_snap" + randString(8)
	createParams := gopowerstore.VolumeClone{}
	createParams.Name = &name
	snapVol, err := C.CreateVolumeFromSnapshot(context.Background(), &createParams, snap.ID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, snapVol.ID)
	deleteVol(t, snapVol.ID)
}

func TestGetVolumes(t *testing.T) {
	_, err := C.GetVolumes(context.Background())
	checkAPIErr(t, err)
}

func TestGetVolume(t *testing.T) {
	volID, volName := createVol(t)
	volume, err := C.GetVolume(context.Background(), volID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, volume.Name)
	assert.Equal(t, volName, volume.Name)
	deleteVol(t, volID)
}

func TestGetVolumeByName(t *testing.T) {
	volID, volName := createVol(t)
	volume, err := C.GetVolumeByName(context.Background(), volName)
	checkAPIErr(t, err)
	assert.NotEmpty(t, volume.Name)
	assert.Equal(t, volName, volume.Name)
	deleteVol(t, volID)
}

func TestCreateDeleteVolume(t *testing.T) {
	volID, _ := createVol(t)
	deleteVol(t, volID)
}

func TestDeleteUnknownVol(t *testing.T) {
	volID := "f98de58e-9223-4fdc-86bd-d4ff268e20e1"
	_, err := C.DeleteVolume(context.Background(), nil, volID)
	if err != nil {
		apiError, ok := err.(gopowerstore.APIError)
		if !ok {
			t.Log("Unexpected API response")
			t.FailNow()
		}
		assert.True(t, apiError.NotFound())
	}
}

func TestGetVolumesWithTrace(t *testing.T) {
	ctx := C.SetTraceID(context.Background(),
		"126c9213-11d4-40b4-8da2-8cd70e277fe4")
	_, err := C.GetVolumes(ctx)
	checkAPIErr(t, err)
}

func TestVolumeAlreadyExist(t *testing.T) {
	volID, name := createVol(t)
	defer deleteVol(t, volID)
	createReq := gopowerstore.VolumeCreate{}
	createReq.Name = &name
	size := DefaultVolSize
	createReq.Size = &size
	_, err := C.CreateVolume(context.Background(), &createReq)
	assert.NotNil(t, err)
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.VolumeNameIsAlreadyUse())
}

func TestSnapshotAlreadyExist(t *testing.T) {
	volID, volName := createVol(t)
	defer deleteVol(t, volID)
	snap := createSnap(volID, t, volName)
	assert.NotEmpty(t, snap.ID)

	snapName := volName + "snapshot"
	snapDesc := "just a description"
	snap, err := C.CreateSnapshot(context.Background(), &gopowerstore.SnapshotCreate{
		Name:        &snapName,
		Description: &snapDesc,
	}, volID)
	assert.NotNil(t, err)
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.SnapshotNameIsAlreadyUse())
}

func TestGetInvalidVolume(t *testing.T) {
	_, err := C.GetVolume(context.Background(), "4961282c-c5c5-4234-935f-2742fed499d0")
	assert.NotNil(t, err)
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.NotFound())
}

func TestComputeDifferences(t *testing.T) {
	// Create volume
	volID, volName := createVol(t)
	defer deleteVol(t, volID)

	// Create snap of volume
	snap1 := createSnapWithSuffix(volID, t, volName, "snapshot1")
	assert.NotEmpty(t, snap1.ID)
	// Create another snap of volume
	snap2 := createSnapWithSuffix(volID, t, volName, "snapshot2")
	assert.NotEmpty(t, snap2.ID)
	// Run snap diff and validate there are no differences

	basesnapShotID := snap2.ID
	offset := int64(0)
	chunkSize := int64(DefaultChunkSize)
	length := int64(DefaultVolSize)
	snapdiffParams := gopowerstore.VolumeComputeDifferences{
		BaseSnapshotID: &basesnapShotID,
		ChunkSize:      &chunkSize,
		Length:         &length,
		Offset:         &offset,
	}
	defaultHeaders := C.GetCustomHTTPHeaders()
	if defaultHeaders == nil {
		defaultHeaders = make(http.Header)
	}
	customHeaders := defaultHeaders
	// for accessing internal REST-APIs
	customHeaders.Add("DELL-VISIBILITY", "internal")
	C.SetCustomHTTPHeaders(customHeaders)

	resp, err := C.ComputeDifferences(context.Background(), &snapdiffParams, snap1.ID)
	checkAPIErr(t, err)
	// AA== is equivalent to an empty bitmap
	assert.Equal(t, "AA==", *resp.ChunkBitmap)
	assert.Equal(t, int64(-1), *resp.NextOffset)
}
