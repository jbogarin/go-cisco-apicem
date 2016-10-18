package apicem

import (
	"fmt"
	"strings"
)

const hostBasePath = "v1"

// HostService is an interface with the Host API
type HostService service

// HostDto is ...
type HostDto struct {
	HostName                        string `json:"hostName,omitempty"`                        // Name of the host
	Source                          string `json:"source,omitempty"`                          // Source from which the host gets collected. Available option:200 for inventory collection and 300 for trap based data collection
	LastUpdated                     string `json:"lastUpdated,omitempty"`                     // Time when the host info last got updated
	VlanID                          string `json:"vlanId,omitempty"`                          // Vlan Id of the host
	AvgUpdateFrequency              string `json:"avgUpdateFrequency,omitempty"`              // Frequency in which host info gets updated
	HostMac                         string `json:"hostMac,omitempty"`                         // Mac address of the host
	HostType                        string `json:"hostType,omitempty"`                        // Type of the host. Available options are: Wired, Wireless
	PointOfAttachment               string `json:"pointOfAttachment,omitempty"`               // Id of the Host's Point of attachment network device (wlc). Based on mobility
	PointOfPresence                 string `json:"pointOfPresence,omitempty"`                 // Id of the Host's Point of presence network device (wlc). Based on mobility
	HostIP                          string `json:"hostIp,omitempty"`                          // Ip address of the host
	ConnectedAPMacAddress           string `json:"connectedAPMacAddress,omitempty"`           // Mac address of the AP to which wireless host gets connected
	ConnectedAPName                 string `json:"connectedAPName,omitempty"`                 // Name of the AP to which wireless host gets connected
	ConnectedInterfaceID            string `json:"connectedInterfaceId,omitempty"`            // Id of the interface to which host gets connected
	ConnectedInterfaceName          string `json:"connectedInterfaceName,omitempty"`          // Name of the interface to which host gets connected
	ConnectedNetworkDeviceID        string `json:"connectedNetworkDeviceId,omitempty"`        // Id of the network device to which host gets connected
	ConnectedNetworkDeviceIPAddress string `json:"connectedNetworkDeviceIpAddress,omitempty"` // Ip address of the network device to which host gets connected
	AttributeInfo                   string `json:"attributeInfo,omitempty"`
	ID                              string `json:"id,omitempty"`
}

// HostListResult is ...
type HostListResult struct {
	Version  string    `json:"version,omitempty"`
	Response []HostDto `json:"response,omitempty"`
}

// HostResult is ...
type HostResult struct {
	Version  string  `json:"version,omitempty"`
	Response HostDto `json:"response,omitempty"`
}

// GetHostByID is ...
// Get host by id
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *HostResult
func (s *HostService) GetHostByID(id string, scope string) (*HostResult, *Response, error) {

	path := hostBasePath + "/host/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(HostResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetHostCount is ...
// Get host Count
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *HostService) GetHostCount(scope string) (*CountResult, *Response, error) {

	path := hostBasePath + "/host/count"
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

// GetHostsQueryParams is ...
type GetHostsQueryParams struct {
	Limit                  string   `url:"limit,omitempty"`                  // limit
	Offset                 string   `url:"offset,omitempty"`                 // offset
	SortBy                 string   `url:"sortBy,omitempty"`                 // sortBy
	Order                  string   `url:"order,omitempty"`                  // order
	HostName               []string `url:"hostName,omitempty"`               // hostName
	HostMac                []string `url:"hostMac,omitempty"`                // hostMac
	HostType               []string `url:"hostType,omitempty"`               // hostType
	ConnectedInterfaceName []string `url:"connectedInterfaceName,omitempty"` // connectedInterfaceName
	HostIP                 []string `url:"hostIp,omitempty"`                 // hostIp
	ConnectedDeviceIP      []string `url:"connectedDeviceIp,omitempty"`      // connectedDeviceIp
}

// GetHosts is ...
// Get Hosts
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *HostListResult
func (s *HostService) GetHosts(scope string, queryParams *GetHostsQueryParams) (*HostListResult, *Response, error) {

	path := hostBasePath + "/host"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(HostListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
