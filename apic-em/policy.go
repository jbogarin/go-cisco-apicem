package apicem

import (
	"fmt"
	"strings"
)

const policyBasePath = "v1"

// PolicyService is an interface with the Policy API
type PolicyService service

// ActionProperty is ...
type ActionProperty struct {
	Destinations       []string `json:"destinations,omitempty"`
	RelevanceLevel     string   `json:"relevanceLevel,omitempty"`     // relevance level for a policy
	PriorityLevel      string   `json:"priorityLevel,omitempty"`      // priority level for a policy
	ExperienceLevel    string   `json:"experienceLevel,omitempty"`    // experience level for a policy
	PathPreference     string   `json:"pathPreference,omitempty"`     // path preference for a policy
	PathControlFlag    bool     `json:"pathControlFlag,omitempty"`    // path control flag
	PathPreferenceFlag bool     `json:"pathPreferenceFlag,omitempty"` // path preference flag
	PathOfLastResort   string   `json:"PathOfLastResort,omitempty"`
	MaintainExperience string   `json:"maintainExperience,omitempty"`
	PrimaryPathPref    []string `json:"PrimaryPathPref,omitempty"`
	SecondaryPathPref  []string `json:"SecondaryPathPref,omitempty"`
	TertiaryPathPref   []string `json:"TertiaryPathPref,omitempty"`
	TrustLevel         string   `json:"trustLevel,omitempty"` // trust level for a policy
}

// FlowDto is ...
type FlowDto struct {
	Status            string `json:"status,omitempty"`
	Codec             string `json:"codec,omitempty"`             // codec
	SourcePort        string `json:"sourcePort,omitempty"`        // sourcePort
	DestPort          string `json:"destPort,omitempty"`          // destPort
	FlowType          string `json:"flowType,omitempty"`          // flowType
	SourceIP          string `json:"sourceIP,omitempty"`          // sourceIP
	DestIP            string `json:"destIP,omitempty"`            // destIP
	DetailedFlowType  string `json:"detailedFlowType,omitempty"`  // detailedFlowType (more detailed classification than flowType)
	AverageBandwidth  string `json:"averageBandwidth,omitempty"`  // averageBandwidth in kbps (min: 0, max: 2147483647 kbps)
	PeakBandwidth     string `json:"peakBandwidth,omitempty"`     // peakBandwidth in kbps (min: 0, max: 2147483647 kbps)
	ClientReference   string `json:"clientReference,omitempty"`   // clientReference (can be used by the client as a reference to this resource
	NetworkDeviceID   string `json:"networkDeviceId,omitempty"`   // networkDeviceId
	NetworkDeviceName string `json:"networkDeviceName,omitempty"` // networkDeviceName
	InterfaceID       string `json:"interfaceId,omitempty"`       // interfaceId
	InterfaceName     string `json:"interfaceName,omitempty"`     // interfaceName
	FailureReason     string `json:"failureReason,omitempty"`
	ID                string `json:"id,omitempty"`       // id
	Protocol          string `json:"protocol,omitempty"` // protocol
}

// FlowIDResponse is ...
type FlowIDResponse struct {
	TaskID string `json:"taskId,omitempty"`
	FlowID string `json:"flowId,omitempty"`
}

// FlowIDResult is ...
type FlowIDResult struct {
	Response FlowIDResponse `json:"response,omitempty"`
	Version  string         `json:"version,omitempty"`
}

// FlowListResult is ...
type FlowListResult struct {
	Response []FlowDto `json:"response,omitempty"`
	Version  string    `json:"version,omitempty"`
}

// FlowResult is ...
type FlowResult struct {
	Response FlowDto `json:"response,omitempty"`
	Version  string  `json:"version,omitempty"`
}

// NetworkUser is ...
type NetworkUser struct {
	UserIdentifiers  []string                   `json:"userIdentifiers,omitempty"`
	DeviceTypes      []string                   `json:"deviceTypes,omitempty"`
	Locations        []string                   `json:"locations,omitempty"`
	Applications     []PolicyApplication        `json:"applications,omitempty"`
	ApplicationDiffs []PolicyApplicationDiffDto `json:"applicationDiffs,omitempty"`
	Categories       []CategoryDto              `json:"categories,omitempty"`
}

