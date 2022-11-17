/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"fmt"
)

type ActionType string

const (
	RS_ACTION_FAILOVER  ActionType = "failover"
	RS_ACTION_REPROTECT ActionType = "reprotect"
	RS_ACTION_RESUME    ActionType = "resume"
	RS_ACTION_PAUSE     ActionType = "pause"
	RS_ACTION_SYNC      ActionType = "sync"
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
	qp.Select("policies")
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
	qp.Select("name,id,replication_rules(id),volume(id,name),volume_group(id,name)")
	qp.RawArg("name", fmt.Sprintf("eq.%s", policyName))
	qp.RawArg("type", fmt.Sprintf("eq.%s", "Protection"))
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
func (c *ClientIMPL) GetReplicationSessionByID(ctx context.Context, id string) (resp ReplicationSession, err error) {
	var session ReplicationSession
	ses := ReplicationSession{}
	qp := c.APIClient().QueryParamsWithFields(&ses)
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    replicationSessionURL,
			QueryParams: qp,
			ID:          id},
		&session)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}

	return session, err
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

func (c *ClientIMPL) ExecuteActionOnReplicationSession(ctx context.Context, id string, actionType ActionType, params *FailoverParams) (resp EmptyResponse, err error) {
	var res interface{}
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: replicationSessionURL,
			ID:       id,
			Action:   string(actionType),
			Body:     params,
		},
		&res)
	return resp, WrapErr(err)
}
