package apicem

import (
	"fmt"
	"strings"
)

const vlanBasePath = "v1"

// VLANService is an interface with the Networkdeviceidvlan API
type VLANService service

// VLANDto is ...
type VLANDto struct {
	NumberOfIPs    int32  `json:"numberOfIPs,omitempty"`
	Mask           int32  `json:"mask,omitempty"`
	Prefix         string `json:"prefix,omitempty"`
	InterfaceName  string `json:"interfaceName,omitempty"`
	IPAddress      string `json:"ipAddress,omitempty"`
	NetworkAddress string `json:"networkAddress,omitempty"`
	VLANType       string `json:"vlanType,omitempty"`
	VLANNumber     int32  `json:"vlanNumber,omitempty"`
}

// VLANListResult is ...
type VLANListResult struct {
	Version  string    `json:"version,omitempty"`
	Response []VLANDto `json:"response,omitempty"`
}

// GetDeviceVLANDataQueryParams is ...
type GetDeviceVLANDataQueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` // VLAN assocaited with sub-interface
}

// GetDeviceVLANData is ...
// getDeviceVLANData
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *VLANListResult
func (s *VLANService) GetDeviceVLANData(scope string, id string, queryParams *GetDeviceVLANDataQueryParams) (*VLANListResult, *Response, error) {

	path := vlanBasePath + "/network-device/{id}/vlan"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(VLANListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
