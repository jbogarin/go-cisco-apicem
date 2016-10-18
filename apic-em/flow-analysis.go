package apicem

import (
	"fmt"
	"strings"
)

const flowAnalysisBasePath = "v1"

// FlowAnalysisService is an interface with the FlowAnalysis API
type FlowAnalysisService service

// Accuracy is ...
type Accuracy struct {
	Percent int32  `json:"percent,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// AclAce is ...
type AclAce struct {
	Result        string             `json:"result,omitempty"`
	Ace           string             `json:"ace,omitempty"`
	MatchingPorts []AclMatchingPorts `json:"matchingPorts,omitempty"`
}

// AclAnalysisResponse is ...
type AclAnalysisResponse struct {
	Result       string   `json:"result,omitempty"`
	AclName      string   `json:"aclName,omitempty"`
	MatchingAces []AclAce `json:"matchingAces,omitempty"`
}

// AclMatchingPorts is ...
type AclMatchingPorts struct {
	Protocol string        `json:"protocol,omitempty"`
	Ports    []AclPortInfo `json:"ports,omitempty"`
}

// AclPortInfo is ...
type AclPortInfo struct {
	SourcePorts []string `json:"sourcePorts,omitempty"`
	DestPorts   []string `json:"destPorts,omitempty"`
}

// CPUStatistics is ...
type CPUStatistics struct {
	RefreshedAt               int64   `json:"refreshedAt,omitempty"`
	FiveMinUsageInPercentage  float32 `json:"fiveMinUsageInPercentage,omitempty"`
	FiveSecsUsageInPercentage float32 `json:"fiveSecsUsageInPercentage,omitempty"`
	OneMinUsageInPercentage   float32 `json:"oneMinUsageInPercentage,omitempty"`
}

// DetailedStatus is ...
type DetailedStatus struct {
	AclTraceCalculation              string `json:"aclTraceCalculation,omitempty"`
	AclTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"`
}

// DeviceStatistics is ...
type DeviceStatistics struct {
	MemoryStatistics MemoryStatistics `json:"memoryStatistics,omitempty"` // Device Memory statistics
	CPUStatistics    CPUStatistics    `json:"cpuStatistics,omitempty"`    // Device CPU statistics
}

// FlowAnalysis is ...
type FlowAnalysis struct {
	ID              string   `json:"id,omitempty"`         //
	Status          string   `json:"status,omitempty"`     // Aggregated status of flow-analysis request. Value from a set of [INPROGRESS, COMPLETED, FAILED]
	CreateTime      int64    `json:"createTime,omitempty"` // Flow analysis request creation time
	FailureReason   string   `json:"failureReason,omitempty"`
	LastUpdateTime  int64    `json:"lastUpdateTime,omitempty"`  // Flow analysis request last update time
	Protocol        string   `json:"protocol,omitempty"`        // Protocol
	SourceIP        string   `json:"sourceIP,omitempty"`        // Source IP address
	DestIP          string   `json:"destIP,omitempty"`          // Destination IP address
	SourcePort      string   `json:"sourcePort,omitempty"`      // Source Port
	DestPort        string   `json:"destPort,omitempty"`        // Destination Port
	Inclusions      []string `json:"inclusions,omitempty"`      // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` // periodicRefresh of path for every 30 sec
}

// FlowAnalysisListOutput is ...
type FlowAnalysisListOutput struct {
	Version  string         `json:"version,omitempty"`
	Response []FlowAnalysis `json:"response,omitempty"`
}

