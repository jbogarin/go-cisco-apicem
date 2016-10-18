package apicem

import (
	"fmt"
	"strings"
)

const networkDeviceBasePath = "v1"

// NetworkDeviceService is an interface with the Networkdevice API
type NetworkDeviceService service

// NetworkDeviceBriefNio is ...
type NetworkDeviceBriefNio struct {
	ID         string `json:"id,omitempty"`         // Unique identifier of the network device
	Role       string `json:"role,omitempty"`       // Role of device as access, distribution, border router, core
	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto
}

// NetworkDeviceBriefNioResult is ...
type NetworkDeviceBriefNioResult struct {
	Version  string                `json:"version,omitempty"`
	Response NetworkDeviceBriefNio `json:"response,omitempty"`
}

// NetworkDeviceDto is ...
type NetworkDeviceDto struct {
	Family                    string `json:"family,omitempty"`                // Family of device as switch, router, wireless lan controller, accesspoints
	Location                  string `json:"location,omitempty"`              // Location ID that is associated with the device
	Type                      string `json:"type,omitempty"`                  // Type of device as switch, router, wireless lan controller, accesspoints
	SerialNumber              string `json:"serialNumber,omitempty"`          // Serial number of device
	Role                      string `json:"role,omitempty"`                  // Role of device as access, distribution, border router, core
	MacAddress                string `json:"macAddress,omitempty"`            // MAC address of device
	UpTime                    string `json:"upTime,omitempty"`                // Time that shows for how long the device has been up
	SoftwareVersion           string `json:"softwareVersion,omitempty"`       // Software version on the device
	InventoryStatusDetail     string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync
	LastUpdateTime            int    `json:"lastUpdateTime,omitempty"`
	LocationName              string `json:"locationName,omitempty"`              // Name of the associated location
	LastUpdated               string `json:"lastUpdated,omitempty"`               // Time when the network device info last got updated
	TagCount                  string `json:"tagCount,omitempty"`                  // Number of tags associated with the device
	Hostname                  string `json:"hostname,omitempty"`                  // Device name
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        // Device reachability status as Reachable / Unreachable
	RoleSource                string `json:"roleSource,omitempty"`                // Role source as manual / auto
	Series                    string `json:"series,omitempty"`                    // Device series
	SnmpContact               string `json:"snmpContact,omitempty"`               // SNMP contact on device
	SnmpLocation              string `json:"snmpLocation,omitempty"`              // SNMP location on device
	TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             // Mobility protocol port is stored as tunneludpport for WLC
	ApManagerInterfaceIP      string `json:"apManagerInterfaceIp,omitempty"`      // IP address of WLC on AP manager interface
	BootDateTime              string `json:"bootDateTime,omitempty"`              // Device boot time
	CollectionStatus          string `json:"collectionStatus,omitempty"`          // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	InterfaceCount            string `json:"interfaceCount,omitempty"`            // Number of interfaces on the device
	LineCardCount             string `json:"lineCardCount,omitempty"`             // Number of linecards on the device
	LineCardID                string `json:"lineCardId,omitempty"`                // IDs of linecards of the device
	ManagementIPAddress       string `json:"managementIpAddress,omitempty"`       // IP address of the device
	MemorySize                string `json:"memorySize,omitempty"`                // Processor memory size
	PlatformID                string `json:"platformId,omitempty"`                // Platform ID of device
	InstanceUUID              string `json:"instanceUuid,omitempty"`
	ID                        string `json:"id,omitempty"`
}

// NetworkDeviceListResult is ...
type NetworkDeviceListResult struct {
	Version  string             `json:"version,omitempty"`
	Response []NetworkDeviceDto `json:"response,omitempty"`
}

// NetworkDeviceManagementInfo is ...
type NetworkDeviceManagementInfo struct {
	Family              string `json:"family,omitempty"`              // Family of device as switch, router, wireless lan controller, accesspoints
	ID                  string `json:"id,omitempty"`                  // Unique identifier of device
	Type                string `json:"type,omitempty"`                // Type of device as switch, router, wireless lan controller, accesspoints
	Hostname            string `json:"hostname,omitempty"`            // Device name
	Series              string `json:"series,omitempty"`              // Device series
	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device
	Credentials         string `json:"credentials,omitempty"`         // Credential info
}

// NetworkDeviceResult is ...
type NetworkDeviceResult struct {
	Version  string           `json:"version,omitempty"`
	Response NetworkDeviceDto `json:"response,omitempty"`
}

