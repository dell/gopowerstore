/*
 *
 * Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"
	"os"
	"strconv"

	"github.com/dell/gopowerstore/api"
)

// Env variables
const (
	APIURLEnv                 = "GOPOWERSTORE_APIURL"
	UsernameEnv               = "GOPOWERSTORE_USERNAME"
	PasswordEnv               = "GOPOWERSTORE_PASSWORD"
	InsecureEnv               = "GOPOWERSTORE_INSECURE"
	HTTPTimeoutEnv            = "GOPOWERSTORE_HTTP_TIMEOUT"
	DebugEnv                  = "GOPOWERSTORE_DEBUG"
	paginationDefaultPageSize = 1000
)

// Client defines gopowerstore client interface
type Client interface {
	APIClient() api.Client
	SetTraceID(ctx context.Context, value string) context.Context
	SetCustomHTTPHeaders(headers http.Header)
	GetVolume(ctx context.Context, id string) (Volume, error)
	GetVolumeByName(ctx context.Context, name string) (Volume, error)
	GetVolumes(ctx context.Context) ([]Volume, error)
	CreateVolume(ctx context.Context, createParams *VolumeCreate) (CreateResponse, error)
	DeleteVolume(ctx context.Context, deleteParams *VolumeDelete, id string) (EmptyResponse, error)
	GetHost(ctx context.Context, id string) (Host, error)
	GetHostByName(ctx context.Context, name string) (Host, error)
	GetHosts(ctx context.Context) ([]Host, error)
	CreateHost(ctx context.Context, createParams *HostCreate) (CreateResponse, error)
	DeleteHost(ctx context.Context, deleteParams *HostDelete, id string) (EmptyResponse, error)
	ModifyHost(ctx context.Context, modifyParams *HostModify, id string) (CreateResponse, error)
	GetHostVolumeMappings(ctx context.Context) (resp []HostVolumeMapping, err error)
	GetHostVolumeMapping(ctx context.Context, id string) (resp HostVolumeMapping, err error)
	GetHostVolumeMappingByVolumeID(ctx context.Context, volumeID string) (resp []HostVolumeMapping, err error)
	AttachVolumeToHost(ctx context.Context, hostID string, attachParams *HostVolumeAttach) (resp EmptyResponse, err error)
	DetachVolumeFromHost(ctx context.Context, hostID string, detachParams *HostVolumeDetach) (resp EmptyResponse, err error)
	GetStorageISCSITargetAddresses(ctx context.Context) ([]IPPoolAddress, error)
	GetApplianceListCMA(ctx context.Context) ([]Appliance, error)
	GetCapacity(ctx context.Context) (int64, error)
	GetFCPorts(ctx context.Context) (resp []FcPort, err error)
	GetFCPort(ctx context.Context, id string) (resp FcPort, err error)
	SetLogger(logger Logger)
	CreateSnapshot(ctx context.Context, createSnapParams *SnapshotCreate, id string) (resp CreateResponse, err error)
	DeleteSnapshot(ctx context.Context, deleteParams *VolumeDelete, id string) (EmptyResponse, error)
	GetSnapshotsByVolumeID(ctx context.Context, volID string) ([]Volume, error)
	GetSnapshots(ctx context.Context) ([]Volume, error)
	GetSnapshot(ctx context.Context, snapID string) (Volume, error)
	CreateVolumeFromSnapshot(ctx context.Context, createParams *VolumeClone, snapID string) (CreateResponse, error)
}

// ClientIMPL provides basic API client implementation
type ClientIMPL struct {
	API api.Client
}

// SetTraceID method allows to set tracing ID to context which will be used in log messages
func (c *ClientIMPL) SetTraceID(ctx context.Context, value string) context.Context {
	return c.API.SetTraceID(ctx, value)
}

// SetCustomHTTPHeaders method register headers which will be sent with every request
func (c *ClientIMPL) SetCustomHTTPHeaders(headers http.Header) {
	c.API.SetCustomHTTPHeaders(headers)
}

// Logger is interface required for gopowerstore custom logger
type Logger api.Logger

// SetLogger set logger which will be used by client
func (c *ClientIMPL) SetLogger(logger Logger) {
	c.API.SetLogger(api.Logger(logger))
}

// APIClient method returns powerstore API client may be useful for doing raw API requests
func (c *ClientIMPL) APIClient() api.Client {
	return c.API
}

// method allow to read paginated data from backend
func (c *ClientIMPL) readPaginatedData(f func(int) (api.RespMeta, error)) error {
	var err error
	var meta api.RespMeta
	meta, err = f(0)
	if err != nil {
		return err
	}
	if meta.Pagination.IsPaginate {
		for {
			nextOffset := meta.Pagination.Last + 1
			if nextOffset >= meta.Pagination.Total {
				break
			}
			meta, err = f(nextOffset)
			err = WrapErr(err)
			if err != nil {
				apiError, ok := err.(*APIError)
				if !ok {
					return err
				}
				if apiError.BadRange() {
					// could happen if some instances was deleted during pagination
					break
				}
			}
		}
	}
	return nil
}

// NewClient returns new PowerStore API client initialized from env vars
func NewClient() (Client, error) {
	options := NewClientOptions()
	insecure, err := strconv.ParseBool(os.Getenv(InsecureEnv))
	if err == nil {
		options.SetInsecure(insecure)
	}
	httpTimeout, err := strconv.ParseUint(os.Getenv(HTTPTimeoutEnv), 10, 64)

	if err == nil {
		options.SetDefaultTimeout(httpTimeout)
	}
	return NewClientWithArgs(
		os.Getenv(APIURLEnv),
		os.Getenv(UsernameEnv),
		os.Getenv(PasswordEnv),
		options)
}

// NewClientWithArgs returns new PowerStore API client initialized from args
func NewClientWithArgs(
	apiURL string,
	username, password string, options *ClientOptions) (Client, error) {
	client, err := api.New(apiURL, username, password,
		options.Insecure(), options.DefaultTimeout(), options.RequestIDKey())
	if err != nil {
		return nil, err
	}

	return &ClientIMPL{client}, nil
}