// FlowAnalysisRequest is ...
type FlowAnalysisRequest struct {
	Protocol        string   `json:"protocol,omitempty"`        // Protocol
	SourceIP        string   `json:"sourceIP,omitempty"`        // Source IP address
	SourcePort      string   `json:"sourcePort,omitempty"`      // Source Port
	DestIP          string   `json:"destIP,omitempty"`          // Destination IP address
	DestPort        string   `json:"destPort,omitempty"`        // Destination Port
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` // periodicRefresh of path for every 30 sec
	Inclusions      []string `json:"inclusions,omitempty"`      // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
}

// FlowAnalysisRequestResult is ...
type FlowAnalysisRequestResult struct {
	URL            string `json:"url,omitempty"`
	TaskID         string `json:"taskId,omitempty"`
	FlowAnalysisID string `json:"flowAnalysisId,omitempty"`
}

// FlowAnalysisRequestResultOutput is ...
type FlowAnalysisRequestResultOutput struct {
	Version  string                    `json:"version,omitempty"`
	Response FlowAnalysisRequestResult `json:"response,omitempty"`
}

// InterfaceContainer is ...
type InterfaceContainer struct {
	VirtualInterface  []Interface `json:"virtualInterface,omitempty"`  // A list of Virtual interface IDs on a device
	PhysicalInterface Interface   `json:"physicalInterface,omitempty"` // ID of Physical interface on a device
}

// Interface is ...
type Interface struct {
	Name                                  string                  `json:"name,omitempty"`        // Name of interface on a device
	ID                                    string                  `json:"id,omitempty"`          // ID of interface on a device
	VrfName                               string                  `json:"vrfName,omitempty"`     // Name of VRF that the interface on a device belongs to
	AclAnalysis                           AclAnalysisResponse     `json:"aclAnalysis,omitempty"` // Analysis of ACLs on an interface of a device
	QosStatsCollectionFailureReason       string                  `json:"qosStatsCollectionFailureReason,omitempty"`
	InterfaceStatistics                   InterfaceStatistics     `json:"interfaceStatistics,omitempty"`
	QosStatistics                         []QosClassMapStatistics `json:"qosStatistics,omitempty"`
	InterfaceStatsCollection              string                  `json:"interfaceStatsCollection,omitempty"` // A status value from [ INPROGRESS, SUCCESS, FAILED ]
	InterfaceStatsCollectionFailureReason string                  `json:"interfaceStatsCollectionFailureReason,omitempty"`
	QosStatsCollection                    string                  `json:"qosStatsCollection,omitempty"` // A status value from [ INPROGRESS, SUCCESS, FAILED ]
}

// InterfaceStatistics is ...
type InterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`
	RefreshedAt        int64  `json:"refreshedAt,omitempty"`
	InputPackets       int64  `json:"inputPackets,omitempty"`
	InputQueueCount    int32  `json:"inputQueueCount,omitempty"`
	OutputDrop         int64  `json:"outputDrop,omitempty"`
	InputQueueDrops    int64  `json:"inputQueueDrops,omitempty"`
	InputQueueFlushes  int64  `json:"inputQueueFlushes,omitempty"`
	InputQueueMaxDepth int32  `json:"inputQueueMaxDepth,omitempty"`
	InputRatebps       int64  `json:"inputRatebps,omitempty"`
	OperationalStatus  string `json:"operationalStatus,omitempty"`
	OutputPackets      int64  `json:"outputPackets,omitempty"`
	OutputQueueCount   int32  `json:"outputQueueCount,omitempty"`
	OutputQueueDepth   int32  `json:"outputQueueDepth,omitempty"`
	OutputRatebps      int64  `json:"outputRatebps,omitempty"`
}

// MemoryStatistics is ...
type MemoryStatistics struct {
	MemoryUsage int64 `json:"memoryUsage,omitempty"`
	RefreshedAt int64 `json:"refreshedAt,omitempty"`
}

// NetworkElement is ...
type NetworkElement struct {
	Name                               string                  `json:"name,omitempty"`                     // Network Device name
	ID                                 string                  `json:"id,omitempty"`                       // Network Device ID
	Type                               string                  `json:"type,omitempty"`                     // Network Device Type(can be switch,router,wired host or wireless host)
	Role                               string                  `json:"role,omitempty"`                     // Role of device in network(can be access,core,distribution or border router)
	IP                                 string                  `json:"ip,omitempty"`                       // Network Device IP
	IngressVirtualInterface            Interface               `json:"ingressVirtualInterface,omitempty"`  // Ingress virtual interface of the network device
	EgressVirtualInterface             Interface               `json:"egressVirtualInterface,omitempty"`   // Egress virtual interface of the network device
	IngressPhysicalInterface           Interface               `json:"ingressPhysicalInterface,omitempty"` // Igress physical interface of the network device
	EgressPhysicalInterface            Interface               `json:"egressPhysicalInterface,omitempty"`  // Egress physical interface of the network device
	LinkInformationSource              string                  `json:"linkInformationSource,omitempty"`    // The source of the link information to the next hop
	Tunnels                            []string                `json:"tunnels,omitempty"`                  // Tunnels this network element is in
	AccuracyList                       []Accuracy              `json:"accuracyList,omitempty"`
	PerfMonCollection                  string                  `json:"perfMonCollection,omitempty"` // A status value from [ INPROGRESS, SUCCESS, FAILED ]
	DeviceStatistics                   DeviceStatistics        `json:"deviceStatistics,omitempty"`
	PerfMonStatistics                  []PerfMonitorStatistics `json:"perfMonStatistics,omitempty"`
	DeviceStatsCollection              string                  `json:"deviceStatsCollection,omitempty"` // A status value from [ INPROGRESS, SUCCESS, FAILED ]
	PerfMonCollectionFailureReason     string                  `json:"perfMonCollectionFailureReason,omitempty"`
	DeviceStatsCollectionFailureReason string                  `json:"deviceStatsCollectionFailureReason,omitempty"`
	DetailedStatus                     DetailedStatus          `json:"detailedStatus,omitempty"`
}

