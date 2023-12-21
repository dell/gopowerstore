/*
 *
 * Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	k8sClusterURL = "k8s_cluster"
)

// RegisterK8sCluster registers K8s cluster with PowerStore array
func (c *ClientIMPL) RegisterK8sCluster(ctx context.Context,
	createParams *K8sCluster,
) (resp CreateResponse, err error) {
	defaultHeaders := c.GetCustomHTTPHeaders()
	if defaultHeaders == nil {
		defaultHeaders = make(http.Header)
	}

	customHeaders := defaultHeaders

	customHeaders.Add("DELL-VISIBILITY", "internal")
	c.SetCustomHTTPHeaders(customHeaders)

	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: k8sClusterURL,
			Body:     createParams,
		},
		&resp)
	if err != nil {
		logrus.Error(err.Error())
	}

	// reset custom header
	customHeaders.Del("DELL-VISIBILITY")
	c.SetCustomHTTPHeaders(customHeaders)
	logrus.Info("default headers: ", customHeaders)

	return resp, WrapErr(err)
}
