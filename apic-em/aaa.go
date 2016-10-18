package apicem

import (
	"fmt"
	"strings"
)

const aaaBasePath = "v1"

// AAAService is an interface with the AAA API
type AAAService service

// AaaAttribute is ...
type AaaAttribute struct {
	AttributeName string `json:"attributeName,omitempty"` // Attributename
}

// AaaAttributeResult is ...
type AaaAttributeResult struct {
	Version  string       `json:"version,omitempty"`
	Response AaaAttribute `json:"response,omitempty"`
}

// AaaServer is ...
type AaaServer struct {
	Protocol           string `json:"protocol,omitempty"`           // Protocol
	AuthenticationPort int32  `json:"authenticationPort,omitempty"` // Authentication Port
	AccountingPort     int32  `json:"accountingPort,omitempty"`     // Accounting Port
	Retries            int32  `json:"retries,omitempty"`            // Retries
	SocketTimeout      int32  `json:"socketTimeout,omitempty"`      // Timeout
	ServerID           string `json:"serverId,omitempty"`           // Server Id
	ServerIP           string `json:"serverIp,omitempty"`           // Server IP Address
	SharedSecret       string `json:"sharedSecret,omitempty"`       // Shared Secret
}

// AaaServerListResult is ...
type AaaServerListResult struct {
	Version  string      `json:"version,omitempty"`
	Response []AaaServer `json:"response,omitempty"`
}

// AaaServerResult is ...
type AaaServerResult struct {
	Version  string    `json:"version,omitempty"`
	Response AaaServer `json:"response,omitempty"`
}

// ServerID is ...
type ServerID struct {
	ServerID string `json:"serverId,omitempty"`
}

// ServerIDListResult is ...
type ServerIDListResult struct {
	Version  string     `json:"version,omitempty"`
	Response []ServerID `json:"response,omitempty"`
}

// AddAAAAttribute is ...
// This method is used to add a custom AAA Attribute.
//
//
//
//  * @param aaaAttribute aaaAttribute
// * @return *SuccessResult
func (s *AAAService) AddAAAAttribute(aaaAttribute AaaAttribute) (*SuccessResult, *Response, error) {

	path := aaaBasePath + "/aaa-server/authorization-attribute"
	req, err := s.client.NewRequest("POST", path, aaaAttribute)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// AddAAAServer is ...
// This method is used to add a list of AAA servers.
//
//
//
//  * @param aaaServerList aaaServerList
// * @return *ServerIDListResult
func (s *AAAService) AddAAAServer(aaaServerList *[]AaaServer) (*ServerIDListResult, *Response, error) {

	path := aaaBasePath + "/aaa-server"
	req, err := s.client.NewRequest("POST", path, aaaServerList)
	if err != nil {
		return nil, nil, err
	}

	root := new(ServerIDListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteAAAAttribute is ...
// This method is used to delete a custom AAA Attribute.
//
//
//
//
// * @return *SuccessResult
func (s *AAAService) DeleteAAAAttribute() (*SuccessResult, *Response, error) {

	path := aaaBasePath + "/aaa-server/authorization-attribute"
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteServer is ...
// This method is used to delete a registered AAA server.
//
//
//
//
// * @return *SuccessResult
func (s *AAAService) DeleteServer(serverID string) (*SuccessResult, *Response, error) {

	path := aaaBasePath + "/aaa-server/{serverId}"
	path = strings.Replace(path, "{"+"serverId"+"}", fmt.Sprintf("%v", serverID), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAAAAttribute is ...
// This method is used to get all custom AAA Attributes.
//
//
//
//
// * @return *AaaAttributeResult
func (s *AAAService) GetAAAAttribute() (*AaaAttributeResult, *Response, error) {

	path := aaaBasePath + "/aaa-server/authorization-attribute"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AaaAttributeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetServer is ...
// This method is used to get a registered AAA server.
//
//
//
//
// * @return *AaaServerResult
func (s *AAAService) GetServer(serverID string) (*AaaServerResult, *Response, error) {

	path := aaaBasePath + "/aaa-server/{serverId}"
	path = strings.Replace(path, "{"+"serverId"+"}", fmt.Sprintf("%v", serverID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AaaServerResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetServers is ...
// This method is used to get a list of registered AAA servers.
//
//
//
//
// * @return *AaaServerListResult
func (s *AAAService) GetServers() (*AaaServerListResult, *Response, error) {

	path := aaaBasePath + "/aaa-server"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AaaServerListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateServers is ...
// This method is used to update an individual or a list of AAA servers.
//
//
//
//  * @param aaaServerList aaaServerList
// * @return *SuccessResult
func (s *AAAService) UpdateServers(aaaServerList *[]AaaServer) (*SuccessResult, *Response, error) {

	path := aaaBasePath + "/aaa-server"
	req, err := s.client.NewRequest("PUT", path, aaaServerList)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