// NetworkElementInfo is ...
type NetworkElementInfo struct {
	Name                               string                  `json:"name,omitempty"`                  // Network Device name
	ID                                 string                  `json:"id,omitempty"`                    // Network Device ID
	Type                               string                  `json:"type,omitempty"`                  // Network Device Type(can be switch,router,wired host or wireless host)
	Role                               string                  `json:"role,omitempty"`                  // Role of device in network(can be access,core,distribution or border router)
	IP                                 string                  `json:"ip,omitempty"`                    // Network Device IP
	LinkInformationSource              string                  `json:"linkInformationSource,omitempty"` // The source of the link information to the next hop
	Tunnels                            []string                `json:"tunnels,omitempty"`               // Tunnels this network element is in
	AccuracyList                       []Accuracy              `json:"accuracyList,omitempty"`
	PerfMonCollection                  string                  `json:"perfMonCollection,omitempty"`     // A status value from [ INPROGRESS, SUCCESS, FAILED ]
	EgressInterface                    InterfaceContainer      `json:"egressInterface,omitempty"`       // Egress interface of the network device
	IngressInterface                   InterfaceContainer      `json:"ingressInterface,omitempty"`      // Ingress interface of the network device
	DeviceStatistics                   DeviceStatistics        `json:"deviceStatistics,omitempty"`      // Device statistics
	DeviceStatsCollection              string                  `json:"deviceStatsCollection,omitempty"` // A status value from [ INPROGRESS, SUCCESS, FAILED ]
	PerfMonCollectionFailureReason     string                  `json:"perfMonCollectionFailureReason,omitempty"`
	DeviceStatsCollectionFailureReason string                  `json:"deviceStatsCollectionFailureReason,omitempty"`
	DetailedStatus                     DetailedStatus          `json:"detailedStatus,omitempty"`
	PerfMonitorStatistics              []PerfMonitorStatistics `json:"perfMonitorStatistics,omitempty"` // perf mon statistics on the device for give flow
}

// PathResponse is ...
type PathResponse struct {
	Properties          []string             `json:"properties,omitempty"` // Properties for path trace
	Request             FlowAnalysis         `json:"request,omitempty"`    // Describes the source and destination for a path trace
	NetworkElements     []NetworkElement     `json:"networkElements,omitempty"`
	LastUpdate          string               `json:"lastUpdate,omitempty"`          // Last updated time
	NetworkElementsInfo []NetworkElementInfo `json:"networkElementsInfo,omitempty"` // Nodes travesed along a path, including source and destination
	DetailedStatus      DetailedStatus       `json:"detailedStatus,omitempty"`      // Detailed Status of the calculation of Path Trace with its inclusions
}

// PathResponseResult is ...
type PathResponseResult struct {
	Version  string       `json:"version,omitempty"`
	Response PathResponse `json:"response,omitempty"`
}

// PerfMonitorStatistics is ...
type PerfMonitorStatistics struct {
	Protocol             string  `json:"protocol,omitempty"`
	SourcePort           string  `json:"sourcePort,omitempty"`
	DestPort             string  `json:"destPort,omitempty"`
	ByteRate             int64   `json:"byteRate,omitempty"`
	InputInterface       string  `json:"inputInterface,omitempty"`
	Ipv4DSCP             string  `json:"ipv4DSCP,omitempty"`
	Ipv4TTL              int64   `json:"ipv4TTL,omitempty"`
	OutputInterface      string  `json:"outputInterface,omitempty"`
	PacketCount          int64   `json:"packetCount,omitempty"`
	PacketLoss           int64   `json:"packetLoss,omitempty"`
	PacketLossPercentage float32 `json:"packetLossPercentage,omitempty"`
	RefreshedAt          int64   `json:"refreshedAt,omitempty"`
	RtpJitterMean        int64   `json:"rtpJitterMean,omitempty"`
	RtpJitterMax         int64   `json:"rtpJitterMax,omitempty"`
	SourceIPAddress      string  `json:"sourceIpAddress,omitempty"`
	DestIPAddress        string  `json:"destIpAddress,omitempty"`
	RtpJitterMin         int64   `json:"rtpJitterMin,omitempty"`
	PacketBytes          int64   `json:"packetBytes,omitempty"`
}

