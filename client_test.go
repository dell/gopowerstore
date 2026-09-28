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
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dell/gopowerstore/api"
	"github.com/stretchr/testify/assert"
)

var C Client

func initClient() {
	clientOptions := &ClientOptions{}
	clientOptions.SetDefaultTimeout(1 * time.Second)
	C = NewMockClient(clientOptions)
}

func init() {
	initClient()
}

func TestNewClient(t *testing.T) {
	os.Setenv(InsecureEnv, "true")
	os.Setenv(APIURLEnv, "api")
	os.Setenv(UsernameEnv, "admin")
	os.Setenv(PasswordEnv, "password")
	os.Setenv(HTTPTimeoutEnv, "120")
	_, err := NewClient()
	assert.Nil(t, err)
	os.Unsetenv(UsernameEnv)
	_, err = NewClient()
	assert.NotNil(t, err)
}

func TestClientIMPL_SetTraceID(t *testing.T) {
	ctx := context.Background()
	ctx = C.SetTraceID(ctx, "123")
	assert.Equal(t, "123", ctx.Value(api.ContextKey(clientOptionsDefaultRequestIDKey)))
}

type clientRecordingObserver struct {
	observations []api.RequestObservation
}

func (r *clientRecordingObserver) ObservePowerStoreRequest(obs api.RequestObservation) {
	r.observations = append(r.observations, obs)
}

func TestNewClientWithArgs_RequestObserver(t *testing.T) {
	observer := &clientRecordingObserver{}
	options := NewClientOptions()
	options.SetDefaultTimeout(1 * time.Second)
	options.SetRequestObserver(observer)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login_session":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
		case "/mock":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"name":"Foo"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client, err := NewClientWithArgs(server.URL, "admin", "password", options)
	assert.NoError(t, err)
	observer.observations = nil

	_, err = client.APIClient().Query(context.Background(), api.RequestConfig{
		Method:   "GET",
		Endpoint: "mock",
	}, &struct {
		Name string `json:"name"`
	}{})
	assert.NoError(t, err)

	assert.Len(t, observer.observations, 1)
	assert.Equal(t, "GET", observer.observations[0].Method)
	assert.Equal(t, "mock", observer.observations[0].Endpoint)
}
