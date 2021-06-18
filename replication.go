package gopowerstore

import (
	"context"
	"fmt"
)

const (
	replicationRuleURL    = "replication_rule"
	policyURL             = "policy"
	replicationSessionURL = "replication_session"
)

// CreateReplicationRule creates new replication rule
func (c *ClientIMPL) CreateReplicationRule(ctx context.Context,
	createParams *ReplicationRuleCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: replicationRuleURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) GetReplicationRuleByName(ctx context.Context,
	ruleName string) (resp ReplicationRule, err error) {
	var ruleList []ReplicationRule
	rule := ReplicationRule{}
	qp := c.APIClient().QueryParamsWithFields(&rule)
	qp.RawArg("name", fmt.Sprintf("eq.%s", ruleName))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    replicationRuleURL,
			QueryParams: qp},
		&ruleList)

	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(ruleList) != 1 {
		return resp, replicationRuleNotExists()
	}
	return ruleList[0], nil
}

// CreateProtectionPolicy creates new protection policy
func (c *ClientIMPL) CreateProtectionPolicy(ctx context.Context,
	createParams *ProtectionPolicyCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: policyURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

func (c *ClientIMPL) GetProtectionPolicyByName(ctx context.Context,
	policyName string) (resp ProtectionPolicy, err error) {
	var policyList []ProtectionPolicy
	policy := ProtectionPolicy{}
	qp := c.APIClient().QueryParamsWithFields(&policy)
	qp.RawArg("name", fmt.Sprintf("eq.%s", policyName))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    policyURL,
			QueryParams: qp},
		&policyList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(policyList) != 1 {
		return resp, protectionPolicyNotExists()
	}
	return policyList[0], nil
}

func (c *ClientIMPL) GetReplicationSessionByLocalResourceID(ctx context.Context, id string) (resp ReplicationSession, err error) {
	var sessionList []ReplicationSession
	ses := ReplicationSession{}
	qp := c.APIClient().QueryParamsWithFields(&ses)
	qp.RawArg("local_resource_id", fmt.Sprintf("eq.%s", id))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    replicationSessionURL,
			QueryParams: qp},
		&sessionList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(sessionList) != 1 {
		return resp, replicationGroupNotExists()
	}
	return sessionList[0], err
}

// DeleteReplicationRule deletes existing RR
func (c *ClientIMPL) DeleteReplicationRule(ctx context.Context, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: replicationRuleURL,
			ID:       id,
		},
		&resp)
	return resp, WrapErr(err)
}

// DeleteProtectionPolicy deletes existing PP
func (c *ClientIMPL) DeleteProtectionPolicy(ctx context.Context, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: policyURL,
			ID:       id,
		},
		&resp)
	return resp, WrapErr(err)
}
