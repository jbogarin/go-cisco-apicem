package apicem

const topologyVLANBasePath = "v1"

// TopologyVLANService is an interface with the  API
type TopologyVLANService service

// VLANNamesResult is ...
type VLANNamesResult struct {
	Version  string   `json:"version,omitempty"`
	Response []string `json:"response,omitempty"`
}

// GetVLANNames is ...
// This method is used to obtain the list of vlan names
//
//
//
//
// * @return *VLANNamesResult
func (s *TopologyVLANService) GetVLANNames() (*VLANNamesResult, *Response, error) {

	path := topologyVLANBasePath + "/topology/vlan/vlan-names"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(VLANNamesResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetVLANNames1 is ...
// This method is used to obtain the list of vlan names
//
//
//
//
// * @return *VLANNamesResult
func (s *TopologyVLANService) GetVLANNames1() (*VLANNamesResult, *Response, error) {

	path := topologyVLANBasePath + "/vlan/vlan-names"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(VLANNamesResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
