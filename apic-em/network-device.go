package apicem

import (
	"fmt"
	"strings"
	"time"
)

const networkDeviceBasePath = "v1"

// NetworkDeviceService is an interface with the Networkdevice API
type NetworkDeviceService service

// CountResult is ...
type CountResult struct {
	Version  string `json:"version,omitempty"`
	Response int32  `json:"response,omitempty"`
}

// Date is ...
type Date struct {
	Time    int64 `json:"time,omitempty"`
	Hours   int32 `json:"hours,omitempty"`
	Minutes int32 `json:"minutes,omitempty"`
	Seconds int32 `json:"seconds,omitempty"`
	Year    int32 `json:"year,omitempty"`
	Month   int32 `json:"month,omitempty"`
	Date    int32 `json:"date,omitempty"`
}

// NetworkDeviceBriefNio is ...
type NetworkDeviceBriefNio struct {
	// Unique identifier of the network device
	ID string `json:"id,omitempty"`
	// Role of device as access, distribution, border router, core
	Role string `json:"role,omitempty"`
	// Role source as manual / auto
	RoleSource string `json:"roleSource,omitempty"`
}

// NetworkDeviceBriefNioResult is ...
type NetworkDeviceBriefNioResult struct {
	Version  string                `json:"version,omitempty"`
	Response NetworkDeviceBriefNio `json:"response,omitempty"`
}

// NetworkDeviceListResult is ...
type NetworkDeviceListResult struct {
	Version  string              `json:"version,omitempty"`
	Response []*NetworkDeviceDto `json:"response,omitempty"`
}

