package apicem

import (
	"fmt"
	"strings"
)

const licenseBasePath = "v1"

// LicenseService is an interface with the  API
type LicenseService service

// LicenseInfoDto is ...
type LicenseInfoDto struct {
	Description             string `json:"description,omitempty"`             // Description about the license
	Name                    string `json:"name,omitempty"`                    // Name of the license feature
	Priority                string `json:"priority,omitempty"`                // License priority
	Type                    string `json:"type,omitempty"`                    // License type
	Status                  string `json:"status,omitempty"`                  // Status of the license
	DeployPending           int32  `json:"deployPending,omitempty"`           // Deploy Pending information of license
	TotalCount              int32  `json:"totalCount,omitempty"`              // Number of license installed in the device
	ParentID                int32  `json:"parentId,omitempty"`                // Parent Id of the license
	ProvisionState          int32  `json:"provisionState,omitempty"`          // Provision state of the license feature
	PhysicalIndex           string `json:"physicalIndex,omitempty"`           // Physical entity index
	LicenseIndex            int32  `json:"licenseIndex,omitempty"`            // Index of the license
	FeatureVersion          string `json:"featureVersion,omitempty"`          // Version of the license feature
	Counted                 bool   `json:"counted,omitempty"`                 // If license feature is counted as part of the license
	ValidityPeriod          int64  `json:"validityPeriod,omitempty"`          // Validity period the the license
	ValidityPeriodRemaining int64  `json:"validityPeriodRemaining,omitempty"` // Remaining validityPeriod
	MaxUsageCount           int64  `json:"maxUsageCount,omitempty"`           // Maximum usage of the license feature
	UsageCount              int32  `json:"usageCount,omitempty"`              // Usage count of the license feature
	EulaStatus              bool   `json:"eulaStatus,omitempty"`              // EULA status of the license
	UsageCountRemaining     int64  `json:"usageCountRemaining,omitempty"`     // Unused license count
	ExpiredPeriod           int32  `json:"expiredPeriod,omitempty"`           // License expired period
	EvalPeriodLeft          string `json:"evalPeriodLeft,omitempty"`          // Number of days remaining in the eval period
	EvalPeriodUsed          string `json:"evalPeriodUsed,omitempty"`          // Number of days used in the eval period
	ExpiredDate             string `json:"expiredDate,omitempty"`             // Expired date of the license
	IsCounted               bool   `json:"isCounted,omitempty"`               // If the license is counted as part of license usage
	IsEulaAccepted          bool   `json:"isEulaAccepted,omitempty"`          // If the EULA is accepted
	IsEulaApplicable        bool   `json:"isEulaApplicable,omitempty"`        // If the EULA is applicable
	IsTechnologyLicense     bool   `json:"isTechnologyLicense,omitempty"`     // If the license is technology license or not
	LicenseFileCount        int32  `json:"licenseFileCount,omitempty"`        // Number of license file
	LicenseFileName         string `json:"licenseFileName,omitempty"`         // Name of license file
	StoreName               string `json:"storeName,omitempty"`               // Name of the license Store
	StoredUsed              int32  `json:"storedUsed,omitempty"`              // License store usage detail
	HostID                  string `json:"hostId,omitempty"`                  // Device Id/Name of the license info
	AttributeInfo           string `json:"attributeInfo,omitempty"`
}

// LicenseInfoListResult is ...
type LicenseInfoListResult struct {
	Version  string           `json:"version,omitempty"`
	Response []LicenseInfoDto `json:"response,omitempty"`
}

// SuccessResultList is ...
type SuccessResultList struct {
	Version  string   `json:"version,omitempty"`
	Response []string `json:"response,omitempty"`
}

