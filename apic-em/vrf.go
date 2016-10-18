package apicem

const vrfBasePath = "v1"

// VRFService is an interface with the Topologyvrfvrfname API
type VRFService service

// VrfNamesResult is ...
type VrfNamesResult struct {
	Version  string   `json:"version,omitempty"`
	Response []string `json:"response,omitempty"`
}

// GetVLANNames is ...
// This method is used to obtain the list of vrf names
//
//
//
//
// * @return *VrfNamesResult
func (s *VRFService) GetVLANNames() (*VrfNamesResult, *Response, error) {

	path := vrfBasePath + "/topology/vrf/vrf-name"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(VrfNamesResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
