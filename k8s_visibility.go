package gopowerstore

import (
	"context"
	"net/http"
)

const (
	k8sClusterURL = "k8s_cluster"
)

func (c *ClientIMPL) RegisterK8sCluster(ctx context.Context,
	createParams *K8sCluster) (resp CreateResponse, err error) {
	customHeader := http.Header{
		"DELL-VISIBILITY": []string{"internal"},
	}
	c.SetCustomHTTPHeaders(customHeader)
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: k8sClusterURL,
			Body:     createParams},
		&resp)

	// reset custom header
	customHeader = http.Header{}
	c.SetCustomHTTPHeaders(customHeader)

	return resp, WrapErr(err)
}