// NetworkDeviceManagementInfo is ...
type NetworkDeviceManagementInfo struct {
	// Family of device as switch, router, wireless lan controller, accesspoints
	Family string `json:"family,omitempty"`
	// Unique identifier of device
	ID string `json:"id,omitempty"`
	// Type of device as switch, router, wireless lan controller, accesspoints
	Type string `json:"type,omitempty"`
	// Device name
	Hostname string `json:"hostname,omitempty"`
	// Device series
	Series string `json:"series,omitempty"`
	// IP address of the device
	ManagementIPAddress string `json:"managementIpAddress,omitempty"`
	// Credential info
	Credentials string `json:"credentials,omitempty"`
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

// NetworkDeviceDto is ...
type NetworkDeviceDto struct {
	// Family of device as switch, router, wireless lan controller, accesspoints
	Family string `json:"family,omitempty"`
	// Location ID that is associated with the device
	Location string `json:"location,omitempty"`
	// Type of device as switch, router, wireless lan controller, accesspoints
	Type string `json:"type,omitempty"`
	// Serial number of device
	SerialNumber string `json:"serialNumber,omitempty"`
	// Role of device as access, distribution, border router, core
	Role string `json:"role,omitempty"`
	// MAC address of device
	MacAddress string `json:"macAddress,omitempty"`
	// Time that shows for how long the device has been up
	UpTime string `json:"upTime,omitempty"`
	// Software version on the device
	SoftwareVersion string `json:"softwareVersion,omitempty"`
	// Status detail of inventory sync
	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"`
	LastUpdateTime        int    `json:"lastUpdateTime,omitempty"`
	// Name of the associated location
	LocationName string `json:"locationName,omitempty"`
	// Time when the network device info last got updated
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Number of tags associated with the device
	TagCount string `json:"tagCount,omitempty"`
	// Device name
	Hostname string `json:"hostname,omitempty"`
	// Failure reason for unreachable devices
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"`
	// Device reachability status as Reachable / Unreachable
	ReachabilityStatus string `json:"reachabilityStatus,omitempty"`
	// Role source as manual / auto
	RoleSource string `json:"roleSource,omitempty"`
	// Device series
	Series string `json:"series,omitempty"`
	// SNMP contact on device
	SnmpContact string `json:"snmpContact,omitempty"`
	// SNMP location on device
	SnmpLocation string `json:"snmpLocation,omitempty"`
	// Mobility protocol port is stored as tunneludpport for WLC
	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"`
	// IP address of WLC on AP manager interface
	APManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"`
	// Device boot time
	BootDateTime string `json:"bootDateTime,omitempty"`
	// Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	CollectionStatus string `json:"collectionStatus,omitempty"`
	// Number of interfaces on the device
	InterfaceCount string `json:"interfaceCount,omitempty"`
	// Number of linecards on the device
	LineCardCount string `json:"lineCardCount,omitempty"`
	// IDs of linecards of the device
	LineCardID string `json:"lineCardId,omitempty"`
	// IP address of the device
	ManagementIPAddress string `json:"managementIpAddress,omitempty"`
	// Processor memory size
	MemorySize string `json:"memorySize,omitempty"`
	// Platform ID of device
	PlatformID   string `json:"platformId,omitempty"`
	InstanceUUID string `json:"instanceUuid,omitempty"`
	ID           string `json:"id,omitempty"`
}

// NetworkDeviceNio is ...
type NetworkDeviceNio struct {
	// Ingress queue config on device
	IngressQueueConfig string `json:"ingressQueueConfig,omitempty"`
	// Authentication model Id on device
	AuthModelID string `json:"authModelId,omitempty"`
	// Identifier of the duplicate ip of the same device discovered
	DuplicateDeviceID string `json:"duplicateDeviceId,omitempty"`
	// Connected WLC device for AP
	AnchorWlcForAp string `json:"anchorWlcForAp,omitempty"`
	// Collection status of AP devices
	WlcApDeviceStatus string `json:"wlcApDeviceStatus,omitempty"`
	// Type of device as switch, router, wireless lan controller, accesspoints
	Type string `json:"type,omitempty"`
	// Location ID that is associated with the device
	Location string `json:"location,omitempty"`
	// Serial number of device
	SerialNumber string `json:"serialNumber,omitempty"`
	// Role of device as access, distribution, border router, core
	Role string `json:"role,omitempty"`
	// MAC address of device
	MacAddress string `json:"macAddress,omitempty"`
	// Time that shows for how long the device has been up
	UpTime string `json:"upTime,omitempty"`
	// Software version on the device
	SoftwareVersion string `json:"softwareVersion,omitempty"`
	// Name of the associated location
	LocationName string `json:"locationName,omitempty"`
	// Unique identifier of network device
	ID string `json:"id,omitempty"`
	// Image details on the device
	ImageName string `json:"imageName,omitempty"`
	// Time when the network device info last got updated
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Vendor information of the device
	Vendor string `json:"vendor,omitempty"`
	// Range of ports on device
	PortRange string `json:"portRange,omitempty"`
	// Number of tags associated with the device
	TagCount int32 `json:"tagCount,omitempty"`
	// Device name
	Hostname string `json:"hostname,omitempty"`
	// Tag ID that is associated with the device
	Tag string `json:"tag,omitempty"`
	// Failure reason for unreachable devices
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"`
	// Device reachability status as Reachable / Unreachable
	ReachabilityStatus string `json:"reachabilityStatus,omitempty"`
	// Role source as manual / auto
	RoleSource string `json:"roleSource,omitempty"`
	// SNMP contact on device
	SnmpContact string `json:"snmpContact,omitempty"`
	// SNMP location on device
	SnmpLocation string `json:"snmpLocation,omitempty"`
	// Frequency in which interface info gets updated
	AvgUpdateFrequency int32 `json:"avgUpdateFrequency,omitempty"`
	// Number of time network-device info got updated
	NumUpdates int32 `json:"numUpdates,omitempty"`
	// Device boot time
	BootDateTime time.Time `json:"bootDateTime,omitempty"`
	// Family of device as switch, router, wireless lan controller, accesspoints
	Family string `json:"family,omitempty"`
	// Number of interfaces on the device
	InterfaceCount string `json:"interfaceCount,omitempty"`
	// Number of linecards on the device
	LineCardCount string `json:"lineCardCount,omitempty"`
	// IDs of linecards of the device
	LineCardID string `json:"lineCardId,omitempty"`
	// IP address of the device
	ManagementIPAddress string `json:"managementIpAddress,omitempty"`
	// Processor memory size
	MemorySize string `json:"memorySize,omitempty"`
	// Platform ID of device
	PlatformID string `json:"platformId,omitempty"`
	// Qos status on device
	QosStatus string `json:"qosStatus,omitempty"`
}

// NetworkDeviceResult is ...
type NetworkDeviceResult struct {
	Version  string           `json:"version,omitempty"`
	Response NetworkDeviceDto `json:"response,omitempty"`
}

// SiteManagementInfo is ...
type SiteManagementInfo struct {
	// Description of site
	Description string `json:"description,omitempty"`
	// Name of site
	Name string `json:"name,omitempty"`
	// Location of site
	Location string `json:"location,omitempty"`
	// Properties of site
	Properties string `json:"properties,omitempty"`
	// Unique identifier of site
	ID string `json:"id,omitempty"`
	// Unique identifier of devices that are associated with site
	DeviceIds []string `json:"deviceIds,omitempty"`
}

// TaskIDResponse is ...
type TaskIDResponse struct {
	URL    string `json:"url,omitempty"`
	TaskID TaskID `json:"taskId,omitempty"`
}

// TaskID is ...
type TaskID struct {
}

// TaskIDResult is ...
type TaskIDResult struct {
	Version  string         `json:"version,omitempty"`
	Response TaskIDResponse `json:"response,omitempty"`
}

// AddNetworkDeviceLocation is ...
// Associates the given location to the given device
//
//  * @param scope Authorization Scope for RBAC
//  * @param networkDeviceNIO networkDeviceNIO
//  * @return *TaskIDResult
func (s *NetworkDeviceService) AddNetworkDeviceLocation(scope string, networkDeviceNIO *NetworkDeviceNio) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location"

	req, err := s.client.NewRequest("Post", path, networkDeviceNIO)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param id Device ID
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *NetworkDeviceService) DeleteDevicebyID(id *string, scope string) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Delete", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param id Device ID
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *NetworkDeviceService) DeleteNetworkLocationByID(id *string, scope string) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/location"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Delete", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @return *Object
func (s *NetworkDeviceService) GetAllNetworkDevice(scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param id Device ID
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceBriefNioResult
func (s *NetworkDeviceService) GetNetworkDeviceBrief(id *string, scope string) (*NetworkDeviceBriefNioResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/brief"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param id Device ID
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceByID(id *string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param ipAddress Device IP address
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceByIP(ipAddress *string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param locationID Location ID
//  * @param startIndex Start index
//  * @param recordsToReturn Number of records to return
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceByLocationByRange(locationID *string, startIndex *int32, recordsToReturn *int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{locationID}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"locationID"+"}", fmt.Sprintf("%v", locationID), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param locationID Location ID
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceByLocationID(locationID *string, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{locationID}"
	path = strings.Replace(path, "{"+"locationID"+"}", fmt.Sprintf("%v", locationID), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param serialNumber Device serial number
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceBySerialNumber(serialNumber *string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @return *CountResult
func (s *NetworkDeviceService) GetNetworkDeviceCount(scope string) (*CountResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/count"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceLocation(scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param id Device ID
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceResult
func (s *NetworkDeviceService) GetNetworkDeviceLocationByID(id *string, scope string) (*NetworkDeviceResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{id}/location"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param startIndex startIndex
//  * @param recordsToReturn recordsToReturn
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceLocationByRange(startIndex *int32, recordsToReturn *int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/location/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param startIndex Start index
//  * @param recordsToReturn Number of records to return
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkDeviceListResult
func (s *NetworkDeviceService) GetNetworkDeviceRange(startIndex *int32, recordsToReturn *int32, scope string) (*NetworkDeviceListResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @return *NetworkManagementInfoResult
func (s *NetworkDeviceService) GetNetworkManagementInfo(scope string) (*NetworkManagementInfoResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/management-info"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @return *CountResult
func (s *NetworkDeviceService) GetNetworkManagementInfoCount(scope string) (*CountResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/management-info/count"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param scope Authorization Scope for RBAC
//  * @param networkDeviceBriefNIO networkDeviceBriefNIO
//  * @return *TaskIDResult
func (s *NetworkDeviceService) UpdateNetworkDevice(scope string, networkDeviceBriefNIO *NetworkDeviceBriefNio) (*TaskIDResult, *Response, error) {

	path := networkDeviceBasePath + "/network-device/brief"

	req, err := s.client.NewRequest("Put", path, networkDeviceBriefNIO)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
