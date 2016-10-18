package apicem

import (
	"fmt"
	"strings"
)

const policyV2BasePath = "v1"

// PolicyV2Service is an interface with the Vpolicy API
type PolicyV2Service service

// ConsumerDto is ...
type ConsumerDto struct {
	ScalableGroups []ScalableGroupDto `json:"scalableGroups,omitempty"`
}

// PolicyContractDto is ...
type PolicyContractDto struct {
	ID   string `json:"id,omitempty"`   // id
	Name string `json:"name,omitempty"` // name
}

// PolicyDto is ...
type PolicyDto struct {
	Description    string            `json:"description,omitempty"`    // description
	Contract       PolicyContractDto `json:"contract,omitempty"`       // contract
	Producer       ProducerDto       `json:"producer,omitempty"`       // producer
	Consumer       ConsumerDto       `json:"consumer,omitempty"`       // consumer
	Priority       int32             `json:"priority,omitempty"`       // priority
	Name           string            `json:"name,omitempty"`           // name
	ID             string            `json:"id,omitempty"`             // id
	CreateTime     int64             `json:"createTime,omitempty"`     // createTime
	LastUpdateTime int64             `json:"lastUpdateTime,omitempty"` // lastUpdateTime
}

// PolicyDtoListResult is ...
type PolicyDtoListResult struct {
	Response []PolicyDto `json:"response,omitempty"`
	Version  string      `json:"version,omitempty"`
}

// PolicyDtoResult is ...
type PolicyDtoResult struct {
	Response PolicyDto `json:"response,omitempty"`
	Version  string    `json:"version,omitempty"`
}

// ProducerDto is ...
type ProducerDto struct {
	ScalableGroups []ScalableGroupDto `json:"scalableGroups,omitempty"`
}

// Add is ...
// Create Policy(s)
//
//
//
//  * @param policyDTOs Policy Object(s)
// * @return *TaskIDResult
func (s *PolicyV2Service) Add(policyDTOs *[]PolicyDto) (*TaskIDResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy"
	req, err := s.client.NewRequest("POST", path, policyDTOs)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete is ...
// Delete a policy by id
//
//
//
//
// * @return *TaskIDResult
func (s *PolicyV2Service) Delete(id string) (*TaskIDResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetByID is ...
// Get a policy by id
//
//
//
//
// * @return *PolicyDtoResult
func (s *PolicyV2Service) GetByID(id string) (*PolicyDtoResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyDtoResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCount is ...
// Get total count of policies
//
//
//
//
// * @return *CountResult
func (s *PolicyV2Service) GetCount() (*CountResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy/count"
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

// GetPolicyByFiltersQueryParams is ...
type GetPolicyByFiltersQueryParams struct {
	Name                      string `url:"name,omitempty"`                      // Retrieve policies for a given name
	Contractname              string `url:"contractname,omitempty"`              // Retrieve policies for a given contract name
	ProducerscalableGroupName string `url:"producerscalableGroupName,omitempty"` // Retrieve policies for a given producer scalable group name
	ConsumerscalableGroupName string `url:"consumerscalableGroupName,omitempty"` // Retrieve policies for a given consumer scalable group name
	ScalableGroupName         string `url:"scalableGroupName,omitempty"`         // Retrieve policies for a given scalable group name (used within producer or consumer)
	Offset                    string `url:"offset,omitempty"`                    // Starting index of the resources (1 based)
	Limit                     string `url:"limit,omitempty"`                     // Number of resources to return
}

// GetPolicyByFilters is ...
// Retrieves policies based on a given filter
//
//  * @param queryParams
//
//
// * @return *PolicyDtoListResult
func (s *PolicyV2Service) GetPolicyByFilters(queryParams *GetPolicyByFiltersQueryParams) (*PolicyDtoListResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyDtoListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update is ...
// Update Policy(s)
//
//
//
//  * @param policyDTOs Policy Object(s)
// * @return *TaskIDResult
func (s *PolicyV2Service) Update(policyDTOs *[]PolicyDto) (*TaskIDResult, *Response, error) {

	path := policyV2BasePath + "/v2/policy"
	req, err := s.client.NewRequest("PUT", path, policyDTOs)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
