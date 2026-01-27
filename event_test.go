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

func TestClientIMPL_GetEvents(t *testing.T) {
	volumeResource := "csivol-0001abc"
	event := Event{
		ID:           "test",
		EventCode:    "0001",
		Severity:     "minor",
		ResourceName: volumeResource,
		Description:  "some description",
		Timestamp:    "2020-01-01T00:00:00Z",
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		opts          GetEventsOpts
		mockResponder httpmock.Responder
		want          *GetEventsResponse
		wantErr       bool
	}{
		{
			name: "no options",
			opts: GetEventsOpts{},
			mockResponder: func() httpmock.Responder {
				mockEventsResp, err := json.Marshal(Events{event})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockEventsResp))
			}(),
			want: &GetEventsResponse{
				EventsResponseMeta{
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
				Events{event},
			},
			wantErr: false,
		},
		{
			name: "with resource name",
			opts: func() GetEventsOpts {
				return GetEventsOpts{
					Queries: map[string]string{"resource_name": volumeResource},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockEventsResp, err := json.Marshal(Events{event})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				return httpmock.NewStringResponder(http.StatusOK, string(mockEventsResp))
			}(),
			want: &GetEventsResponse{
				EventsResponseMeta{
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
				Events{event},
			},
			wantErr: false,
		},
		{
			name: "pagination",
			opts: func() GetEventsOpts {
				return GetEventsOpts{
					RequestPagination: RequestPagination{
						PageSize:   1,
						StartIndex: 0,
					},
					Queries: map[string]string{"resource_name": volumeResource},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockEventsResp, err := json.Marshal(Events{event})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				// indicate there are more results to be read
				return httpmock.NewStringResponder(http.StatusPartialContent, string(mockEventsResp)).HeaderAdd(http.Header{"content-range": {"0-0/2"}})
			}(),
			want: &GetEventsResponse{
				EventsResponseMeta{
					api.RespMeta{
						Status: http.StatusPartialContent,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       1,
							Total:      2,
							IsPaginate: true,
						},
					},
				},
				Events{event},
			},
			wantErr: false,
		},
		{
			name: "pagination result returns last page",
			opts: func() GetEventsOpts {
				return GetEventsOpts{
					RequestPagination: RequestPagination{
						PageSize:   1,
						StartIndex: 0,
					},
					Queries: map[string]string{"resource_name": volumeResource},
				}
			}(),
			mockResponder: func() httpmock.Responder {
				mockEventsResp, err := json.Marshal(Events{event})
				if err != nil {
					t.Fatalf("unable to marshal mock response to json: %v", err)
				}
				// indicate we've reached the end of the results
				return httpmock.NewStringResponder(http.StatusPartialContent, string(mockEventsResp)).HeaderAdd(http.Header{"content-range": {"0-0/1"}})
			}(),
			want: &GetEventsResponse{
				EventsResponseMeta{
					api.RespMeta{
						Status: http.StatusPartialContent,
						Pagination: api.PaginationInfo{
							First:      0,
							Last:       0,
							Next:       0,
							Total:      1,
							IsPaginate: true,
						},
					},
				},
				Events{event},
			},
			wantErr: false,
		},
		{
			name: "error response",
			opts: GetEventsOpts{},
			mockResponder: func() httpmock.Responder {
				return httpmock.NewStringResponder(500, "")
			}(),
			want:    &GetEventsResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize http mocking for REST requests
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", eventEndpoint, tt.mockResponder)

			// create mock client
			options := NewClientOptions()
			options.SetDefaultTimeout(1 * time.Second)
			client := api.MockClient(options.DefaultTimeout(), options.RateLimit(), options.RequestIDKey())
			c := &ClientIMPL{
				API: client,
			}

			got, gotErr := c.GetEvents(context.Background(), tt.opts)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetEvents() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetEvents() succeeded unexpectedly")
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