// GetDeviceIDByFilename is ...
// Gets the list of devices with licensefilename
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *SuccessResultList
func (s *LicenseService) GetDeviceIDByFilename(licenseFileName string, scope string) (*SuccessResultList, *Response, error) {

	path := licenseBasePath + "/network-device/license/{licenseFileName}"
	path = strings.Replace(path, "{"+"licenseFileName"+"}", fmt.Sprintf("%v", licenseFileName), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(SuccessResultList)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLicenseInfoQueryParams is ...
type GetLicenseInfoQueryParams struct {
	Limit                       string   `url:"limit,omitempty"`                       // limit
	Offset                      string   `url:"offset,omitempty"`                      // offset
	SortBy                      string   `url:"sortBy,omitempty"`                      // sortBy
	Order                       string   `url:"order,omitempty"`                       // order
	CountedList                 []string `url:"countedList,omitempty"`                 // countedList
	DeployPendingList           []string `url:"deployPendingList,omitempty"`           // deployPendingList
	EULAStatusList              []string `url:"eulaStatusList,omitempty"`              // eulaStatusList
	EvalPeriodLeftList          []string `url:"evalPeriodLeftList,omitempty"`          // evalPeriodLeftList
	EvalPeriodUsedList          []string `url:"evalPeriodUsedList,omitempty"`          // evalPeriodUsedList
	ExpiredDateList             []string `url:"expiredDateList,omitempty"`             // expiredDateList
	ExpiredPeriodList           []string `url:"expiredPeriodList,omitempty"`           // expiredPeriodList
	FeatureVersionList          []string `url:"featureVersionList,omitempty"`          // featureVersionList
	HostIDList                  []string `url:"hostIdList,omitempty"`                  // hostIdList
	IsCountedList               []string `url:"isCountedList,omitempty"`               // isCountedList
	IsEulaAcceptedList          []string `url:"isEulaAcceptedList,omitempty"`          // isEulaAcceptedList
	IsEulaApplicableList        []string `url:"isEulaApplicableList,omitempty"`        // isEulaApplicableList
	IsTechnologyLicenseList     []string `url:"isTechnologyLicenseList,omitempty"`     // isTechnologyLicenseList
	LicenseFileCountList        []string `url:"licenseFileCountList,omitempty"`        // licenseFileCountList
	LicenseFileNameList         []string `url:"licenseFileNameList,omitempty"`         // licenseFileNameList
	LicenseIndexList            []string `url:"licenseIndexList,omitempty"`            // licenseIndexList
	MaxUsageCountList           []string `url:"maxUsageCountList,omitempty"`           // maxUsageCountList
	ParentIDList                []string `url:"parentIdList,omitempty"`                // parentIdList
	PhysicalIndexList           []string `url:"physicalIndexList,omitempty"`           // physicalIndexList
	PriorityList                []string `url:"priorityList,omitempty"`                // priorityList
	ProvisionStateList          []string `url:"provisionStateList,omitempty"`          // provisionStateList
	StatusList                  []string `url:"statusList,omitempty"`                  // statusList
	StoredUsedList              []string `url:"storedUsedList,omitempty"`              // storedUsedList
	StoreNameList               []string `url:"storeNameList,omitempty"`               // storeNameList
	TotalCountList              []string `url:"totalCountList,omitempty"`              // totalCountList
	LicenseTypeList             []string `url:"licenseTypeList,omitempty"`             // licenseTypeList
	UsageCountList              []string `url:"usageCountList,omitempty"`              // usageCountList
	UsageCountRemainingList     []string `url:"usageCountRemainingList,omitempty"`     // usageCountRemainingList
	ValidityPeriodList          []string `url:"validityPeriodList,omitempty"`          // validityPeriodList
	ValidityPeriodRemainingList []string `url:"validityPeriodRemainingList,omitempty"` // validityPeriodRemainingList
}

// GetLicenseInfo is ...
// Gets the list of licenses for network device with given filters
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *LicenseInfoListResult
func (s *LicenseService) GetLicenseInfo(scope string, deviceID string, queryParams *GetLicenseInfoQueryParams) (*LicenseInfoListResult, *Response, error) {

	path := licenseBasePath + "/license-info/network-device/{deviceId}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(LicenseInfoListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLicenseInfoCountQueryParams is ...
type GetLicenseInfoCountQueryParams struct {
	CountedList                 []string `url:"countedList,omitempty"`                 // countedList
	DeployPendingList           []string `url:"deployPendingList,omitempty"`           // deployPendingList
	EulaStatusList              []string `url:"eulaStatusList,omitempty"`              // eulaStatusList
	EvalPeriodLeftList          []string `url:"evalPeriodLeftList,omitempty"`          // evalPeriodLeftList
	EvalPeriodUsedList          []string `url:"evalPeriodUsedList,omitempty"`          // evalPeriodUsedList
	ExpiredDateList             []string `url:"expiredDateList,omitempty"`             // expiredDateList
	ExpiredPeriodList           []string `url:"expiredPeriodList,omitempty"`           // expiredPeriodList
	FeatureVersionList          []string `url:"featureVersionList,omitempty"`          // featureVersionList
	HostIDList                  []string `url:"hostIdList,omitempty"`                  // hostIdList
	IsCountedList               []string `url:"isCountedList,omitempty"`               // isCountedList
	IsEulaAcceptedList          []string `url:"isEulaAcceptedList,omitempty"`          // isEulaAcceptedList
	IsEulaApplicableList        []string `url:"isEulaApplicableList,omitempty"`        // isEulaApplicableList
	IsTechnologyLicenseList     []string `url:"isTechnologyLicenseList,omitempty"`     // isTechnologyLicenseList
	LicenseFileCountList        []string `url:"licenseFileCountList,omitempty"`        // licenseFileCountList
	LicenseFileNameList         []string `url:"licenseFileNameList,omitempty"`         // licenseFileNameList
	LicenseIndexList            []string `url:"licenseIndexList,omitempty"`            // licenseIndexList
	MaxUsageCountList           []string `url:"maxUsageCountList,omitempty"`           // maxUsageCountList
	ParentIDList                []string `url:"parentIdList,omitempty"`                // parentIdList
	PhysicalIndexList           []string `url:"physicalIndexList,omitempty"`           // physicalIndexList
	PriorityList                []string `url:"priorityList,omitempty"`                // priorityList
	ProvisionStateList          []string `url:"provisionStateList,omitempty"`          // provisionStateList
	StatusList                  []string `url:"statusList,omitempty"`                  // statusList
	StoredUsedList              []string `url:"storedUsedList,omitempty"`              // storedUsedList
	StoreNameList               []string `url:"storeNameList,omitempty"`               // storeNameList
	TotalCountList              []string `url:"totalCountList,omitempty"`              // totalCountList
	LicenseTypeList             []string `url:"licenseTypeList,omitempty"`             // licenseTypeList
	UsageCountList              []string `url:"usageCountList,omitempty"`              // usageCountList
	UsageCountRemainingList     []string `url:"usageCountRemainingList,omitempty"`     // usageCountRemainingList
	ValidityPeriodList          []string `url:"validityPeriodList,omitempty"`          // validityPeriodList
	ValidityPeriodRemainingList []string `url:"validityPeriodRemainingList,omitempty"` // validityPeriodRemainingList
}

// GetLicenseInfoCount is ...
// Gets the list of licenses for network device with given filters
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *CountResult
func (s *LicenseService) GetLicenseInfoCount(scope string, deviceID string, queryParams *GetLicenseInfoCountQueryParams) (*CountResult, *Response, error) {

	path := licenseBasePath + "/license-info/network-device/{deviceId}/count"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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
