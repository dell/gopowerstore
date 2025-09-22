package gopowerstore

import (
	"context"

	"github.com/dell/gopowerstore/api"
)

const (
	eventURL = "event"
)

func getEventsDefaultQueryParams(c Client) api.QueryParamsEncoder {
	event := Event{}
	return c.APIClient().QueryParamsWithFields(&event)
}

// GetVolumes returns a list of volumes
func (c *ClientIMPL) GetEvents(ctx context.Context) ([]Event, error) {
	var result []Event
	qp := getEventsDefaultQueryParams(c)
	qp.Order("generated_timestamp")
	// // In order to retrieve based on the resource_type. You can also exclude specific resource_type by using 'neq.'
	// filterType := "volume_group"
	// qp.RawArg("resource_type", fmt.Sprintf("eq.%s", filterType))
	_, err := c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    eventURL,
			QueryParams: qp,
		},
		&result)
	err = WrapErr(err)
	if err != nil {
		return nil, err
	}

	return result, nil
}
