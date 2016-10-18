package apicem

import "time"

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

// SuccessResult is ...
type SuccessResult struct {
	Version  string `json:"version,omitempty"`
	Response string `json:"response,omitempty"`
}

// NetworkDeviceNio is ...
type NetworkDeviceNio struct {
	AnchorWlcForAp            string    `json:"anchorWlcForAp,omitempty"`            // Connected WLC device for AP
	AuthModelID               string    `json:"authModelId,omitempty"`               // Authentication model Id on device
	AvgUpdateFrequency        int32     `json:"avgUpdateFrequency,omitempty"`        // Frequency in which interface info gets updated
	BootDateTime              time.Time `json:"bootDateTime,omitempty"`              // Device boot time
	DuplicateDeviceID         string    `json:"duplicateDeviceId,omitempty"`         // Identifier of the duplicate ip of the same device discovered
	Family                    string    `json:"family,omitempty"`                    // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                  string    `json:"hostname,omitempty"`                  // Device name
	ID                        string    `json:"id,omitempty"`                        // Unique identifier of network device
	ImageName                 string    `json:"imageName,omitempty"`                 // Image details on the device
	IngressQueueConfig        string    `json:"ingressQueueConfig,omitempty"`        // Ingress queue config on device
	InterfaceCount            string    `json:"interfaceCount,omitempty"`            // Number of interfaces on the device
	LastUpdated               string    `json:"lastUpdated,omitempty"`               // Time when the network device info last got updated
	LineCardCount             string    `json:"lineCardCount,omitempty"`             // Number of linecards on the device
	LineCardID                string    `json:"lineCardId,omitempty"`                // IDs of linecards of the device
	Location                  string    `json:"location,omitempty"`                  // Location ID that is associated with the device
	LocationName              string    `json:"locationName,omitempty"`              // Name of the associated location
	MacAddress                string    `json:"macAddress,omitempty"`                // MAC address of device
	ManagementIPAddress       string    `json:"managementIpAddress,omitempty"`       // IP address of the device
	MemorySize                string    `json:"memorySize,omitempty"`                // Processor memory size
	NumUpdates                int32     `json:"numUpdates,omitempty"`                // Number of time network-device info got updated
	PlatformID                string    `json:"platformId,omitempty"`                // Platform ID of device
	PortRange                 string    `json:"portRange,omitempty"`                 // Range of ports on device
	QosStatus                 string    `json:"qosStatus,omitempty"`                 // Qos status on device
	ReachabilityFailureReason string    `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices
	ReachabilityStatus        string    `json:"reachabilityStatus,omitempty"`        // Device reachability status as Reachable / Unreachable
	Role                      string    `json:"role,omitempty"`                      // Role of device as access, distribution, border router, core
	RoleSource                string    `json:"roleSource,omitempty"`                // Role source as manual / auto
	SerialNumber              string    `json:"serialNumber,omitempty"`              // Serial number of device
	SnmpContact               string    `json:"snmpContact,omitempty"`               // SNMP contact on device
	SnmpLocation              string    `json:"snmpLocation,omitempty"`              // SNMP location on device
	SoftwareVersion           string    `json:"softwareVersion,omitempty"`           // Software version on the device
	Tag                       string    `json:"tag,omitempty"`                       // Tag ID that is associated with the device
	TagCount                  int32     `json:"tagCount,omitempty"`                  // Number of tags associated with the device
	Type                      string    `json:"type,omitempty"`                      // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                    string    `json:"upTime,omitempty"`                    // Time that shows for how long the device has been up
	Vendor                    string    `json:"vendor,omitempty"`                    // Vendor information of the device
	WlcApDeviceStatus         string    `json:"wlcApDeviceStatus,omitempty"`         // Collection status of AP devices
}

// Void is ...
type Void struct {
}

// CategoryDto is ...
type CategoryDto struct {
	ID   string `json:"id,omitempty"`   // id
	Name string `json:"name,omitempty"` // Category Name
}

// LinkWrapper is ...
type LinkWrapper struct {
	Target               string `json:"target,omitempty"`               // Device ID corresponding to the target device
	Source               string `json:"source,omitempty"`               // Device ID correspondng to the source device
	Tag                  string `json:"tag,omitempty"`                  // Tag for the devices
	Id                   string `json:"id,omitempty"`                   // Unified identifier for device
	StartPortID          string `json:"startPortID,omitempty"`          // Device port ID corresponding to start devices
	EndPortID            string `json:"endPortID,omitempty"`            // Device port ID corresponding to end devices
	LinkStatus           string `json:"linkStatus,omitempty"`           // Indicates whether link is working
	EndPortName          string `json:"endPortName,omitempty"`          // Interface port name corresponding to end devices
	EndPortSpeed         string `json:"endPortSpeed,omitempty"`         // Interface port speed corresponding to end devices
	StartPortName        string `json:"startPortName,omitempty"`        // Interface port name corresponding to start devices
	StartPortSpeed       string `json:"startPortSpeed,omitempty"`       // Interface port speed corresponding to start devices
	GreyOut              bool   `json:"greyOut,omitempty"`              // Device greyout
	EndPortIpv4Address   string `json:"endPortIpv4Address,omitempty"`   // Interface port IPv4 address corresponding to end devices
	EndPortIpv4Mask      string `json:"endPortIpv4Mask,omitempty"`      // Interface port IPv4 mask corresponding to end devices
	StartPortIpv4Address string `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start devices
	StartPortIpv4Mask    string `json:"startPortIpv4Mask,omitempty"`    // Interface port IPv4 mask corresponding to start devices
	AttributeInfo        string `json:"attributeInfo,omitempty"`
}

