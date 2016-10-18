package apicem

import (
	"fmt"
	"strings"
)

const ippoolBasePath = "v1"

// IPPoolService is an interface with the IPPool API
type IPPoolService service

// IPPoolInfo is ...
type IPPoolInfo struct {
	Role                           string   `json:"role,omitempty"`                           // Used to group IP Address Pools
	CreateTime                     int64    `json:"createTime,omitempty"`                     // IP Address Pool creation time
	ID                             string   `json:"id,omitempty"`                             // UUID of IP Address Pool
	StartAddress                   string   `json:"startAddress,omitempty"`                   // First IP address in IP Address Pool
	NextAddress                    string   `json:"nextAddress,omitempty"`                    // Next available IP address in IP Address Pool
	EndAddress                     string   `json:"endAddress,omitempty"`                     // Last IP address in IP Address Pool
	IPPool                         string   `json:"ipPool,omitempty"`                         // IP subnet in CIDR format
	ApicAppName                    string   `json:"apicAppName,omitempty"`                    // APIC-EM App Name
	ExcludeNetworkBroadcastAddress bool     `json:"excludeNetworkBroadcastAddress,omitempty"` // If true then network and broadcast IP address will not be used in the usable range. Default value is false
	IPPoolName                     string   `json:"ipPoolName,omitempty"`                     // IP Address Pool name
	DHCPServerIP                   []string `json:"dhcpServerIp,omitempty"`                   // DHCP server hostname or IP address list
	Gateway                        []string `json:"gateway,omitempty"`                        // Gateway IP address list
	InterAppOverlap                bool     `json:"interAppOverlap,omitempty"`                // If true then overlapping IP pool is supported between APIC-EM Apps. Default value is false
	CreationOrder                  int32    `json:"creationOrder,omitempty"`                  // Creation order of IP Address Pool
	LastUpdateTime                 int64    `json:"lastUpdateTime,omitempty"`                 // IP Address Pool last updated time
	ParentID                       string   `json:"parentId,omitempty"`                       // Parent IP Address Pool UUID
	FreeIPCount                    int32    `json:"freeIpCount,omitempty"`                    // IP Address Pool free IP count
	UsagePercentage                int32    `json:"usagePercentage,omitempty"`                // Current usage percentage of IP Address Pool
	Shared                         bool     `json:"shared,omitempty"`                         // If true then duplicate/overlapping pool is supported. Default value is false
}

// IPPoolInfoListResult is ...
type IPPoolInfoListResult struct {
	Version  string       `json:"version,omitempty"`
	Response []IPPoolInfo `json:"response,omitempty"`
}

// IPPoolInfoResult is ...
type IPPoolInfoResult struct {
	Version  string     `json:"version,omitempty"`
	Response IPPoolInfo `json:"response,omitempty"`
}

// CreateIPPool is ...
// This API is used to create global IP address pool(s).
//
//
//  * @param scope Authorization Scope for RBAC * @param username requestorUsername
//  * @param ipPoolInfoList List of IPPoolInfo objects
// * @return *TaskIDResult
func (s *IPPoolService) CreateIPPool(ipPoolInfoList *[]IPPoolInfo, scope string, username string) (*TaskIDResult, *Response, error) {

	path := ippoolBasePath + "/ippool"
	req, err := s.client.NewRequest("POST", path, ipPoolInfoList)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)
	req.Header.Add("username", username)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteIPPool is ...
// This API is used to delete global IP address pool based on its ID.
//
//
//  * @param scope Authorization Scope for RBAC * @param username requestorUsername
//
// * @return *TaskIDResult
func (s *IPPoolService) DeleteIPPool(id string, scope string, username string) (*TaskIDResult, *Response, error) {

	path := ippoolBasePath + "/ippool/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)
	req.Header.Add("username", username)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Get is ...
// This API is used to retrieve IP address pool based on its ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *IPPoolInfoResult
func (s *IPPoolService) Get(id string, scope string) (*IPPoolInfoResult, *Response, error) {

	path := ippoolBasePath + "/ippool/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(IPPoolInfoResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCount is ...
// This API is used to retrieve the number of configured IP address pools
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *IPPoolService) GetCount(scope string) (*CountResult, *Response, error) {

	path := ippoolBasePath + "/ippool/count"
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

// GetListQueryParams is ...
type GetListQueryParams struct {
	Limit       string `url:"limit,omitempty"`       // limit
	Offset      string `url:"offset,omitempty"`      // offset
	ApicAppName string `url:"apicAppName,omitempty"` // Apic APP Name
}

// GetList is ...
// This API is used to retrieve list of configured IP address pools.
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *IPPoolInfoListResult
func (s *IPPoolService) GetList(scope string, queryParams *GetListQueryParams) (*IPPoolInfoListResult, *Response, error) {

	path := ippoolBasePath + "/ippool"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(IPPoolInfoListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
