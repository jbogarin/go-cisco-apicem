package apicem

import (
	"fmt"
	"strings"
)

const discoveryBasePath = "v1"

// DiscoveryService is an interface with the Discovery API
type DiscoveryService service

// DiscoveryNio is ...
type DiscoveryNio struct {
	Name                   string   `json:"name,omitempty"`                   // Name for the discovery
	SnmpRwCommunity        string   `json:"snmpRwCommunity,omitempty"`        // Snmp RW community of the devices to be discovered
	DiscoveryCondition     string   `json:"discoveryCondition,omitempty"`     // To indicate the discovery status. Available options: Complete or In Progress
	GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` // To get the list of global credential of the discovery
	IPFilterList           string   `json:"ipFilterList,omitempty"`           // IP addresses of the devices to be filtered
	IsAutoCdp              bool     `json:"isAutoCdp,omitempty"`              // Flag to mention if CDP discovery or not
	arentDiscoveryID       string   `json:"parentDiscoveryId,omitempty"`      // Parent Discovery Id from which the discovery initiated
	UserNameList           string   `json:"userNameList,omitempty"`           // Username of the devices to be discovered
	PasswordList           string   `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	EnablePasswordList     string   `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	ProtocolOrder          string   `json:"protocolOrder,omitempty"`          // Order of protocol in which device connection establishment to be tried
	CdpLevel               int32    `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	SnmpRoCommunity        string   `json:"snmpRoCommunity,omitempty"`        // Snmp RO community of the devices to be discovered
	NumDevices             int32    `json:"numDevices,omitempty"`             // Number of devices discovered in a discovery
	ID                     string   `json:"id,omitempty"`                     // Unique identifier for discovery
	RetryCount             int32    `json:"retryCount,omitempty"`             // Number of times to try establishing connection to device
	DiscoveryType          string   `json:"discoveryType,omitempty"`          // Available types are: single, auto cdp discovery, range, multi range
	TimeOut                int32    `json:"timeOut,omitempty"`                // Time to wait for device response.
	IPAddressList          string   `json:"ipAddressList,omitempty"`          // Ip address of the device to be discovered
	DiscoveryStatus        string   `json:"discoveryStatus,omitempty"`        // Available options are: active, inactive
	DeviceIDs              string   `json:"deviceIds,omitempty"`              // IDs of the devices discovered in a discovery
	AttributeInfo          string   `json:"attributeInfo,omitempty"`
}

// DiscoveryNioListResult is ...
type DiscoveryNioListResult struct {
	Version  string         `json:"version,omitempty"`
	Response []DiscoveryNio `json:"response,omitempty"`
}

// DiscoveryNioResult is ...
type DiscoveryNioResult struct {
	Version  string       `json:"version,omitempty"`
	Response DiscoveryNio `json:"response,omitempty"`
}

// InventoryRequest is ...
type InventoryRequest struct {
	SnmpROCommunity        string   `json:"snmpROCommunity,omitempty"`        // Snmp RO community of the devices to be discovered
	SnmpRWCommunity        string   `json:"snmpRWCommunity,omitempty"`        // Snmp RW community of the devices to be discovered
	SnmpUserName           string   `json:"snmpUserName,omitempty"`           // SNMP username of the device
	Name                   string   `json:"name,omitempty"`                   // Name for discovery
	EnablePasswordList     []string `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	PasswordList           []string `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	UserNameList           []string `json:"userNameList,omitempty"`           // Username of the devices to be discovered
	ProtocolOrder          string   `json:"protocolOrder,omitempty"`          // Order of protocol in which device connection establishment to be tried
	CdpLevel               int32    `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` // List of global credential ids to be used
	IPFilterList           []string `json:"ipFilterList,omitempty"`           // Username of the devices to be discovered
	SnmpAuthProtocol       string   `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. Available values:'SHA' or 'MD5'
	SnmpAuthPassphrase     string   `json:"snmpAuthPassphrase,omitempty"`     // Auth Pass phrase for SNMP
	SnmpPrivProtocol       string   `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. Available values:'DES' or 'AES128'
	SnmpPrivPassphrase     string   `json:"snmpPrivPassphrase,omitempty"`     // Pass phrase for SNMP privacy
	ReDiscovery            bool     `json:"reDiscovery,omitempty"`            // Flag to indicate is rediscovery or not
	Retry                  int32    `json:"retry,omitempty"`                  // Number of times to try establishing connection to device
	Timeout                int32    `json:"timeout,omitempty"`                // Time to wait for device response in seconds
	DiscoveryType          string   `json:"discoveryType,omitempty"`          // Available types are: single, auto cdp discovery, range, multi range
	IPAddressList          string   `json:"ipAddressList,omitempty"`          // IP address of the device to be discovered
	SnmpVersion            string   `json:"snmpVersion,omitempty"`            // Version of SNMP. Can be v2 or v3
	SnmpMode               string   `json:"snmpMode,omitempty"`               // Mode of SNMP. Available values:'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
}

