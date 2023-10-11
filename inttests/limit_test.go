package inttests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxVolumeSize(t *testing.T) {
	customHeaders := C.GetCustomHTTPHeaders()
	if customHeaders == nil {
		customHeaders = make(http.Header)
	}
	customHeaders.Add("DELL-VISIBILITY", "internal")
	C.SetCustomHTTPHeaders(customHeaders)

	limit, err := C.GetMaxVolumeSize(context.Background())

	checkAPIErr(t, err)
	assert.Positive(t, limit)
}

func TestGetMaxVolumeSizeEndpointNotFound(t *testing.T) {
	limit, err := C.GetMaxVolumeSize(context.Background())

	assert.Equal(t, "The REST endpoint [GET /api/rest/limit?select=id%2Climit] cannot be found.", err.Error())
	assert.Negative(t, limit)
}
