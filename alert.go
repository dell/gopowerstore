package gopowerstore

import (
	"context"
	"fmt"

	"github.com/dell/gopowerstore/api"
)

const (
	alertURL = "alert"
)

func getAlertsDefaultQueryParams(c Client) api.QueryParamsEncoder {
	alert := Alert{}
	return c.APIClient().QueryParamsWithFields(&alert)
}

// GetVolumes returns a list of volumes
func (c *ClientIMPL) GetAlerts(ctx context.Context) ([]Alert, error) {
	var result []Alert
	qp := getAlertsDefaultQueryParams(c)
	qp.Order("generated_timestamp")

	// In order to retrieve based on the resource_type. You can also exclude specific resource_type by using 'neq.'
	filterType := "ACTIVE"
	qp.RawArg("state", fmt.Sprintf("eq.%s", filterType))

	_, err := c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    alertURL,
			QueryParams: qp,
		},
		&result)
	err = WrapErr(err)
	if err != nil {
		return nil, err
	}

	return result, nil
}
