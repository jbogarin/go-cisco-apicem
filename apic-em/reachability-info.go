package apicem

import (
	"fmt"
	"strings"
)

const reachabilityInfoBasePath = "v1"

// ReachabilityInfoService is an interface with the Reachabilityinfo API
type ReachabilityInfoService service

// NetworkDeviceReachabilityInfoNio is ...
type NetworkDeviceReachabilityInfoNio struct {
	SnmpCommunity             string `json:"snmpCommunity,omitempty"`             // SNMP community used for device connectivity
	MgmtIP                    string `json:"mgmtIp,omitempty"`                    // IP address of the device
	ProtocolList              string `json:"protocolList,omitempty"`              // Protocol order given for discovery
	ProtocolUsed              string `json:"protocolUsed,omitempty"`              // Protocol used for device discovery
	DiscoveryStartTime        string `json:"discoveryStartTime,omitempty"`        // Time when the discovery was started
	UserName                  string `json:"userName,omitempty"`                  // CLI username used for device connectivity
	Password                  string `json:"password,omitempty"`                  // CLI password used for device connectivity
	ID                        string `json:"id,omitempty"`                        // Unique identifier for reachability-info
	EnablePassword            string `json:"enablePassword,omitempty"`            // CLI enable password used for device connectivity
	DiscoveryID               string `json:"discoveryId,omitempty"`               // ID of discovery thorugh which device was discovered
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for failure if the device is not discovered successfully
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        // Reachability status of the device as Reachable / Unreachable
	AttributeInfo             string `json:"attributeInfo,omitempty"`
}

// NetworkDeviceReachabilityInfoNioListResult is ...
type NetworkDeviceReachabilityInfoNioListResult struct {
	Version  string                             `json:"version,omitempty"`
	Response []NetworkDeviceReachabilityInfoNio `json:"response,omitempty"`
}

// NetworkDeviceReachabilityInfoNioResult is ...
type NetworkDeviceReachabilityInfoNioResult struct {
	Version  string                           `json:"version,omitempty"`
	Response NetworkDeviceReachabilityInfoNio `json:"response,omitempty"`
}

// GetAllNetworkDevicesReachabilityInfo is ...
// Get reachability info of all devices
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceReachabilityInfoNioListResult
func (s *ReachabilityInfoService) GetAllNetworkDevicesReachabilityInfo(scope string) (*NetworkDeviceReachabilityInfoNioListResult, *Response, error) {

	path := reachabilityInfoBasePath + "/reachability-info"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceReachabilityInfoNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceReachabilityInfoByRange is ...
// Get reachability info of devices for the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceReachabilityInfoNioListResult
func (s *ReachabilityInfoService) GetNetworkDeviceReachabilityInfoByRange(startIndex int32, recordsToReturn int32, scope string) (*NetworkDeviceReachabilityInfoNioListResult, *Response, error) {

	path := reachabilityInfoBasePath + "/reachability-info/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceReachabilityInfoNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDevicesReachabilityInfoCount is ...
// Get the count of reachability-info of devices
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *ReachabilityInfoService) GetNetworkDevicesReachabilityInfoCount(scope string) (*CountResult, *Response, error) {

	path := reachabilityInfoBasePath + "/reachability-info/count"
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

// GetNetworkReachabilityInfoByID is ...
// Get reachability info of the given device by ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceReachabilityInfoNioResult
func (s *ReachabilityInfoService) GetNetworkReachabilityInfoByID(networkDeviceID string, scope string) (*NetworkDeviceReachabilityInfoNioResult, *Response, error) {

	path := reachabilityInfoBasePath + "/reachability-info/{networkDeviceID}"
	path = strings.Replace(path, "{"+"networkDeviceID"+"}", fmt.Sprintf("%v", networkDeviceID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceReachabilityInfoNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkReachabilityInfoByIPaddress is ...
// Get reachability info of the given device by IP address
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceReachabilityInfoNioResult
func (s *ReachabilityInfoService) GetNetworkReachabilityInfoByIPaddress(ipAddress string, scope string) (*NetworkDeviceReachabilityInfoNioResult, *Response, error) {

	path := reachabilityInfoBasePath + "/reachability-info/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceReachabilityInfoNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
