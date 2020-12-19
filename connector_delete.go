package afas

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (s *ConnectorService) NewDeleteRequest() ConnectorDeleteRequest {
	return ConnectorDeleteRequest{
		api:          s.api,
		method:       http.MethodDelete,
		urlParams:    ConnectorDeleteURLParams{},
		queryParams:  ConnectorDeleteQueryParams{},
		requestBody:  s.NewDeleteRequestBody(),
		responseBody: s.NewDeleteResponseBody(),
	}
}

type ConnectorDeleteRequest struct {
	api          *API
	method       string
	urlParams    ConnectorDeleteURLParams
	queryParams  ConnectorDeleteQueryParams
	requestBody  ConnectorDeleteRequestBody
	responseBody *ConnectorDeleteResponseBody
}

func (r *ConnectorDeleteRequest) Method() string {
	return r.method
}

func (r *ConnectorDeleteRequest) SetMethod(method string) {
	r.method = method
}

func (r *ConnectorDeleteRequest) RequestBody() ConnectorDeleteRequestBody {
	return r.requestBody
}

func (r *ConnectorDeleteRequest) SetRequestBody(body ConnectorDeleteRequestBody) {
	r.requestBody = body
}

func (r *ConnectorDeleteRequest) ResponseBody() *ConnectorDeleteResponseBody {
	return r.responseBody
}

func (r *ConnectorDeleteRequest) URL() url.URL {
	path := "/connectors/{connectorid}[/{subelement}]"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	path = strings.Replace(path, "[/{subelement}]", r.urlParams.SubElement, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ConnectorDeleteRequest) Do() (*ConnectorDeleteResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return r.ResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return r.ResponseBody(), err
	}

	responseBody := r.ResponseBody()
	_, err = r.api.Do(req, responseBody)
	return responseBody, err
}

func (s *ConnectorService) NewDeleteResponseBody() *ConnectorDeleteResponseBody {
	return &ConnectorDeleteResponseBody{}
}

func (r *ConnectorDeleteRequest) QueryParams() *ConnectorDeleteQueryParams {
	return &r.queryParams
}

type ConnectorDeleteQueryParams struct {
}

func (r *ConnectorDeleteRequest) URLParams() *ConnectorDeleteURLParams {
	return &r.urlParams
}

type ConnectorDeleteURLParams struct {
	ConnectorID string
	SubElement  string
}

func (s *ConnectorService) NewDeleteRequestBody() ConnectorDeleteRequestBody {
	return struct{}{}
}

type ConnectorDeleteRequestBody interface{}

// {"results":{"FiEntryPar":{"UnId":"4","EnNo":"54605"}}}
type ConnectorDeleteResponseBody struct {
	Results json.RawMessage `json:"results"`
}
