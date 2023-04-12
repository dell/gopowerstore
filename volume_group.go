/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/dell/gopowerstore/api"
	"github.com/sirupsen/logrus"
)

const (
	volumeGroupURL = "volume_group"
	snapshotURL    = "/snapshot"
)

func getVolumeGroupDefaultQueryParams(c Client) api.QueryParamsEncoder {
	vol := VolumeGroup{}
	return c.APIClient().QueryParamsWithFields(&vol)
}

// GetVolumeGroup query and return specific volume group by id
func (c *ClientIMPL) GetVolumeGroup(ctx context.Context, id string) (resp VolumeGroup, err error) {
	qp := getVolumeGroupDefaultQueryParams(c)
	qp.Select("volume.volume_group_membership(id,name,protection_policy_id,state,protection_data)")
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeGroupURL,
			ID:          id,
			QueryParams: qp},
		&resp)
	return resp, WrapErr(err)
}

// GetVolumeGroupByName query and return specific volume group by name
func (c *ClientIMPL) GetVolumeGroupByName(ctx context.Context, name string) (resp VolumeGroup, err error) {
	var groups []VolumeGroup
	qp := getVolumeGroupDefaultQueryParams(c)
	qp.RawArg("name", fmt.Sprintf("eq.%s", name))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeGroupURL,
			QueryParams: qp},
		&groups)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(groups) != 1 {
		return resp, NewNotFoundError()
	}
	return groups[0], err
}

// GetVolumeGroups returns a list of volume groups
func (c *ClientIMPL) GetVolumeGroups(ctx context.Context) ([]VolumeGroup, error) {
	var result []VolumeGroup
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []VolumeGroup
		volume_group := VolumeGroup{}
		qp := c.APIClient().QueryParamsWithFields(&volume_group)
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    volumeGroupURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// CreateVolumeGroup creates new volume group
func (c *ClientIMPL) CreateVolumeGroup(ctx context.Context,
	createParams *VolumeGroupCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeGroupURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) GetVolumeGroupsByVolumeID(ctx context.Context, id string) (resp VolumeGroups, err error) {
	qp := c.API.QueryParams()
	qp.Select("volume_group.volume_group_membership(id,name,protection_policy_id)")
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeURL,
			ID:          id,
			QueryParams: qp},
		&resp)
	logrus.Info(resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) UpdateVolumeGroupProtectionPolicy(ctx context.Context, id string, params *VolumeGroupChangePolicy) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "PATCH",
			Endpoint: volumeGroupURL,
			ID:       id,
			Body:     params,
		},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) RemoveMembersFromVolumeGroup(ctx context.Context,
	params *VolumeGroupMembers, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeGroupURL,
			ID:       id,
			Body:     params,
			Action:   "remove_members",
		},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) AddMembersToVolumeGroup(ctx context.Context,
	params *VolumeGroupMembers, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeGroupURL,
			ID:       id,
			Body:     params,
			Action:   "add_members",
		},
		&resp)
	return resp, WrapErr(err)
}

// DeleteVolumeGroup deletes existing VG
func (c *ClientIMPL) DeleteVolumeGroup(ctx context.Context, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			ID:       id,
			Endpoint: volumeGroupURL,
		},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) ModifyVolumeGroup(ctx context.Context,
	modifyParams *VolumeGroupModify, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "PATCH",
			Endpoint: volumeGroupURL,
			ID:       id,
			Body:     modifyParams},
		&resp)
	return resp, WrapErr(err)
}

// CreateVolumeGroupSnapshot Creates a new volume group snapshot from the existing volume group
func (c *ClientIMPL) CreateVolumeGroupSnapshot(ctx context.Context, volumeGroupID string,
	createParams *VolumeGroupSnapshotCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeGroupURL + "/" + volumeGroupID + snapshotURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

// GetVolumeGroupSnapshot query and return specific snapshot by id
func (c *ClientIMPL) GetVolumeGroupSnapshot(ctx context.Context, snapID string) (resVol VolumeGroup, err error) {
	qp := getVolumeDefaultQueryParams(c)
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeGroupURL,
			ID:          snapID,
			QueryParams: qp},
		&resVol)
	return resVol, WrapErr(err)
}
