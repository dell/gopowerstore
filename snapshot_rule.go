package gopowerstore

import (
	"context"

	"github.com/dell/gopowerstore/api"
)

const (
	snapshotRuleURL = "snapshot_rule"
)

func getSnapshotRuleDefaultQueryParams(c Client) api.QueryParamsEncoder {
	snapshotRule := SnapshotRule{}
	return c.APIClient().QueryParamsWithFields(&snapshotRule)
}

// GetSnapshotRule query and return specific snapshot rule by id
func (c *ClientIMPL) GetSnapshotRule(ctx context.Context, id string) (resp SnapshotRule, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    snapshotRuleURL,
			ID:          id,
			QueryParams: getSnapshotRuleDefaultQueryParams(c)},
		&resp)
	return resp, WrapErr(err)
}

// GetSnapshotRules returns a list of snapshot rules
func (c *ClientIMPL) GetSnapshotRules(ctx context.Context) ([]SnapshotRule, error) {
	var result []SnapshotRule
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []SnapshotRule
		qp := getSnapshotRuleDefaultQueryParams(c)
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    snapshotRuleURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// CreateSnapshotRule creates new snapshot rule
func (c *ClientIMPL) CreateSnapshotRule(ctx context.Context,
	createParams *SnapshotRuleCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: snapshotRuleURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

// DeleteSnapshotRule deletes existing snapshot rule
func (c *ClientIMPL) DeleteSnapshotRule(ctx context.Context,
	deleteParams *SnapshotRuleDelete, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: snapshotRuleURL,
			ID:       id,
			Body:     deleteParams},
		&resp)
	return resp, WrapErr(err)
}