// NodeWrapperCustom is ...
type NodeWrapperCustom struct {
	Y             int32  `json:"y,omitempty"`            // Y - Coordinate for this Node in the topology View
	X             int32  `json:"x,omitempty"`            // X - Coordinate for this Node in the topology View
	Label         string `json:"label,omitempty"`        // Label for the node
	ParentNodeId  string `json:"parentNodeId,omitempty"` // Unique Id of the Node for ehich the custom properties are being represented
	AttributeInfo string `json:"attributeInfo,omitempty"`
	Id            string `json:"id,omitempty"`
}

// NodeWrapper is ...
type NodeWrapper struct {
	Tags            []string          `json:"tags,omitempty"`            // List of tags applied on this device
	Y               int32             `json:"y,omitempty"`               // Y position of the device on the displayed topology view
	X               int32             `json:"x,omitempty"`               // X position of the device on the displayed topology view
	Fixed           bool              `json:"fixed,omitempty"`           // Indication of whether the position is fixed or will use auto layout
	Family          string            `json:"family,omitempty"`          // Product family which is part of the product identifier
	NodeType        string            `json:"nodeType,omitempty"`        // Host or device
	Role            string            `json:"role,omitempty"`            // Role of the device
	Order           int32             `json:"order,omitempty"`           // Device order by link number
	Id              string            `json:"id,omitempty"`              // Unique identifier for device
	Label           string            `json:"label,omitempty"`           // Hostname of the device
	CustomParam     NodeWrapperCustom `json:"customParam,omitempty"`     // Device custom parameters
	PlatformId      string            `json:"platformId,omitempty"`      // Device platform description
	DeviceType      string            `json:"deviceType,omitempty"`      // Type of the device
	VlanId          string            `json:"vlanId,omitempty"`          // VLan id
	RoleSource      string            `json:"roleSource,omitempty"`      // Indicates whether role is assigned manually or automatically
	UserId          string            `json:"userId,omitempty"`          // ID of the host
	SoftwareVersion string            `json:"softwareVersion,omitempty"` // Device OS version
	NetworkType     string            `json:"networkType,omitempty"`     // Type of network
	Ip              string            `json:"ip,omitempty"`              // IP address of the device
	UpperNode       string            `json:"upperNode,omitempty"`       // Start node ID
	AclApplied      bool              `json:"aclApplied,omitempty"`      // Indicates if the ACL that is applied on the device
	GreyOut         bool              `json:"greyOut,omitempty"`         // Indicates if the device is active for this topology view
	OsType          string            `json:"osType,omitempty"`          // OS type of the device
	DataPathId      string            `json:"dataPathId,omitempty"`      // ID of the path between devices
	AttributeInfo   string            `json:"attributeInfo,omitempty"`
}

// Topology is ...
type Topology struct {
	Nodes         []NodeWrapper `json:"nodes,omitempty"` // List of devices and hosts
	Links         []LinkWrapper `json:"links,omitempty"` // List of link between devices
	AttributeInfo string        `json:"attributeInfo,omitempty"`
	Id            string        `json:"id,omitempty"`
}

// ScalableGroupDto is ...
type ScalableGroupDto struct {
	Description                 string `json:"description,omitempty"`                 // description
	IdentitySourceIPAddress     string `json:"identitySourceIpAddress,omitempty"`     // identitySourceIpAddress
	CreateTime                  int64  `json:"createTime,omitempty"`                  // createTime
	ScalableGroupExternalHandle string `json:"scalableGroupExternalHandle,omitempty"` // scalableGroupExternalHandle
	LastUpdateTime              int64  `json:"lastUpdateTime,omitempty"`              // lastUpdateTime
	IdentitySourceID            string `json:"identitySourceId,omitempty"`            // identitySourceId
	IdentitySourceType          string `json:"identitySourceType,omitempty"`          // identitySourceType
	Name                        string `json:"name,omitempty"`                        // name
	ID                          string `json:"id,omitempty"`                          // id
}