// QosClassMapStatistics is ...
type QosClassMapStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`
	RefreshedAt        int64  `json:"refreshedAt,omitempty"`
	DropRate           int64  `json:"dropRate,omitempty"`
	NumBytes           int64  `json:"numBytes,omitempty"`
	NumPackets         int64  `json:"numPackets,omitempty"`
	OfferedRate        int64  `json:"offeredRate,omitempty"`
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`
	QueueDepth         int32  `json:"queueDepth,omitempty"`
	QueueNoBufferDrops int64  `json:"queueNoBufferDrops,omitempty"`
	QueueTotalDrops    int64  `json:"queueTotalDrops,omitempty"`
}

// DeleteFlowAnalysis is ...
// Deletes a flow analysis request by its id
//
//
//
//
// * @return *TaskIDResult
func (s *FlowAnalysisService) DeleteFlowAnalysis(flowAnalysisID string) (*TaskIDResult, *Response, error) {

	path := flowAnalysisBasePath + "/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{"+"flowAnalysisId"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAllFlowAnalysisQueryParams is ...
type GetAllFlowAnalysisQueryParams struct {
	PeriodicRefresh bool   `url:"periodicRefresh,omitempty"` // Is analysis periodically refreshed?
	SourceIP        string `url:"sourceIP,omitempty"`        // Source IP address
	DestIP          string `url:"destIP,omitempty"`          // Destination IP adress
	SourcePort      string `url:"sourcePort,omitempty"`      // Source port
	DestPort        string `url:"destPort,omitempty"`        // Destination port
	GtCreateTime    int64  `url:"gtCreateTime,omitempty"`    // Analyses requested after this time
	LtCreateTime    int64  `url:"ltCreateTime,omitempty"`    // Analyses requested before this time
	Protocol        string `url:"protocol,omitempty"`        // Protocol
	Status          string `url:"status,omitempty"`          // Status
	TaskID          string `url:"taskId,omitempty"`          // Task ID
	LastUpdateTime  string `url:"lastUpdateTime,omitempty"`  // Last update time
	Limit           string `url:"limit,omitempty"`           // Number of resources returned
	Offset          string `url:"offset,omitempty"`          // Start index of resources returned (1-based)
	Order           string `url:"order,omitempty"`           // Order by this field
	SortBy          string `url:"sortBy,omitempty"`          // Sort by this field
}

// GetAllFlowAnalysis is ...
// Retrieves a summary of all flow analyses stored. Filters the results by given parameters.
//
//  * @param queryParams
//
//
// * @return *FlowAnalysisListOutput
func (s *FlowAnalysisService) GetAllFlowAnalysis(queryParams *GetAllFlowAnalysisQueryParams) (*FlowAnalysisListOutput, *Response, error) {

	path := flowAnalysisBasePath + "/flow-analysis"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FlowAnalysisListOutput)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetFullFlowAnalysisResult is ...
// Retrieves result of a previously requested flow analysis by its Flow Analysis id
//
//
//
//
// * @return *PathResponseResult
func (s *FlowAnalysisService) GetFullFlowAnalysisResult(flowAnalysisID string) (*PathResponseResult, *Response, error) {

	path := flowAnalysisBasePath + "/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{"+"flowAnalysisId"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PathResponseResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// InitiateFlowAnalysis is ...
// Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to get results and follow progress.
//
//
//
//  * @param flowAnalysisRequest flowAnalysisRequest
// * @return *FlowAnalysisRequestResultOutput
func (s *FlowAnalysisService) InitiateFlowAnalysis(flowAnalysisRequest FlowAnalysisRequest) (*FlowAnalysisRequestResultOutput, *Response, error) {

	path := flowAnalysisBasePath + "/flow-analysis"
	req, err := s.client.NewRequest("POST", path, flowAnalysisRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(FlowAnalysisRequestResultOutput)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
