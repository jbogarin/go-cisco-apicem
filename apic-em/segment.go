package apicem

const segmentBasePath = "v1"

// SegmentService is an interface with the Segment API
type SegmentService service

// GetSegmentInfoQueryParams is ...
type GetSegmentInfoQueryParams struct {
	Type      string `url:"type,omitempty"`
	PolicyTag string `url:"policyTag,omitempty"`
}

// NetworkDeviceBrief is ...
type NetworkDeviceBrief struct {
	ID       string `json:"id,omitempty"`
	HostName string `json:"hostName,omitempty"`
}

// SegmentDto is ...
type SegmentDto struct {
	Name           string               `json:"name,omitempty"`
	NetworkDevices []NetworkDeviceBrief `json:"networkDevices,omitempty"`
}

// SegmentResult is ...
type SegmentResult struct {
	Version  string       `json:"version,omitempty"`
	Response []SegmentDto `json:"response,omitempty"`
}

// GetSegmentInfo is ...
// Gets list of segment info based on policyTag
//
//  * @param type_ Type of segment
//  * @param policyTag Policy tag
//  * @return *SegmentResult
func (s *SegmentService) GetSegmentInfo(queryParams *GetSegmentInfoQueryParams) (*SegmentResult, *Response, error) {

	path := segmentBasePath + "/segment"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SegmentResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
