package main

import (
	"encoding/json"
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"runtime"
	"strings"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"

	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"

	"github.com/pkg/errors"

	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
	"github.com/ucloud/ucloud-sdk-go/ucloud/version"
)

// SetupRequest will init request by client configuration
func (c *Client) SetupRequest(req request.Common) request.Common {
	req.SetRetryable(true)
	cfg := c.GetConfig()
	if cfg == nil {
		return req
	}

	// set optional client level variables
	if len(req.GetRegion()) == 0 && len(cfg.Region) > 0 {
		req.SetRegion(cfg.Region)
	}

	if len(req.GetZone()) == 0 && len(cfg.Zone) > 0 {
		req.SetZone(cfg.Zone)
	}

	if len(req.GetProjectId()) == 0 && len(cfg.ProjectId) > 0 {
		req.SetProjectId(cfg.ProjectId)
	}

	if req.GetTimeout() == 0 && cfg.Timeout != 0 {
		req.WithTimeout(cfg.Timeout)
	}

	if req.GetMaxretries() == 0 && cfg.MaxRetries != 0 {
		req.WithRetry(cfg.MaxRetries)
	}

	return req
}

func (c *Client) buildHTTPRequest(req request.Common) (*http.HttpRequest, error) {
	// convert request struct to query map
	query, err := request.ToQueryMap(req)

	if err != nil {
		return nil, errors.Errorf("convert request to map failed, %s", err)
	}

	credential := c.GetCredential()
	config := c.GetConfig()
	httpReq := http.NewHttpRequest()
	httpReq.SetURL(config.BaseUrl)
	httpReq.SetMethod("POST")

	// set timeout with client configuration
	httpReq.SetTimeout(config.Timeout)

	// keep query string is ordered and append credential signature as the last query param
	httpReq.SetQueryString(credential.BuildCredentialedQuery(query))

	ua := fmt.Sprintf("GO/%s GO-SDK/%s %s", runtime.Version(), version.Version, config.UserAgent)
	httpReq.SetHeader("User-Agent", strings.TrimSpace(ua))

	return httpReq, nil
}

// unmarshalHTTPResponse will get body from http response and unmarshal it's data into response struct
func (c *Client) unmarshalHTTPResponse(body []byte, resp response.Common) error {
	if len(body) == 0 {
		return uerr.NewEmptyResponseBodyError()
	}
	if r, ok := resp.(response.GenericResponse); ok {
		m := make(map[string]interface{})
		if err := json.Unmarshal(body, &m); err != nil {
			return uerr.NewResponseBodyError(err, string(body))
		}
		if err := r.SetPayload(m); err != nil {
			return uerr.NewResponseBodyError(err, string(body))
		}
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return uerr.NewResponseBodyError(err, string(body))
	}

	return nil
}
