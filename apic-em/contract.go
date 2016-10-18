package apicem

import (
	"fmt"
	"strings"
)

const contractBasePath = "v1"

// ContractService is an interface with the Vcontract API
type ContractService service

// AccessClauseDto is ...
type AccessClauseDto struct {
	Access string `json:"access,omitempty"` // access
}

// ContractDto is ...
type ContractDto struct {
	AccessClause   AccessClauseDto `json:"accessClause,omitempty"`   // accessClause
	Description    string          `json:"description,omitempty"`    // description
	Name           string          `json:"name,omitempty"`           // name
	ID             string          `json:"id,omitempty"`             // id
	CreateTime     int64           `json:"createTime,omitempty"`     // createTime
	LastUpdateTime int64           `json:"lastUpdateTime,omitempty"` // lastUpdateTime
}

// ContractListResult is ...
type ContractListResult struct {
	Response []ContractDto `json:"response,omitempty"`
	Version  string        `json:"version,omitempty"`
}

// ContractResult is ...
type ContractResult struct {
	Response ContractDto `json:"response,omitempty"`
	Version  string      `json:"version,omitempty"`
}

// Add is ...
// Create contract(s)
//
//
//
//  * @param contractDTOs Contract Object(s)
// * @return *TaskIDResult
func (s *ContractService) Add(contractDTOs *[]ContractDto) (*TaskIDResult, *Response, error) {

	path := contractBasePath + "/v2/contract"
	req, err := s.client.NewRequest("POST", path, contractDTOs)
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
// Delete a contract by id
//
//
//
//
// * @return *TaskIDResult
func (s *ContractService) Delete(id string) (*TaskIDResult, *Response, error) {

	path := contractBasePath + "/v2/contract/{id}"
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
// Get a contract by id
//
//
//
//
// * @return *ContractResult
func (s *ContractService) GetByID(id string) (*ContractResult, *Response, error) {

	path := contractBasePath + "/v2/contract/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ContractResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetContractByFiltersQueryParams is ...
type GetContractByFiltersQueryParams struct {
	Name         string `url:"name,omitempty"`         // Retrieve contracts for a given name
	AccessAction string `url:"accessAction,omitempty"` // Retrieve contracts for a given accessAction
	Offset       string `url:"offset,omitempty"`       // Starting index of the resources (1 based)
	Limit        string `url:"limit,omitempty"`        // Number of resources to return
}

// GetContractByFilters is ...
// Retrieves contracts based on a given filter
//
//  * @param queryParams
//
//
// * @return *ContractListResult
func (s *ContractService) GetContractByFilters(queryParams *GetContractByFiltersQueryParams) (*ContractListResult, *Response, error) {

	path := contractBasePath + "/v2/contract"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ContractListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCount is ...
// Get total count of contracts
//
//
//
//
// * @return *CountResult
func (s *ContractService) GetCount() (*CountResult, *Response, error) {

	path := contractBasePath + "/v2/contract/count"
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

// Update is ...
// Update contract(s)
//
//
//
//  * @param contractDTOs Contract Object(s)
// * @return *TaskIDResult
func (s *ContractService) Update(contractDTOs *[]ContractDto) (*TaskIDResult, *Response, error) {

	path := contractBasePath + "/v2/contract"
	req, err := s.client.NewRequest("PUT", path, contractDTOs)
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
