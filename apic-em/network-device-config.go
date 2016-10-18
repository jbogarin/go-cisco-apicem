package apicem

import (
	"fmt"
	"strings"
)

const networkDeviceConfigBasePath = "v1"

// NetworkDeviceConfigService is an interface with the Networkdevice API
type NetworkDeviceConfigService service

// RawCliInfoNio is ...
type RawCliInfoNio struct {
	Snmp            string `json:"snmp,omitempty"`            // SNMP configuration info of the device
	Inventory       string `json:"inventory,omitempty"`       // Inventory configuration info of the device
	IPIntfBrief     string `json:"ipIntfBrief,omitempty"`     // IP interface brief configuration info of the device
	IntfDescription string `json:"intfDescription,omitempty"` // Interface configuration info of the device
	MacAddressTable string `json:"macAddressTable,omitempty"` // MAC address configuration info of the device
	HealthMonitor   string `json:"healthMonitor,omitempty"`   // Health monitor configuration info of the device
	Version         string `json:"version,omitempty"`         // Version configuration info of the device
	ID              string `json:"id,omitempty"`              // Unique identifier for config
	RunningConfig   string `json:"runningConfig,omitempty"`   // Running-config of the device
	CdpNeighbors    string `json:"cdpNeighbors,omitempty"`    // CDP configuration info of the device
	AttributeInfo   string `json:"attributeInfo,omitempty"`
}

// RawCliInfoNioListResult is ...
type RawCliInfoNioListResult struct {
	Version  string          `json:"version,omitempty"`
	Response []RawCliInfoNio `json:"response,omitempty"`
}

// GetNetworkConfigCount is ...
// Gets the count of device configs
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *NetworkDeviceConfigService) GetNetworkConfigCount(scope string) (*CountResult, *Response, error) {

	path := networkDeviceConfigBasePath + "/network-device/config/count"
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

// GetRunningConfig is ...
// Gets the config for all devices
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *RawCliInfoNioListResult
func (s *NetworkDeviceConfigService) GetRunningConfig(scope string) (*RawCliInfoNioListResult, *Response, error) {

	path := networkDeviceConfigBasePath + "/network-device/config"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(RawCliInfoNioListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetRunningConfig1 is ...
// Gets the device config by device ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *SuccessResult
func (s *NetworkDeviceConfigService) GetRunningConfig1(networkDeviceID string, scope string) (*SuccessResult, *Response, error) {

	path := networkDeviceConfigBasePath + "/network-device/{networkDeviceID}/config"
	path = strings.Replace(path, "{"+"networkDeviceID"+"}", fmt.Sprintf("%v", networkDeviceID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
