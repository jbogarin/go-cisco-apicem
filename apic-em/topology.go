package apicem

import (
	"fmt"
	"strings"
)

const topologyBasePath = "v1"

// TopologyService is an interface with the Topology API
type TopologyService service

// TopologyResult is ...
type TopologyResult struct {
	Version  string   `json:"version,omitempty"`
	Response Topology `json:"response,omitempty"`
}

// GetL2Topology is ...
// This method is used to obtain the Layer 2 topology by Vlan ID
//
//
//
//
// * @return *TopologyResult
func (s *TopologyService) GetL2Topology(vlanID string) (*TopologyResult, *Response, error) {

	path := topologyBasePath + "/topology/l2/{vlanID}"
	path = strings.Replace(path, "{"+"vlanID"+"}", fmt.Sprintf("%v", vlanID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetL3Topology is ...
// This method is used to obtain Layer 3 device topology by routing protocol type
//
//
//
//
// * @return *TopologyResult
func (s *TopologyService) GetL3Topology(topologyType string) (*TopologyResult, *Response, error) {

	path := topologyBasePath + "/topology/l3/{topologyType}"
	path = strings.Replace(path, "{"+"topologyType"+"}", fmt.Sprintf("%v", topologyType), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetL3TopologyForVrf is ...
// This method is used to obtain Layer 3 device topology by vrf name
//
//
//
//
// * @return *TopologyResult
func (s *TopologyService) GetL3TopologyForVrf(vrfName string) (*TopologyResult, *Response, error) {

	path := topologyBasePath + "/topology/l3/vrf/{vrfName}"
	path = strings.Replace(path, "{"+"vrfName"+"}", fmt.Sprintf("%v", vrfName), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPhysicalTopology is ...
// This method is used to obtain the raw physical topology
//
//
//
//
// * @return *TopologyResult
func (s *TopologyService) GetPhysicalTopology() (*TopologyResult, *Response, error) {

	path := topologyBasePath + "/topology/physical-topology"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// LoadCustomTopology is ...
// This method is used to obtain the topology by customized layout
//
//
//
//
// * @return *TopologyResult
func (s *TopologyService) LoadCustomTopology() (*TopologyResult, *Response, error) {

	path := topologyBasePath + "/topology/custom"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// SaveCustomTopology is ...
// This method is used to save the topology in canvas implementation
//
//
//
//  * @param topo Topology
// * @return *TaskIDResult
func (s *TopologyService) SaveCustomTopology(topo Topology) (*TaskIDResult, *Response, error) {

	path := topologyBasePath + "/topology/custom"
	req, err := s.client.NewRequest("POST", path, topo)
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
