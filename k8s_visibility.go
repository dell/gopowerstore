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
	createParams *K8sCluster) (resp CreateResponse, err error) {
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
			Body:     createParams},
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
