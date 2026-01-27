/*
 *
 * Copyright © 2026 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	ethPortMockURL = apiEthPortURL
)

var (
	ethPortID   = "236c0fa40bba4ce28d8e65598854e9e7"
	ethPortID2  = "4b1a3ed44c15438e957781ec4bb1c1e9"
	ethPortName = "BaseEnclosure-NodeA-IoModule1-FEPort0"
)

func TestClientIMPL_GetEthPorts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, ethPortID, ethPortID2)
	httpmock.RegisterResponder("GET", ethPortMockURL,
		httpmock.NewStringResponder(200, respData))
	ports, err := C.GetEthPorts(context.Background())
	assert.Nil(t, err)
	assert.Len(t, ports, 2)
	assert.Equal(t, ethPortID, ports[0].ID)
}

func TestClientIMPL_GetEthPort(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s", "name": "%s", "is_link_up": true}`, ethPortID, ethPortName)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", ethPortMockURL, ethPortID),
		httpmock.NewStringResponder(200, respData))
	port, err := C.GetEthPort(context.Background(), ethPortID)
	assert.Nil(t, err)
	assert.Equal(t, ethPortID, port.ID)
	assert.Equal(t, ethPortName, port.Name)
	assert.True(t, port.IsLinkUp)
}

func TestClientIMPL_GetEthPortByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", ethPortMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s", "name": "%s"}]`, ethPortID, ethPortName)
	setResponder(respData)
	port, err := C.GetEthPortByName(context.Background(), ethPortName)
	assert.Nil(t, err)
	assert.Equal(t, ethPortID, port.ID)
	assert.Equal(t, ethPortName, port.Name)
	httpmock.Reset()
	setResponder("[]")
	_, err = C.GetEthPortByName(context.Background(), "nonexistent")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetEthPortNotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := `{"messages":[{"severity":"Error","code":"0xE04040020009","message_l10n":"Instance with id invalid-id was not found."}]}`
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", ethPortMockURL, "invalid-id"),
		httpmock.NewStringResponder(404, respData))
	_, err := C.GetEthPort(context.Background(), "invalid-id")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_ModifyEthPort(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", ethPortMockURL, ethPortID),
		httpmock.NewStringResponder(204, ""))

	requestedSpeed := EthPortSpeedEnumAuto
	modifyParams := EthPortModify{
		RequestedSpeed: &requestedSpeed,
	}

	resp, err := C.ModifyEthPort(context.Background(), &modifyParams, ethPortID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_GetEthPortWithL2Discovery(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{
		"id": "%s",
		"name": "%s",
		"is_link_up": true,
		"current_speed": "25_Gbps",
		"l2_discovery_details": {
			"remote_mac": "3c:2c:30:59:7a:00",
			"remote_port_name": "ethernet1/1/7",
			"remote_name": "drm-pod1-25g-a1",
			"remote_mtu": 9216,
			"remote_native_vlan": 1212
		}
	}`, ethPortID, ethPortName)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", ethPortMockURL, ethPortID),
		httpmock.NewStringResponder(200, respData))
	port, err := C.GetEthPort(context.Background(), ethPortID)
	assert.Nil(t, err)
	assert.Equal(t, ethPortID, port.ID)
	assert.Equal(t, EthPortSpeedEnum25Gbps, port.CurrentSpeed)
	assert.Equal(t, "3c:2c:30:59:7a:00", port.L2DiscoveryDetails.RemoteMac)
	assert.Equal(t, "ethernet1/1/7", port.L2DiscoveryDetails.RemotePortName)
	assert.Equal(t, "drm-pod1-25g-a1", port.L2DiscoveryDetails.RemoteName)
	assert.Equal(t, int32(9216), port.L2DiscoveryDetails.RemoteMTU)
	assert.Equal(t, int32(1212), port.L2DiscoveryDetails.RemoteNativeVlan)
}
