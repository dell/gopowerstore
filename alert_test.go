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
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/dell/gopowerstore/api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClientIMPL_GetAlerts(t *testing.T) {
	activeUnacknowledgedAlert := Alert{
		ID:                 "1234",
		EventCode:          "001",
		Severity:           "info",
		ResourceType:       "metro_session",
		ResourceName:       "csivol-001abc",
		Description:        "some description",
		GeneratedTimestamp: "2020-01-01T00:00:00Z",
		RaisedTimestamp:    "2020-01-01T00:00:00Z",
		ClearedTimestamp:   "",
		State:              "Active",
		IsAcknowledged:     false,
		Events:             []Event{},
	}
	clearedUnacknowledgedAlert := Alert{
		ID:                 "1001",
		EventCode:          "001",
		Severity:           "info",
		ResourceType:       "metro_session",
		ResourceName:       "csivol-001abc",
		Description:        "some else went wrong",
		GeneratedTimestamp: "2020-01-01T00:00:00Z",
		RaisedTimestamp:    "2020-01-01T00:00:00Z",
		ClearedTimestamp:   "2020-01-02T00:00:00Z",
		State:              "Cleared",
		IsAcknowledged:     false,
		Events:             []Event{},
	}
	activeAcknowledgedAlert := Alert{
		ID:                 "1234",
		EventCode:          "001",
		Severity:           "info",
		ResourceType:       "metro_session",
		ResourceName:       "csivol-001abc",
		Description:        "some description",
		GeneratedTimestamp: "2020-01-01T00:00:00Z",
		RaisedTimestamp:    "2020-01-01T00:00:00Z",
		ClearedTimestamp:   "",
		State:              "Active",
		IsAcknowledged:     true,
		Events:             []Event{},
	}
	clearedAcknowledgedAlert := Alert{
		ID:                 "1001",
		EventCode:          "001",
		Severity:           "info",
		ResourceType:       "metro_session",
		ResourceName:       "csivol-001abc",
		Description:        "some else went wrong",
		GeneratedTimestamp: "2020-01-01T00:00:00Z",
		RaisedTimestamp:    "2020-01-01T00:00:00Z",
		ClearedTimestamp:   "2020-01-02T00:00:00Z",
		State:              "Cleared",
		IsAcknowledged:     true,
		Events:             []Event{},
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		opts          GetAlertsOpts
		mockResponder httpmock.Responder
		want          *GetAlertsResponse
		wantErr       bool
	}{
		{
			name: "no options",
			opts: GetAlertsOpts{},
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					activeUnacknowledgedAlert,
					clearedUnacknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockResponse))
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta{
					api.RespMeta{
						Status: http.StatusOK,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      0,
							IsPaginate: false,
						},
					},
				},
				Alerts{
					activeUnacknowledgedAlert,
					clearedUnacknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "get active alerts",
			opts: func() GetAlertsOpts {
				return GetAlertsOpts{
					Queries: map[string]string{"state": "Active"},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					activeUnacknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockResponse))
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta: AlertsResponseMeta{
					RespMeta: api.RespMeta{
						Status: http.StatusOK,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      0,
							IsPaginate: false,
						},
					},
				},
				Alerts: Alerts{
					activeUnacknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "get cleared alerts",
			opts: func() GetAlertsOpts {
				return GetAlertsOpts{
					Queries: map[string]string{"state": "Cleared"},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					clearedUnacknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockResponse))
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta{
					api.RespMeta{
						Status: http.StatusOK,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      0,
							IsPaginate: false,
						},
					},
				},
				Alerts{
					clearedUnacknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "get acknowledged alerts",
			opts: func() GetAlertsOpts {
				return GetAlertsOpts{
					Queries: map[string]string{"is_acknowledged": "true"},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					clearedAcknowledgedAlert,
					activeAcknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockResponse))
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta{
					api.RespMeta{
						Status: http.StatusOK,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      0,
							IsPaginate: false,
						},
					},
				},
				Alerts{
					clearedAcknowledgedAlert,
					activeAcknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "get unacknowledged alerts",
			opts: func() GetAlertsOpts {
				return GetAlertsOpts{
					Queries: map[string]string{"is_acknowledged": "false"},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					clearedUnacknowledgedAlert,
					activeUnacknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockResponse))
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta{
					api.RespMeta{
						Status: http.StatusOK,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      0,
							IsPaginate: false,
						},
					},
				},
				Alerts{
					clearedUnacknowledgedAlert,
					activeUnacknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "get paginated alerts",
			opts: GetAlertsOpts{
				RequestPagination: RequestPagination{
					PageSize:   2,
					StartIndex: 0,
				},
			},
			mockResponder: func() httpmock.Responder {
				mockResponse, err := json.Marshal(Alerts{
					clearedUnacknowledgedAlert,
					activeUnacknowledgedAlert,
				})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusPartialContent, string(mockResponse)).HeaderAdd(http.Header{"content-range": {"0-1/4"}})
			}(),
			want: &GetAlertsResponse{
				AlertsResponseMeta{
					api.RespMeta{
						Status: http.StatusPartialContent,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       1,
							Next:       2,
							Total:      4,
							IsPaginate: true,
						},
					},
				},
				Alerts{
					clearedUnacknowledgedAlert,
					activeUnacknowledgedAlert,
				},
			},
			wantErr: false,
		},
		{
			name: "error response",
			opts: GetAlertsOpts{},
			mockResponder: func() httpmock.Responder {
				return httpmock.NewStringResponder(500, "")
			}(),
			want:    &GetAlertsResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize http mocking for REST requests
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", alertEndpoint, tt.mockResponder)

			// create mock client
			options := NewClientOptions()
			options.SetDefaultTimeout(1 * time.Second)
			client := api.MockClient(options.DefaultTimeout(), options.RateLimit(), options.RequestIDKey())
			c := &ClientIMPL{
				API: client,
			}

			got, gotErr := c.GetAlerts(context.Background(), tt.opts)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAlerts() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAlerts() succeeded unexpectedly")
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
