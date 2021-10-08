package gopowerstore

import (
	"context"
	"fmt"
	"github.com/dell/gopowerstore/api"
	"github.com/sirupsen/logrus"
)

const (
	volumeGroupURL = "volume_group"
)

func getVolumeGroupDefaultQueryParams(c Client) api.QueryParamsEncoder {
	vol := VolumeGroup{}
	return c.APIClient().QueryParamsWithFields(&vol)
}

// GetVolumeGroup query and return specific volume group by id
func (c *ClientIMPL) GetVolumeGroup(ctx context.Context, id string) (resp VolumeGroup, err error) {
	qp := getVolumeGroupDefaultQueryParams(c)
	qp.Select("volume.volume_group_membership(id,name,protection_policy_id)")
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
	params *VolumeGroupRemoveMember, id string) (resp EmptyResponse, err error) {
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
