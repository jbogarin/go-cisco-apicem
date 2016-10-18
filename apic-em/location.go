package apicem

import (
	"fmt"
	"strings"
)

const locationBasePath = "v1"

// LocationService is an interface with the  API
type LocationService service

// LocationNio is ...
type LocationNio struct {
	CivicAddress        string `json:"civicAddress,omitempty"`        // Civic address of the location
	Description         string `json:"description,omitempty"`         // Description of the location
	LocationName        string `json:"locationName,omitempty"`        // Name of the location
	ID                  string `json:"id,omitempty"`                  // Unique identifier for location
	GeographicalAddress string `json:"geographicalAddress,omitempty"` // Geographic address of the location
	Tag                 string `json:"tag,omitempty"`                 // Tag associated with the location
}

// LocationNioListResult is ...
type LocationNioListResult struct {
	Version  string        `json:"version,omitempty"`
	Response []LocationNio `json:"response,omitempty"`
}

// LocationNioResult is ...
type LocationNioResult struct {
	Version  string      `json:"version,omitempty"`
	Response LocationNio `json:"response,omitempty"`
}

// TagNio is ...
type TagNio struct {
	ID              string `json:"id,omitempty"`              // Unique identifier of tag
	LastUpdated     string `json:"lastUpdated,omitempty"`     // Time when the device interface info last got updated
	LocationID      string `json:"locationId,omitempty"`      // Unique identifier of location
	LinkID          string `json:"linkId,omitempty"`          // Unique identifier of link
	Tag             string `json:"tag,omitempty"`             // Tag ID
	InterfaceID     string `json:"interfaceId,omitempty"`     // Unique identifier of the interface
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Unique identifier of device
}

// AddLocation is ...
// Creates a new location with the attributes given
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param locationNIO Attributes of a location
// * @return *TaskIDResult
func (s *LocationService) AddLocation(locationNIO LocationNio, scope string) (*TaskIDResult, *Response, error) {

	path := locationBasePath + "/location"
	req, err := s.client.NewRequest("POST", path, locationNIO)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteLocationByID is ...
// Deletes the location with the given ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *LocationService) DeleteLocationByID(id string, scope string) (*TaskIDResult, *Response, error) {

	path := locationBasePath + "/location/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAllLocation is ...
// Gets list of locations
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *LocationNioListResult
func (s *LocationService) GetAllLocation(scope string) (*LocationNioListResult, *Response, error) {

	path := locationBasePath + "/location"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LocationNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLocationByID is ...
// Gets the location with the given ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *LocationNioResult
func (s *LocationService) GetLocationByID(id string, scope string) (*LocationNioResult, *Response, error) {

	path := locationBasePath + "/location/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LocationNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLocationByName is ...
// Gets the location with the given location name
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *LocationNioResult
func (s *LocationService) GetLocationByName(locationName string, scope string) (*LocationNioResult, *Response, error) {

	path := locationBasePath + "/location/location-name/{location-name}"
	path = strings.Replace(path, "{"+"location-name"+"}", fmt.Sprintf("%v", locationName), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LocationNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLocationByRange is ...
// Gets list of locations in the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *LocationNioListResult
func (s *LocationService) GetLocationByRange(startIndex int32, recordsToReturn int32, scope string) (*LocationNioListResult, *Response, error) {

	path := locationBasePath + "/location/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LocationNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLocationByTag is ...
// Gets list of locations associated with the given tag
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *LocationNioListResult
func (s *LocationService) GetLocationByTag(tagID string, scope string) (*LocationNioListResult, *Response, error) {

	path := locationBasePath + "/tag/{tagId}/location"
	path = strings.Replace(path, "{"+"tagId"+"}", fmt.Sprintf("%v", tagID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LocationNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLocationCount is ...
// Gets the count of all locations
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *LocationService) GetLocationCount(scope string) (*CountResult, *Response, error) {

	path := locationBasePath + "/location/count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(CountResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateLocation is ...
// Updates a location with the given attributes
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param locationNIO Attributes of a location
// * @return *TaskIDResult
func (s *LocationService) UpdateLocation(locationNIO LocationNio, scope string) (*TaskIDResult, *Response, error) {

	path := locationBasePath + "/location"
	req, err := s.client.NewRequest("PUT", path, locationNIO)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateLocationTag is ...
// Associates the given tag to the given location
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param tagNIO TagNIO with tag ID and location ID
// * @return *TaskIDResult
func (s *LocationService) UpdateLocationTag(tagNIO TagNio, scope string) (*TaskIDResult, *Response, error) {

	path := locationBasePath + "/location/tag"
	req, err := s.client.NewRequest("POST", path, tagNIO)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