// PolicyApplicationDiffDto is ...
type PolicyApplicationDiffDto struct {
	Operation         string            `json:"operation,omitempty"`         // operation
	PolicyApplication PolicyApplication `json:"policyApplication,omitempty"` // application
}

// PolicyApplication is ...
type PolicyApplication struct {
	Raw     string `json:"raw,omitempty"` // Either raw Application of the form port:protocol should be specified or appID should be specified
	ID      string `json:"id,omitempty"`  // id
	AppName string `json:"appName,omitempty"`
}

// PolicyDiffDto is ...
type PolicyDiffDto struct {
	Policy    Policy `json:"policy,omitempty"`    // policy
	Operation string `json:"operation,omitempty"` // operation
}

// Policy is ...
type Policy struct {
	ID                   string         `json:"id,omitempty"`           // id
	PolicyScope          string         `json:"policyScope,omitempty"`  // policyScope
	InstanceUUID         string         `json:"instanceUuid,omitempty"` //
	Resource             PolicyResource `json:"resource,omitempty"`     // Resource
	PolicyName           string         `json:"policyName,omitempty"`   // name of the policy
	TaskID               string         `json:"taskId,omitempty"`       // Task ID
	ScopeWirelessSegment string         `json:"scopeWirelessSegment,omitempty"`
	NetworkUser          NetworkUser    `json:"networkUser,omitempty"`    // Network User
	PolicyOwner          string         `json:"policyOwner,omitempty"`    // Policy Owner
	ActionProperty       ActionProperty `json:"actionProperty,omitempty"` // ActionProperty
	PolicyPriority       int32          `json:"policyPriority,omitempty"` // Policy Priority
	Actions              []string       `json:"actions,omitempty"`        // Action Set
	State                string         `json:"state,omitempty"`
}

// PolicyFlowSettingDto is ...
type PolicyFlowSettingDto struct {
	FlowPolicyEnabled bool `json:"flowPolicyEnabled,omitempty"` // Allowed values are true, false
}

// PolicyFlowSettingResult is ...
type PolicyFlowSettingResult struct {
	Response PolicyFlowSettingDto `json:"response,omitempty"`
	Version  string               `json:"version,omitempty"`
}

// PolicyListResult is ...
type PolicyListResult struct {
	Response []Policy `json:"response,omitempty"`
	Version  string   `json:"version,omitempty"`
}

// PolicyResource is ...
type PolicyResource struct {
	UserIdentifiers  []string                   `json:"userIdentifiers,omitempty"`
	DeviceTypes      []string                   `json:"deviceTypes,omitempty"`
	Locations        []string                   `json:"locations,omitempty"`
	Applications     []PolicyApplication        `json:"applications,omitempty"`
	ApplicationDiffs []PolicyApplicationDiffDto `json:"applicationDiffs,omitempty"`
	Categories       []CategoryDto              `json:"categories,omitempty"`
}

// PolicyResult is ...
type PolicyResult struct {
	Response Policy `json:"response,omitempty"`
	Version  string `json:"version,omitempty"`
}

// PolicyStatusDto is ...
type PolicyStatusDto struct {
	Status                         string              `json:"status,omitempty"`
	InstanceUUID                   string              `json:"instanceUuid,omitempty"`
	PolicyScope                    string              `json:"policyScope,omitempty"`
	ScopeWirelessSegment           string              `json:"scopeWirelessSegment,omitempty"`
	NetworkDeviceID                string              `json:"networkDeviceId,omitempty"`
	NetworkDeviceName              string              `json:"networkDeviceName,omitempty"`
	FailureReason                  string              `json:"failureReason,omitempty"`
	LastUpdated                    string              `json:"lastUpdated,omitempty"`
	OutOfScope                     bool                `json:"outOfScope,omitempty"`
	BusinessRelevantApplications   []PolicyApplication `json:"businessRelevantApplications,omitempty"`
	BusinessIrrelevantApplications []PolicyApplication `json:"businessIrrelevantApplications,omitempty"`
	NetworkDeviceIP                string              `json:"networkDeviceIp,omitempty"`
	ApplicationPolicyCount         int32               `json:"applicationPolicyCount,omitempty"`
}

