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

package gopowerstore

import (
	"context"

	"github.com/dell/gopowerstore/api"
)

const (
	smbShareURL = "smb_share"
)

// CreateSMBShare creates new SMB share
func (c *ClientIMPL) CreateSMBShare(ctx context.Context, createParams *SMBShareCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: smbShareURL,
			Body:     createParams,
		},
		&resp)
	return resp, WrapErr(err)
}

// ModifySMBShare modifies new SMB share
func (c *ClientIMPL) ModifySMBShare(ctx context.Context, id string, modifyParams *SMBShareCreate) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "PATCH",
			Endpoint: smbShareURL,
			ID:       id,
			Body:     modifyParams,
		},
		&resp)
	return resp, WrapErr(err)
}

// DeleteSMBShare deletes existing SMB share
func (c *ClientIMPL) DeleteSMBShare(ctx context.Context, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: smbShareURL,
			ID:       id,
		},
		&resp)
	return resp, WrapErr(err)
}

// GetSMBShare returns specific smb share by id
func (c *ClientIMPL) GetSMBShare(ctx context.Context, id string) (resp SMBShare, err error) {
	share := SMBShare{}
	qp := c.APIClient().QueryParamsWithFields(&share)

	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    smbShareURL,
			ID:          id,
			QueryParams: qp,
		},
		&resp)
	return resp, WrapErr(err)
}

// GetSMBShares returns SMB shares satisfying the filter
func (c *ClientIMPL) GetSMBShares(ctx context.Context, filter *string) (resp []SMBShare, err error) {
	var result []SMBShare
	err = c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []SMBShare
		share := SMBShare{}
		qp := c.APIClient().QueryParamsWithFields(&share)
		if filter != nil {
			qp.RawArg("", *filter)
		}
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    smbShareURL,
				QueryParams: qp,
			},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}
