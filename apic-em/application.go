package apicem

import (
	"fmt"
	"strings"
)

const applicationBasePath = "v1"

// ApplicationService is an interface with the Application API
type ApplicationService service

// ApplicationDTO is ...
type ApplicationDTO struct {
	TrafficClass                    string `json:"trafficClass,omitempty"`                    // Traffic class to which the app belongs
	Rank                            int32  `json:"rank,omitempty"`                            // rank
	Status                          string `json:"status,omitempty"`                          // Gives status of the app
	Category                        string `json:"category,omitempty"`                        // Category name
	AppProtocol                     string `json:"appProtocol,omitempty"`                     // protocol of the app. Valid values are tcp, udp, tcp/udp, ip or it could be empty. Values are case sensitive.
	TCPPorts                        string `json:"tcpPorts,omitempty"`                        // list of tcp ports
	IndicativeTCPPorts              string `json:"indicativeTcpPorts,omitempty"`              // Indicative tcp ports
	UDPPorts                        string `json:"udpPorts,omitempty"`                        // list of udp ports
	IndicativeUDPPorts              string `json:"indicativeUdpPorts,omitempty"`              // Indicative udp ports
	IPPorts                         string `json:"ipPorts,omitempty"`                         // list of ip ports
	URL                             string `json:"url,omitempty"`                             // url of the app
	LongDescription                 string `json:"longDescription,omitempty"`                 // Long description of the app
	SubCategory                     string `json:"subCategory,omitempty"`                     // Sub-Category Id
	GlobalID                        string `json:"globalId,omitempty"`                        // global id
	EngineID                        string `json:"engineId,omitempty"`                        // engine id
	SelectorID                      string `json:"selectorId,omitempty"`                      // selector id
	HelpString                      string `json:"helpString,omitempty"`                      // help string to describe the app
	References                      string `json:"references,omitempty"`                      // references of the app
	ApplicationGroup                string `json:"applicationGroup,omitempty"`                // App group name
	Encrypted                       string `json:"encrypted,omitempty"`                       // If the app is encrypted
	Tunnel                          string `json:"tunnel,omitempty"`                          // If the app is a tunnel
	InstanceUUID                    string `json:"instanceUuid,omitempty"`                    //
	CategoryID                      string `json:"categoryId,omitempty"`                      // Category id
	Dscp                            string `json:"dscp,omitempty"`                            // dscp value
	PfrThresholdJitter              int32  `json:"pfrThresholdJitter,omitempty"`              // PfR Threshold Jitter
	PfrThresholdLossRate            int32  `json:"pfrThresholdLossRate,omitempty"`            // PfR Threshold Loss Rate
	PfrThresholdOneWayDelay         int32  `json:"pfrThresholdOneWayDelay,omitempty"`         // PfR Threshold One Way Delay
	Popularity                      int32  `json:"popularity,omitempty"`                      // popularity of the app
	TransportIps                    string `json:"transportIps,omitempty"`                    // Transport IP of the app
	Enabled                         string `json:"enabled,omitempty"`                         // If the app enabled
	IndicativeIPPorts               string `json:"indicativeIpPorts,omitempty"`               // Indicative ip ports
	IsRepresentativeApp             bool   `json:"isRepresentativeApp,omitempty"`             // If the app is representative
	NbarID                          string `json:"nbarId,omitempty"`                          // nbar id
	P2pTechnology                   string `json:"p2pTechnology,omitempty"`                   // If the app is a p2p technology
	PfrThresholdJitterPriority      int32  `json:"pfrThresholdJitterPriority,omitempty"`      // PfR Threshold Jitter Priority
	PfrThresholdLossRatePriority    int32  `json:"pfrThresholdLossRatePriority,omitempty"`    // PfR Threshold Loss Rate Priority
	PfrThresholdOneWayDelayPriority int32  `json:"pfrThresholdOneWayDelayPriority,omitempty"` // PfR Threshold One Way Delay Priority
	IgnoreConflict                  bool   `json:"ignoreConflict,omitempty"`                  // If true ignore conflicts with other Applications
	Name                            string `json:"name,omitempty"`                            // App Name
	ID                              string `json:"id,omitempty"`                              // id
}

// ApplicationListResult is ...
type ApplicationListResult struct {
	Response []ApplicationDTO `json:"response,omitempty"`
	Version  string           `json:"version,omitempty"`
}

// ApplicationResult is ...
type ApplicationResult struct {
	Response ApplicationDTO `json:"response,omitempty"`
	Version  string         `json:"version,omitempty"`
}

// AddApplicationQueryParams is ...
type AddApplicationQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // scheduleAt
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // scheduleDesc
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // scheduleOrigin
}

// AddApplication is ...
// Create an application
//
//  * @param queryParams
//
//  * @param applicationDTOList applicationDTOList
// * @return *TaskIDResult
func (s *ApplicationService) AddApplication(queryParams *AddApplicationQueryParams, applicationDTOList []ApplicationDTO) (*TaskIDResult, *Response, error) {

	path := applicationBasePath + "/application"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("POST", path, applicationDTOList)
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

// DeleteApplicationQueryParams is ...
type DeleteApplicationQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // scheduleAt
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // scheduleDesc
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // scheduleOrigin
}

// DeleteApplication is ...
// Delete an application by id
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *ApplicationService) DeleteApplication(id string, queryParams *DeleteApplicationQueryParams) (*TaskIDResult, *Response, error) {

	path := applicationBasePath + "/application/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// GetApplication is ...
// Get an application by id
//
//
//
//
// * @return *ApplicationResult
func (s *ApplicationService) GetApplication(id string) (*ApplicationResult, *Response, error) {

	path := applicationBasePath + "/application/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ApplicationResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetApplicationCount is ...
// Get total count of application(s)
//
//
//
//
// * @return *CountResult
func (s *ApplicationService) GetApplicationCount() (*CountResult, *Response, error) {

	path := applicationBasePath + "/application/count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CountResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetFilterApplicationQueryParams is ...
type GetFilterApplicationQueryParams struct {
	IsCustom         string `url:"isCustom,omitempty"`         // Retrieve custom applications
	IsRepresentative string `url:"isRepresentative,omitempty"` // Retrieve representative applications
	CategoryID       string `url:"categoryId,omitempty"`       // Retrieve applications by categoryId
	Name             string `url:"name,omitempty"`             // name
	Offset           string `url:"offset,omitempty"`           // Starting index of the resources (1 based), This should be only used in conjuction with the limit param.
	Limit            string `url:"limit,omitempty"`            // Number of resources to return, WARNING: results may take an unexpectely long time with more than 50 results requested.
}

// GetFilterApplication is ...
// Get application(s) based on a filter provided
//
//  * @param queryParams
//
//
// * @return *ApplicationListResult
func (s *ApplicationService) GetFilterApplication(queryParams *GetFilterApplicationQueryParams) (*ApplicationListResult, *Response, error) {

	path := applicationBasePath + "/application"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ApplicationListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateApplicationQueryParams is ...
type UpdateApplicationQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // scheduleAt
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // scheduleDesc
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // scheduleOrigin
}

// UpdateApplication is ...
// Update an application
//
//  * @param queryParams
//
//  * @param applicationDTOList applicationDTOList
// * @return *TaskIDResult
func (s *ApplicationService) UpdateApplication(queryParams *UpdateApplicationQueryParams, applicationDTOList []ApplicationDTO) (*TaskIDResult, *Response, error) {

	path := applicationBasePath + "/application"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("PUT", path, applicationDTOList)
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
