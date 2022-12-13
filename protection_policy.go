package gopowerstore

import (
        "context"

        "github.com/dell/gopowerstore/api"
)

const (
        protectionPolicyURL = "policy"
)

func getProtectionPolicyDefaultQueryParams(c Client) api.QueryParamsEncoder {
        protectionPolicy := ProtectionPolicy{}
        return c.APIClient().QueryParamsWithFields(&protectionPolicy)
}

//GetProtectionPolicy query and return specific protection policy id
func (c *ClientIMPL) GetProtectionPolicy(ctx context.Context, id string) (resp ProtectionPolicy, err error) {
        _, err = c.APIClient().Query(
                ctx,
                RequestConfig{
                        Method:      "GET",
                        Endpoint:    protectionPolicyURL,
                        ID:          id,
                        QueryParams: getProtectionPolicyDefaultQueryParams(c)},
                &resp)
        return resp, WrapErr(err)
}