// NetworkManagementInfo is ...
type NetworkManagementInfo struct {
	NetworkDeviceManagementInfo []NetworkDeviceManagementInfo `json:"networkDeviceManagementInfo,omitempty"`
	SiteManagementInfo          []SiteManagementInfo          `json:"siteManagementInfo,omitempty"`
}

// NetworkManagementInfoResult is ...
type NetworkManagementInfoResult struct {
	Version  string                `json:"version,omitempty"`
	Response NetworkManagementInfo `json:"response,omitempty"`
}

// SiteManagementInfo is ...
type SiteManagementInfo struct {
	Description string   `json:"description,omitempty"` // Description of site
	Name        string   `json:"name,omitempty"`        // Name of site
	Location    string   `json:"location,omitempty"`    // Location of site
	Properties  string   `json:"properties,omitempty"`  // Properties of site
	ID          string   `json:"id,omitempty"`          // Unique identifier of site
	DeviceIDs   []string `json:"deviceIds,omitempty"`   // Unique identifier of devices that are associated with site
}

// AddNetworkDeviceLocation is ...
// Associates the given location to the given device
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param networkDeviceNIO networkDeviceNIO
// * @return *TaskIDResult
func (s *NetworkDeviceService) AddNetworkDeviceLocation(scope string, networkDeviceNIO NetworkDeviceNio) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location"
	req, err := s.client.NewRequest("POST", path, networkDeviceNIO)
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

// DeleteDevicebyID is ...
// Removes the network device for the given ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *NetworkDeviceService) DeleteDevicebyID(id string, scope string) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}"
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

// DeleteNetworkLocationByID is ...
// Removes the association between device and location for the given device
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *NetworkDeviceService) DeleteNetworkLocationByID(id string, scope string) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/location"
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

// GetAllNetworkDevice is ...
// Gets the list of network devices filtered using management IP address, mac address, hostname and location name
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *string
func (s *NetworkDeviceService) GetAllNetworkDevice(scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceBrief is ...
// Gets brief network device info such as hostname, management IP address for the given device ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceBriefNioResult
func (s *NetworkDeviceService) GetNetworkDeviceBrief(id string, scope string) (*NetworkDeviceBriefNioResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/brief"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceBriefNioResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceByID is ...
// Gets the network device for the given device ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceByID(id string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceByIP is ...
// Gets the network device with the given IP address
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceByIP(ipAddress string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceByLocationByRange is ...
// Gets network devices associated with the given location in the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceByLocationByRange(locationID string, startIndex int32, recordsToReturn int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{locationID}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"locationID"+"}", fmt.Sprintf("%v", locationID), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceByLocationID is ...
// Gets list of network devices that are associated with the given location
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceByLocationID(locationID string, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{locationID}"
	path = strings.Replace(path, "{"+"locationID"+"}", fmt.Sprintf("%v", locationID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceBySerialNumber is ...
// Gets the network device with the given serial number
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceBySerialNumber(serialNumber string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceCount is ...
// Gets the count of network devices filtered by management IP address, mac address, hostname and location name
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *NetworkDeviceService) GetNetworkDeviceCount(scope string) (*CountResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/count"
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

// GetNetworkDeviceLocation is ...
// Gets the list of network devices that has a location
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceLocation(scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceLocationByID is ...
// Gets the location for the given device ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceLocationByID(id string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/location"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceLocationByRange is ...
// Gets the location for the devices in the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceLocationByRange(startIndex int32, recordsToReturn int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkDeviceRange is ...
// Gets the list of network devices for the given range
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceRange(startIndex int32, recordsToReturn int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkDeviceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkManagementInfo is ...
// Gets the managment information of network devices and sites. Maximum allowed limit is 100.
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *NetworkManagementInfoResult
func (s *NetworkDeviceService) GetNetworkManagementInfo(scope string) (*NetworkManagementInfoResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/management-info"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(NetworkManagementInfoResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetNetworkManagementInfoCount is ...
// Gets the number of network devices or sites, whichever is maximum. This count is used to paginate and query the /network-device/management-info API
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *NetworkDeviceService) GetNetworkManagementInfoCount(scope string) (*CountResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/management-info/count"
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

// UpdateNetworkDevice is ...
// Updates the role of the device as access, core, distribution, border router
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param networkDeviceBriefNIO networkDeviceBriefNIO
// * @return *TaskIDResult
func (s *NetworkDeviceService) UpdateNetworkDevice(scope string, networkDeviceBriefNIO NetworkDeviceBriefNio) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/brief"
	req, err := s.client.NewRequest("PUT", path, networkDeviceBriefNIO)
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