// NetworkDeviceNioListResult is ...
type NetworkDeviceNioListResult struct {
	Version  string             `json:"version,omitempty"`
	Response []NetworkDeviceNio `json:"response,omitempty"`
}

// DeleteAllDiscovery is ...
// Stops all the discoveries and removes them
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *DiscoveryService) DeleteAllDiscovery(scope string) (*TaskIDResult, *Response, error) {

	path := discoveryBasePath + "/discovery"
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

// DeleteDiscoveryByID is ...
// Stops the discovery for the given ID and removes it
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *DiscoveryService) DeleteDiscoveryByID(id string, scope string) (*TaskIDResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{id}"
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

// DeleteDiscoveryByRange is ...
// Stops discovery for the given range and removes them
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *DiscoveryService) DeleteDiscoveryByRange(startIndex int32, recordsToDelete int32, scope string) (*TaskIDResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{startIndex}/{recordsToDelete}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToDelete"+"}", fmt.Sprintf("%v", recordsToDelete), -1)
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

// GetDiscoveryByID is ...
// Retrieves discovery by ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DiscoveryNioResult
func (s *DiscoveryService) GetDiscoveryByID(id string, scope string) (*DiscoveryNioResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DiscoveryNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetDiscoveryByRange is ...
// Gets the discovery for the range specified
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DiscoveryNioListResult
func (s *DiscoveryService) GetDiscoveryByRange(startIndex int32, recordsToReturn int32, scope string) (*DiscoveryNioListResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DiscoveryNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetDiscoveryCount is ...
// Gets the count of all available discovery
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *DiscoveryService) GetDiscoveryCount(scope string) (*CountResult, *Response, error) {

	path := discoveryBasePath + "/discovery/count"
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

// GetNetworkDeviceByDiscoveryID is ...
// Gets the network devices discovered for the given discovery
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceNioListResult
func (s *DiscoveryService) GetNetworkDeviceByDiscoveryID(id string, scope string) (*NetworkDeviceNioListResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{id}/network-device"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceByDiscoveryIDByRange is ...
// Gets the network devices discovered for the given discovery and for the given range. The maximum number of records that could be retrieved is 500
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceNioListResult
func (s *DiscoveryService) GetNetworkDeviceByDiscoveryIDByRange(id string, startIndex int32, recordsToReturn int32, scope string) (*NetworkDeviceNioListResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{id}/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceCountByDiscoveryID is ...
// Gets the count of network devices discovered in the given discovery
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *DiscoveryService) GetNetworkDeviceCountByDiscoveryID(id string, scope string) (*CountResult, *Response, error) {

	path := discoveryBasePath + "/discovery/{id}/network-device/count"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
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

// InsertDiscovery is ...
// Initiates discovery with the given parameters
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param request Discovery request that holds the parameters required for discovery
// * @return *TaskIDResult
func (s *DiscoveryService) InsertDiscovery(request InventoryRequest, scope string) (*TaskIDResult, *Response, error) {

	path := discoveryBasePath + "/discovery"
	req, err := s.client.NewRequest("POST", path, request)
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

// UpdateDiscovery is ...
// Stops or starts an existing discovery
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param discovery Discovery request that holds the status of discovery as active / inactive
// * @return *TaskIDResult
func (s *DiscoveryService) UpdateDiscovery(discovery DiscoveryNio, scope string) (*TaskIDResult, *Response, error) {

	path := discoveryBasePath + "/discovery"
	req, err := s.client.NewRequest("PUT", path, discovery)
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
