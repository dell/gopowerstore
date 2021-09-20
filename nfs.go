/*
 *
 * Copyright 2020 Dell EMC Corporation
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
	"fmt"
	"github.com/dell/gopowerstore/api"
)

const (
	nfsURL           = "nfs_export"
	nfsServerURL     = "nfs_server"
	fileInterfaceURL = "file_interface"
)

func getNFSExportDefaultQueryParams(c Client) api.QueryParamsEncoder {
	nfs := NFSExport{}
	return c.APIClient().QueryParamsWithFields(&nfs)
}

func getFileInterfaceDefaultQueryParams(c Client) api.QueryParamsEncoder {
	fi := FileInterface{}
	return c.APIClient().QueryParamsWithFields(&fi)
}

// GetNFSExportByName query and return specific NFS export by name
func (c *ClientIMPL) GetNFSExportByName(ctx context.Context, name string) (resp NFSExport, err error) {
	var nfsList []NFSExport
	qp := getNFSExportDefaultQueryParams(c)
	qp.RawArg("name", fmt.Sprintf("eq.%s", name))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    nfsURL,
			QueryParams: qp},
		&nfsList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(nfsList) != 1 {
		return resp, NewNotFoundError()
	}
	return nfsList[0], err
}

// GetNFSExportByName query and return specific NFS export by its filesystems name
func (c *ClientIMPL) GetNFSExportByFileSystemID(ctx context.Context, fsID string) (resp NFSExport, err error) {
	var nfsList []NFSExport
	qp := getNFSExportDefaultQueryParams(c)
	qp.RawArg("file_system_id", fmt.Sprintf("eq.%s", fsID))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    nfsURL,
			QueryParams: qp},
		&nfsList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(nfsList) != 1 {
		return resp, NewNotFoundError()
	}
	return nfsList[0], err
}

// CreateNFSExport creates new NFS export on storage array
func (c *ClientIMPL) CreateNFSExport(ctx context.Context, createParams *NFSExportCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: nfsURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

// DeleteNFSExport deletes existing NFS export from storage array
func (c *ClientIMPL) DeleteNFSExport(ctx context.Context, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: nfsURL,
			ID:       id},
		&resp)
	return resp, WrapErr(err)
}

// ModifyNFSExport patches existing NFS export, adding or removing new hosts
func (c *ClientIMPL) ModifyNFSExport(ctx context.Context,
	modifyParams *NFSExportModify, id string) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "PATCH",
			Endpoint: nfsURL,
			ID:       id,
			Body:     modifyParams},
		&resp)
	return resp, WrapErr(err)
}

// CreateNFSServer creates new NFS server on storage array
func (c *ClientIMPL) CreateNFSServer(ctx context.Context,
	createParams *NFSServerCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: nfsServerURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

// GetFileInterface returns FileInterface from storage array by id
func (c *ClientIMPL) GetFileInterface(ctx context.Context, id string) (resp FileInterface, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    fileInterfaceURL,
			ID:          id,
			QueryParams: getFileInterfaceDefaultQueryParams(c)},
		&resp)
	return resp, WrapErr(err)
}
