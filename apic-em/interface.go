package apicem

import (
	"fmt"
	"strings"
)

const interfaceBasePath = "v1"

// InterfaceService is an interface with the Interface API
type InterfaceService service

// DeviceIfDto is ...
type DeviceIfDto struct {
	Description                 string `json:"description,omitempty"`                 // interface description
	Status                      string `json:"status,omitempty"`                      // Interface status as Down / Up
	InterfaceType               string `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	MacAddress                  string `json:"macAddress,omitempty"`                  // MAC address of interface
	IfIndex                     string `json:"ifIndex,omitempty"`                     // Interface index
	Speed                       string `json:"speed,omitempty"`                       // Speed of the interface
	Duplex                      string `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	PortMode                    string `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortType                    string `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	PortName                    string `json:"portName,omitempty"`                    // Interface name
	VlanID                      string `json:"vlanId,omitempty"`                      // Vlan ID of interface
	DeviceID                    string `json:"deviceId,omitempty"`                    // ID of the device
	Series                      string `json:"series,omitempty"`                      // Series of the device
	Ipv4Address                 string `json:"ipv4Address,omitempty"`                 // IPv4 address assigned for interface
	Ipv4Mask                    string `json:"ipv4Mask,omitempty"`                    // Subnet mask for IPv4 address assigned for interface
	IsisSupport                 string `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	NativeVlanID                string `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string `json:"pid,omitempty"`                         // Platform ID of the device
	SerialNo                    string `json:"serialNo,omitempty"`                    // Serial number of the device
	InstanceUUID                string `json:"instanceUuid,omitempty"`
	ID                          string `json:"id,omitempty"`
}

// DeviceIfListResult is ...
type DeviceIfListResult struct {
	Version  string        `json:"version,omitempty"`
	Response []DeviceIfDto `json:"response,omitempty"`
}

// DeviceIfResult is ...
type DeviceIfResult struct {
	Version  string      `json:"version,omitempty"`
	Response DeviceIfDto `json:"response,omitempty"`
}

// GetInterface is ...
// Gets all available interfaces with a maximum limit of 500
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterface(scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceByDeviceID is ...
// Gets list of interfaces for the given device
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterfaceByDeviceID(deviceID string, scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceByDeviceIDAndNameQueryParams is ...
type GetInterfaceByDeviceIDAndNameQueryParams struct {
	Name string `url:"name,omitempty"` // Interface name
}

// GetInterfaceByDeviceIDAndName is ...
// Gets the interface for the given device ID and for the given interface name
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfResult
func (s *InterfaceService) GetInterfaceByDeviceIDAndName(deviceID string, scope string, queryParams *GetInterfaceByDeviceIDAndNameQueryParams) (*DeviceIfResult, *Response, error) {

	path := interfaceBasePath + "/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceByID is ...
// Gets the interface for the given interface ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfResult
func (s *InterfaceService) GetInterfaceByID(id string, scope string) (*DeviceIfResult, *Response, error) {

	path := interfaceBasePath + "/interface/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceByIsis is ...
// Gets the interfaces that has ISIS enabled
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterfaceByIsis(scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface/isis"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceByOspf is ...
// Gets the interfaces that has OSPF enabled
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterfaceByOspf(scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface/ospf"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceCount is ...
// Gets the count of interfaces for all devices
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *InterfaceService) GetInterfaceCount(scope string) (*CountResult, *Response, error) {

	path := interfaceBasePath + "/interface/count"
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

// GetInterfaceCountByDevice is ...
// Gets the interface count for the given device
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *InterfaceService) GetInterfaceCountByDevice(deviceID string, scope string) (*CountResult, *Response, error) {

	path := interfaceBasePath + "/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
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

// GetInterfaceListByIP is ...
// Gets list of interfaces with the given IP address
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterfaceListByIP(ipAddress string, scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetInterfaceRangeByDevice is ...
// Gets list of interfaces for the device and for the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *DeviceIfListResult
func (s *InterfaceService) GetInterfaceRangeByDevice(deviceID string, startIndex int32, recordsToReturn int32, scope string) (*DeviceIfListResult, *Response, error) {

	path := interfaceBasePath + "/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(DeviceIfListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
