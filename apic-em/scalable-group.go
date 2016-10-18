package apicem

import (
	"fmt"
	"strings"
)

const scalableGroupBasePath = "v1"

// ScalableGroupService is an interface with the Vscalablegroup API
type ScalableGroupService service

// ScalableGroupListResult is ...
type ScalableGroupListResult struct {
	Response []ScalableGroupDto `json:"response,omitempty"`
	Version  string             `json:"version,omitempty"`
}

// ScalableGroupResult is ...
type ScalableGroupResult struct {
	Response ScalableGroupDto `json:"response,omitempty"`
	Version  string           `json:"version,omitempty"`
}

// GetCount is ...
// This method is used to obtain the number of scalable groups known to APIC-EM
//
//
//
//
// * @return *CountResult
func (s *ScalableGroupService) GetCount() (*CountResult, *Response, error) {

	path := scalableGroupBasePath + "/v2/scalable-group/count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CountResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetEndPointGroupbyID is ...
// Get Scalable group based on id
//
//
//
//
// * @return *ScalableGroupResult
func (s *ScalableGroupService) GetEndPointGroupbyID(id string) (*ScalableGroupResult, *Response, error) {

	path := scalableGroupBasePath + "/v2/scalable-group/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ScalableGroupResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetScalableGroupByFiltersQueryParams is ...
type GetScalableGroupByFiltersQueryParams struct {
	name   string `url:"name,omitempty"`   // Retrieve policies for a given name
	offset string `url:"offset,omitempty"` // Starting index of the resources (1 based)
	limit  string `url:"limit,omitempty"`  // Number of resources to return
}

// GetScalableGroupByFilters is ...
// Retrieves scalable group based on a given filter
//
//  * @param queryParams
//
//
// * @return *ScalableGroupListResult
func (s *ScalableGroupService) GetScalableGroupByFilters(queryParams *GetScalableGroupByFiltersQueryParams) (*ScalableGroupListResult, *Response, error) {

	path := scalableGroupBasePath + "/v2/scalable-group"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ScalableGroupListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
