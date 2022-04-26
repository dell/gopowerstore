package gopowerstore

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	remoteSystemURL = "remote_system"
	clusterURL      = "cluster"
)

// GetCluster returns info about first cluster found
func (c *ClientIMPL) GetCluster(ctx context.Context) (resp Cluster, err error) {
	var systemList []Cluster
	cluster := Cluster{}
	qp := c.APIClient().QueryParamsWithFields(&cluster)

	majorMinorVersion, err := c.GetSoftwareMajorMinorVersion(ctx)
	if err != nil {
		log.Errorf("Couldn't find the major array version %s", err.Error())
	} else {
		if majorMinorVersion >= 3.1 {
			qp.Select("nvm_subsystem_nqn")
		}
	}
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    clusterURL,
			QueryParams: qp,
		},
		&systemList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	return systemList[0], err
}

// GetRemoteSystem query and return specific remote system by id
func (c *ClientIMPL) GetRemoteSystem(ctx context.Context, id string) (resp RemoteSystem, err error) {
	sys := RemoteSystem{}
	qp := c.APIClient().QueryParamsWithFields(&sys)
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    remoteSystemURL,
			ID:          id,
			QueryParams: qp,
		},
		&resp)
	return resp, WrapErr(err)
}

// GetRemoteSystemByName query and return specific remote system by name
func (c *ClientIMPL) GetRemoteSystemByName(ctx context.Context, name string) (resp RemoteSystem, err error) {
	var systemList []RemoteSystem
	sys := RemoteSystem{}
	qp := c.APIClient().QueryParamsWithFields(&sys)
	qp.RawArg("name", fmt.Sprintf("eq.%s", name))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    remoteSystemURL,
			QueryParams: qp},
		&systemList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(systemList) != 1 {
		return resp, NewHostIsNotExistError()
	}
	return systemList[0], err
}

// Queries all Remote Systems
func (c *ClientIMPL) GetAllRemoteSystems(ctx context.Context) (resp []RemoteSystem, err error) {
	sys := RemoteSystem{}
	var retsys []RemoteSystem
	qp := c.APIClient().QueryParamsWithFields(&sys)
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    remoteSystemURL,
			QueryParams: qp,
		},
		&retsys)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	return retsys, err
}
