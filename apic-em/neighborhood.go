package apicem

import (
	"fmt"
	"strings"
)

const neighborhoodBasePath = "v1"

// NeighborhoodService is an interface with the Neighborhood API
type NeighborhoodService service

// NeighborhoodDto is ...
type NeighborhoodDto struct {
	Description          string   `json:"description,omitempty"`          // Description of the group
	ListofScalableGroups []string `json:"listofScalableGroups,omitempty"` // List of Scalable Groups
	Name                 string   `json:"name,omitempty"`                 // Neighborhood Name
	ID                   string   `json:"id,omitempty"`                   // UUID of the Neighborhood
	Type                 string   `json:"type,omitempty"`                 // Isolation Property
}

// NeighborhoodListResult is ...
type NeighborhoodListResult struct {
	Response []NeighborhoodDto `json:"response,omitempty"`
	Version  string            `json:"version,omitempty"`
}

// NeighborhoodResult is ...
type NeighborhoodResult struct {
	Response NeighborhoodDto `json:"response,omitempty"`
	Version  string          `json:"version,omitempty"`
}

// AddNeighbor is ...
// Create a Neighborhood
//
//
//
//  * @param nbr Neighborhood Object
// * @return *TaskIDResult
func (s *NeighborhoodService) AddNeighbor(nbr NeighborhoodDto) (*TaskIDResult, *Response, error) {

	path := neighborhoodBasePath + "/v2/neighborhood"
	req, err := s.client.NewRequest("POST", path, nbr)
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

// DeleteNeighbor is ...
// Delete specific Neighborhood
//
//
//
//
// * @return *TaskIDResult
func (s *NeighborhoodService) DeleteNeighbor(id string) (*TaskIDResult, *Response, error) {

	path := neighborhoodBasePath + "/v2/neighborhood/{id}"
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

// GetAllNeighbors is ...
// Get all Neighborhood(s)
//
//
//
//
// * @return *NeighborhoodListResult
func (s *NeighborhoodService) GetAllNeighbors() (*NeighborhoodListResult, *Response, error) {

	path := neighborhoodBasePath + "/v2/neighborhood"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(NeighborhoodListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNeighbor is ...
// Get specific Neighborhood
//
//
//
//
// * @return *NeighborhoodResult
func (s *NeighborhoodService) GetNeighbor(id string) (*NeighborhoodResult, *Response, error) {

	path := neighborhoodBasePath + "/v2/neighborhood/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(NeighborhoodResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateNeighbor is ...
// Update a Neighborhood
//
//
//
//  * @param nbr Neighborhood Object
// * @return *TaskIDResult
func (s *NeighborhoodService) UpdateNeighbor(nbr NeighborhoodDto) (*TaskIDResult, *Response, error) {

	path := neighborhoodBasePath + "/v2/neighborhood"
	req, err := s.client.NewRequest("PUT", path, nbr)
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
