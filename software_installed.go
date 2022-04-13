package gopowerstore

import (
	"context"
	"github.com/dell/gopowerstore/api"
)

const apiSoftwareInstalledURL = "software_installed"

func getSoftwareInstalledDefaultQueryParams(c Client) api.QueryParamsEncoder {
	softwareInstalled := SoftwareInstalled{}
	return c.APIClient().QueryParamsWithFields(&softwareInstalled)
}

// GetSoftwareInstalled queries the software packages that are installed on each appliance, or on the cluster as a whole
func (c *ClientIMPL) GetSoftwareInstalled(
	ctx context.Context) (resp []SoftwareInstalled, err error) {
	err = c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []SoftwareInstalled
		qp := getSoftwareInstalledDefaultQueryParams(c)
		qp.Limit(paginationDefaultPageSize)
		qp.Offset(offset)
		qp.Order("id")
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    apiSoftwareInstalledURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			resp = append(resp, page...)
		}
		return meta, err
	})
	return resp, err
}