// PolicyStatusListResult is ...
type PolicyStatusListResult struct {
	Response []PolicyStatusDto `json:"response,omitempty"`
	Version  string            `json:"version,omitempty"`
}

// PolicyTagAssociationDeviceDto is ...
type PolicyTagAssociationDeviceDto struct {
	DeviceID   string `json:"deviceId,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceIP   string `json:"deviceIp,omitempty"`
}

// PolicyTagAssociationDto is ...
type PolicyTagAssociationDto struct {
	NetworkDevices     []PolicyTagAssociationDeviceDto `json:"networkDevices,omitempty"`
	PolicyTag          string                          `json:"policyTag,omitempty"`
	UnModifiable       bool                            `json:"unModifiable,omitempty"`
	UnModifiableReason string                          `json:"unModifiableReason,omitempty"`
}

// PolicyTagAssociationListResult is ...
type PolicyTagAssociationListResult struct {
	Response []PolicyTagAssociationDto `json:"response,omitempty"`
	Version  string                    `json:"version,omitempty"`
}

// PolicyTagDto is ...
type PolicyTagDto struct {
	PolicyTag string `json:"policyTag,omitempty"` // Policy Tag value
}

// PolicyTagListResult is ...
type PolicyTagListResult struct {
	Response []PolicyTagDto `json:"response,omitempty"`
	Version  string         `json:"version,omitempty"`
}

// PolicyVersionNumberListResult is ...
type PolicyVersionNumberListResult struct {
	Response []VersionNumberDto `json:"response,omitempty"`
	Version  string             `json:"version,omitempty"`
}

// VersionDiffDto is ...
type VersionDiffDto struct {
	ID       string          `json:"id,omitempty"`       // id
	Policies []PolicyDiffDto `json:"policies,omitempty"` // policies
}

// VersionDiffResult is ...
type VersionDiffResult struct {
	Response VersionDiffDto `json:"response,omitempty"`
	Version  string         `json:"version,omitempty"`
}

// VersionNumberDto is ...
type VersionNumberDto struct {
	Version    int64 `json:"version,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
}

// AddQueryParams is ...
type AddQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the policy should be scheduled (Optional)
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // Custom Description (Optional)
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // Originator of this call (Optional)
}

// Add is ...
// Create a policy
//
//  * @param queryParams
//
//  * @param policyList Policy Object
// * @return *TaskIDResult
func (s *PolicyService) Add(policyList []*Policy, queryParams *AddQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("POST", path, policyList)
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

// AddFlow is ...
// addFlow
//
//
//
//  * @param flowDTO flowDTO
// * @return *FlowIDResult
func (s *PolicyService) AddFlow(flowDTO FlowDto) (*FlowIDResult, *Response, error) {

	path := policyBasePath + "/policy/flow"
	req, err := s.client.NewRequest("POST", path, flowDTO)
	if err != nil {
		return nil, nil, err
	}

	root := new(FlowIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// AddPolicyDiff is ...
// Adds the policy diff
//
//
//
//  * @param versionDiffDTO VersionDiff
// * @return *TaskIDResult
func (s *PolicyService) AddPolicyDiff(versionDiffDTO VersionDiffDto) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/diff"
	req, err := s.client.NewRequest("POST", path, versionDiffDTO)
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

// AddPolicyTag is ...
// addPolicyTag
//
//
//
//  * @param policyTagDTO policyTagDTO
// * @return *TaskIDResult
func (s *PolicyService) AddPolicyTag(policyTagDTO PolicyTagDto) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/tag"
	req, err := s.client.NewRequest("POST", path, policyTagDTO)
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

// AddPolicyTagAssociation is ...
// addPolicyTagAssociation
//
//
//
//  * @param policyTagAssociationDTO policyTagAssociationDTO
// * @return *TaskIDResult
func (s *PolicyService) AddPolicyTagAssociation(policyTagAssociationDTO PolicyTagAssociationDto) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/tag/association"
	req, err := s.client.NewRequest("POST", path, policyTagAssociationDTO)
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

// Delete is ...
// Delete a policy by id
//
//
//
//
// * @return *TaskIDResult
func (s *PolicyService) Delete(id string) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
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

// DeleteFilterPoliciesQueryParams is ...
type DeleteFilterPoliciesQueryParams struct {
	PolicyScope          string `url:"policyScope,omitempty"`          // Delete policies for a given scope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Delete policies for a given wireless segment, within a policyScope (policyScope is mandatory for this)
}

// DeleteFilterPolicies is ...
// Delete policies based on a given filter
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *PolicyService) DeleteFilterPolicies(queryParams *DeleteFilterPoliciesQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy"
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

// DeleteFlow is ...
// Delete a policy for the flow by its id
//
//
//
//
// * @return *TaskIDResult
func (s *PolicyService) DeleteFlow(id string) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/flow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
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

// DeleteFlowsByClientReferenceQueryParams is ...
type DeleteFlowsByClientReferenceQueryParams struct {
	ClientReference string `url:"clientReference,omitempty"` // clientReference
}

// DeleteFlowsByClientReference is ...
// deleteFlowsByClientReference
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *PolicyService) DeleteFlowsByClientReference(queryParams *DeleteFlowsByClientReferenceQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/flow"
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

// DeletePolicyTagQueryParams is ...
type DeletePolicyTagQueryParams struct {
	PolicyTag string `url:"policyTag,omitempty"` // Policy Tag
}

// DeletePolicyTag is ...
// deletePolicyTag
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *PolicyService) DeletePolicyTag(queryParams *DeletePolicyTagQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/tag"
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

// DeletePolicyTagAssociationQueryParams is ...
type DeletePolicyTagAssociationQueryParams struct {
	PolicyTag       string `url:"policyTag,omitempty"`       // Policy Tag
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` // NetworkDeviceId
}

// DeletePolicyTagAssociation is ...
// deletePolicyTagAssociation
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *PolicyService) DeletePolicyTagAssociation(queryParams *DeletePolicyTagAssociationQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/tag/association"
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

// Get is ...
// Get a policy by id
//
//
//
//
// * @return *PolicyResult
func (s *PolicyService) Get(id string) (*PolicyResult, *Response, error) {

	path := policyBasePath + "/policy/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCount is ...
// Get total count of policies
//
//
//
//
// * @return *CountResult
func (s *PolicyService) GetCount() (*CountResult, *Response, error) {

	path := policyBasePath + "/policy/count"
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

// GetCountByPolicyScopeQueryParams is ...
type GetCountByPolicyScopeQueryParams struct {
	PolicyScope          string `url:"policyScope,omitempty"`          // Retrieve count of policy status(es) based on a given policyScope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Retrieve count of policy status(es) based on a given scopeWirelessSegment
}

// GetCountByPolicyScope is ...
// Get count of policy status(es) based on a given policyScope &amp; scopeWirelessSegment
//
//  * @param queryParams
//
//
// * @return *CountResult
func (s *PolicyService) GetCountByPolicyScope(queryParams *GetCountByPolicyScopeQueryParams) (*CountResult, *Response, error) {

	path := policyBasePath + "/policy/status/count"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// GetFilterPoliciesQueryParams is ...
type GetFilterPoliciesQueryParams struct {
	PolicyScope          string `url:"policyScope,omitempty"`          // Retrieve policies for a given scope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Retrieve policies for a given wireless segment, within a policyScope (policyScope is mandatory for this)
	ApplicationID        string `url:"applicationId,omitempty"`        // Retrieve policies for a given Resource Application Id
	Offset               string `url:"offset,omitempty"`               // Starting index of the resources (1 based)
	Limit                string `url:"limit,omitempty"`                // Number of resources to return
}

// GetFilterPolicies is ...
// Get policy(s) based on a filter provided
//
//  * @param queryParams
//
//
// * @return *PolicyListResult
func (s *PolicyService) GetFilterPolicies(queryParams *GetFilterPoliciesQueryParams) (*PolicyListResult, *Response, error) {

	path := policyBasePath + "/policy"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetFlow is ...
// getFlow
//
//
//
//
// * @return *FlowResult
func (s *PolicyService) GetFlow(id string) (*FlowResult, *Response, error) {

	path := policyBasePath + "/policy/flow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FlowResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetFlowCount is ...
// getFlowCount
//
//
//
//
// * @return *CountResult
func (s *PolicyService) GetFlowCount() (*CountResult, *Response, error) {

	path := policyBasePath + "/policy/flow/count"
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

// GetFlowsQueryParams is ...
type GetFlowsQueryParams struct {
	Offset string `url:"offset,omitempty"` // Starting index of the resources (1 based)
	Limit  string `url:"limit,omitempty"`  // Number of resources to return
}

// GetFlows is ...
// getFlows
//
//  * @param queryParams
//
//
// * @return *FlowListResult
func (s *PolicyService) GetFlows(queryParams *GetFlowsQueryParams) (*FlowListResult, *Response, error) {

	path := policyBasePath + "/policy/flow"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FlowListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyDiffQueryParams is ...
type GetPolicyDiffQueryParams struct {
	Version              string `url:"version,omitempty"`              // Retrieve policy diffs for the given version
	PolicyScope          string `url:"policyScope,omitempty"`          // Retrieve policy diffs for the given scope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Retrieve policy diffs for the given scopeWirelessSegment
}

// GetPolicyDiff is ...
// Retrieves the policy diff
//
//  * @param queryParams
//
//
// * @return *VersionDiffResult
func (s *PolicyService) GetPolicyDiff(queryParams *GetPolicyDiffQueryParams) (*VersionDiffResult, *Response, error) {

	path := policyBasePath + "/policy/diff"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(VersionDiffResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyFlowSetting is ...
// getPolicyFlowSetting
//
//
//
//
// * @return *PolicyFlowSettingResult
func (s *PolicyService) GetPolicyFlowSetting() (*PolicyFlowSettingResult, *Response, error) {

	path := policyBasePath + "/policy/flow/setting"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyFlowSettingResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyStatusQueryParams is ...
type GetPolicyStatusQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` // Retrieve policies for a given networkDeviceId
	PolicyScope     string `url:"policyScope,omitempty"`     // Retrieve policies for a given scope
	Offset          string `url:"offset,omitempty"`          // Starting index of the resources (1 based)
	Limit           string `url:"limit,omitempty"`           // Number of resources to return (Use smaller limit value for better response times. Max 50)
}

// GetPolicyStatus is ...
// Get policy status(s) based on filter(s) provided. Either provide networkDeviceId or provide policyScope, offset, limit
//
//  * @param queryParams
//
//
// * @return *PolicyStatusListResult
func (s *PolicyService) GetPolicyStatus(queryParams *GetPolicyStatusQueryParams) (*PolicyStatusListResult, *Response, error) {

	path := policyBasePath + "/policy/status"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyStatusListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyStatusSummaryQueryParams is ...
type GetPolicyStatusSummaryQueryParams struct {
	PolicyScope string `url:"policyScope,omitempty"` // Retrieve policy status(es) for a given scope
	Offset      string `url:"offset,omitempty"`      // Starting index of the resources (1 based)
	Limit       string `url:"limit,omitempty"`       // Number of resources to return (Use smaller limit value for better response times. Max 50)
}

// GetPolicyStatusSummary is ...
// Get policy status(s) summary based on a filter provided
//
//  * @param queryParams
//
//
// * @return *PolicyStatusListResult
func (s *PolicyService) GetPolicyStatusSummary(queryParams *GetPolicyStatusSummaryQueryParams) (*PolicyStatusListResult, *Response, error) {

	path := policyBasePath + "/policy/status/summary"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyStatusListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyTagAssociationQueryParams is ...
type GetPolicyTagAssociationQueryParams struct {
	PolicyTag       string `url:"policyTag,omitempty"`       // Policy Tag
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` // Network device Id
	Offset          string `url:"offset,omitempty"`          // Starting index of the resources (1 based)
	Limit           string `url:"limit,omitempty"`           // Number of resources to return
}

// GetPolicyTagAssociation is ...
// getPolicyTagAssociation
//
//  * @param queryParams
//
//
// * @return *PolicyTagAssociationListResult
func (s *PolicyService) GetPolicyTagAssociation(queryParams *GetPolicyTagAssociationQueryParams) (*PolicyTagAssociationListResult, *Response, error) {

	path := policyBasePath + "/policy/tag/association"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyTagAssociationListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyTagsQueryParams is ...
type GetPolicyTagsQueryParams struct {
	PolicyTag       string `url:"policyTag,omitempty"`       // policyTag search substring
	FilterOperation string `url:"filterOperation,omitempty"` // type of search (startsWith, endsWith, contains
}

// GetPolicyTags is ...
// getPolicyTags
//
//  * @param queryParams
//
//
// * @return *PolicyTagListResult
func (s *PolicyService) GetPolicyTags(queryParams *GetPolicyTagsQueryParams) (*PolicyTagListResult, *Response, error) {

	path := policyBasePath + "/policy/tag"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyTagListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyTagsCount is ...
// getPolicyTagsCount
//
//
//
//
// * @return *CountResult
func (s *PolicyService) GetPolicyTagsCount() (*CountResult, *Response, error) {

	path := policyBasePath + "/policy/tag/count"
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

// GetPolicyVersionNumbersQueryParams is ...
type GetPolicyVersionNumbersQueryParams struct {
	PolicyScope          string `url:"policyScope,omitempty"`          // Retrieve policy version numbers for the given scope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Retrieve policy version numbers for the given scopeWirelessSegment
	Offset               string `url:"offset,omitempty"`               // Starting index of the resources (1 based)
	Limit                string `url:"limit,omitempty"`                // Number of resources to return
}

// GetPolicyVersionNumbers is ...
// Retrieves the policy version numbers
//
//  * @param queryParams
//
//
// * @return *PolicyVersionNumberListResult
func (s *PolicyService) GetPolicyVersionNumbers(queryParams *GetPolicyVersionNumbersQueryParams) (*PolicyVersionNumberListResult, *Response, error) {

	path := policyBasePath + "/policy/version"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PolicyVersionNumberListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetPolicyVersionNumbersCountQueryParams is ...
type GetPolicyVersionNumbersCountQueryParams struct {
	PolicyScope          string `url:"policyScope,omitempty"`          // Retrieve the count of policy version numbers for the given scope
	ScopeWirelessSegment string `url:"scopeWirelessSegment,omitempty"` // Retrieve the count of policy version numbers for the given scopeWirelessSegment
}

// GetPolicyVersionNumbersCount is ...
// Retrieves the count of policy version numbers
//
//  * @param queryParams
//
//
// * @return *CountResult
func (s *PolicyService) GetPolicyVersionNumbersCount(queryParams *GetPolicyVersionNumbersCountQueryParams) (*CountResult, *Response, error) {

	path := policyBasePath + "/policy/version/count"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// UpdateQueryParams is ...
type UpdateQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the policy should be scheduled (Optional)
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // Custom Description (Optional)
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // Originator of this call (Optional)
}

// Update is ...
// Update a policy
//
//  * @param queryParams
//
//  * @param policyList Policy Object
// * @return *TaskIDResult
func (s *PolicyService) Update(policyList []*Policy, queryParams *UpdateQueryParams) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("PUT", path, policyList)
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

// UpdateFlowPolicySetting is ...
// updateFlowPolicySetting
//
//
//
//  * @param policyFlowSettingDTO policyFlowSettingDTO
// * @return *TaskIDResult
func (s *PolicyService) UpdateFlowPolicySetting(policyFlowSettingDTO PolicyFlowSettingDto) (*TaskIDResult, *Response, error) {

	path := policyBasePath + "/policy/flow/setting"
	req, err := s.client.NewRequest("PUT", path, policyFlowSettingDTO)
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
