package apicem

import (
	"fmt"
	"strings"
)

const pkiBrokerBasePath = "v1"

// PKIBrokerService is an interface with the  API
type PKIBrokerService service

// PkiTrustPointConfig is ...
type PkiTrustPointConfig struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`
	TrustProfileName string   `json:"trustProfileName,omitempty"`
	PlatformID       string   `json:"platformId,omitempty"`
	IosCli           []string `json:"iosCli,omitempty"`
	Pkcs12Password   string   `json:"pkcs12Password,omitempty"`
	ProvisionType    string   `json:"provisionType,omitempty"`
	SdnIP            string   `json:"sdnIp,omitempty"`
	CaFingerprint    string   `json:"caFingerprint,omitempty"`
	EnrollThreshold  string   `json:"enrollThreshold,omitempty"`
	IosSecureCli     []string `json:"iosSecureCli,omitempty"`
	Pkcs12URL        string   `json:"pkcs12Url,omitempty"`
	RsaKeySize       string   `json:"rsaKeySize,omitempty"`
	Fqdn             string   `json:"fqdn,omitempty"`
	EnrollSubjectDN  string   `json:"enrollSubjectDN,omitempty"`
	EnrollPort       string   `json:"enrollPort,omitempty"`
	EnrollURL        string   `json:"enrollUrl,omitempty"`
	EnrollPassword   string   `json:"enrollPassword,omitempty"`
	ID               string   `json:"id,omitempty"`
}

// PkiTrustPointConfigResult is ...
type PkiTrustPointConfigResult struct {
	Response PkiTrustPointConfig `json:"response,omitempty"`
	Version  string              `json:"version,omitempty"`
}

// PkiTrustPoint is ...
type PkiTrustPoint struct {
	EntityName             string `json:"entityName,omitempty"`             // Devices hostname
	SerialNumber           string `json:"serialNumber,omitempty"`           // Devices serial-number
	ID                     string `json:"id,omitempty"`                     // Trust-point identification. Automatically generated
	TrustProfileName       string `json:"trustProfileName,omitempty"`       // Name of trust-profile (must already exist). Default: sdn-network-infra-iwan
	EntityType             string `json:"entityType,omitempty"`             // Available options: router, switch. Currently not used
	NetworkDeviceID        string `json:"networkDeviceId,omitempty"`        // Device identification. Currently not used
	CertificateAuthorityID string `json:"certificateAuthorityId,omitempty"` // CA identification. Automatically populated
	PlatformID             string `json:"platformId,omitempty"`             // Platform identification. Eg. ASR1006
	ControllerIPAddress    string `json:"controllerIpAddress,omitempty"`    // IP address device uses to connect to APIC-EM. Eg. Proxy server IP address. Automatically populated if not set
	AttributeInfo          string `json:"attributeInfo,omitempty"`
}

// PkiTrustPointListResult is ...
type PkiTrustPointListResult struct {
	Response []PkiTrustPoint `json:"response,omitempty"`
	Version  string          `json:"version,omitempty"`
}

// PkiTrustPointResult is ...
type PkiTrustPointResult struct {
	Response PkiTrustPoint `json:"response,omitempty"`
	Version  string        `json:"version,omitempty"`
}

// TrustpoolUpdateParam is ...
type TrustpoolUpdateParam struct {
	Simulate bool   `json:"simulate,omitempty"`
	ID       string `json:"id,omitempty"`
}

// TrustpoolUpdateStatus is ...
type TrustpoolUpdateStatus struct {
	Update        string `json:"update,omitempty"`
	AttributeInfo string `json:"attributeInfo,omitempty"`
	ID            string `json:"id,omitempty"`
}

// TrustpoolUpdateStatusResult is ...
type TrustpoolUpdateStatusResult struct {
	Response TrustpoolUpdateStatus `json:"response,omitempty"`
	Version  string                `json:"version,omitempty"`
}

// CheckPKCS12Downloaded is ...
// This method is used to check if a specific trust-point has downloaded its PKCS#12 bundle
//
//
//  * @param scope Authorization Scope for RBAC * @param username requestUsername
//
// * @return *SuccessResult
func (s *PKIBrokerService) CheckPKCS12Downloaded(trustPointID string, scope string, username string) (*SuccessResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{trustPointID}/downloaded"
	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)
	req.Header.Add("username", username)

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetDefaultCaPem is ...
// This method is used to download the default certificate
//
//
//
//
// * @return void
// func (s *PKIBrokerService) GetDefaultCaPem(id string, pkiType string) (*Response, error) {

// 	path := pkiBrokerBasePath + "/certificate-authority/ca/{id}/{type}"
// 	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
// 	path = strings.Replace(path, "{"+"type"+"}", fmt.Sprintf("%v", pkiType), -1)
// 	req, err := s.client.NewRequest("GET", path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	resp, err := s.client.Do(req, root)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return root, resp, err
// }

// LegacyUpdateDefaultCaPem is ...
// This method is used to update the default certificate
//
//
//
//  * @param param param
// * @return *TrustpoolUpdateStatusResult
func (s *PKIBrokerService) LegacyUpdateDefaultCaPem(id string, pkiType string, param TrustpoolUpdateParam) (*TrustpoolUpdateStatusResult, *Response, error) {

	path := pkiBrokerBasePath + "/certificate-authority/update/{id}/{type}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"type"+"}", fmt.Sprintf("%v", pkiType), -1)
	req, err := s.client.NewRequest("PUT", path, param)
	if err != nil {
		return nil, nil, err
	}

	root := new(TrustpoolUpdateStatusResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointConfigGet is ...
// This method is used to obtain a specific trust-point&#39;s configuration
//
//
//
//
// * @return *PkiTrustPointConfigResult
func (s *PKIBrokerService) PkiTrustPointConfigGet(trustPointID string) (*PkiTrustPointConfigResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{trustPointID}/config"
	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PkiTrustPointConfigResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointCount is ...
// This method is used to return count of trust-points
//
//
//
//
// * @return *CountResult
func (s *PKIBrokerService) PkiTrustPointCount() (*CountResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/count"
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

// PkiTrustPointDelete is ...
// This method is used to delete a specific trust-point by its ID
//
//
//
//
// * @return *TaskIDResult
func (s *PKIBrokerService) PkiTrustPointDelete(trustPointID string) (*TaskIDResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{trustPointID}"
	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
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

// PkiTrustPointDeleteByDeviceSN is ...
// This method is used to delete a specific trust-point by its device serial-number
//
//
//
//
// * @return *TaskIDResult
func (s *PKIBrokerService) PkiTrustPointDeleteByDeviceSN(serialNumber string) (*TaskIDResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/serial-number/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)
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

// PkiTrustPointGet is ...
// This method is used to return a specific trust-point by its ID
//
//
//
//
// * @return *PkiTrustPointResult
func (s *PKIBrokerService) PkiTrustPointGet(trustPointID string) (*PkiTrustPointResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{trustPointID}"
	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PkiTrustPointResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointGetByDeviceSN is ...
// This method is used to return a specific trust-point by its device serial-number
//
//
//
//
// * @return *PkiTrustPointResult
func (s *PKIBrokerService) PkiTrustPointGetByDeviceSN(serialNumber string) (*PkiTrustPointResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/serial-number/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PkiTrustPointResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointListGet is ...
// This method is used to return a list of trust-points
//
//
//
//
// * @return *PkiTrustPointListResult
func (s *PKIBrokerService) PkiTrustPointListGet() (*PkiTrustPointListResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PkiTrustPointListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointListGetByRange is ...
// This method is used to return a list of trust-points by range
//
//
//
//
// * @return *PkiTrustPointListResult
func (s *PKIBrokerService) PkiTrustPointListGetByRange(startIndex int32, recordsToReturn int32) (*PkiTrustPointListResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PkiTrustPointListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// PkiTrustPointPkcs12Download is ...
// This method is used to download a trust-point&#39;s PKCS#12 bundle
//
//
//
//
// * @return void
// func (s *PKIBrokerService) PkiTrustPointPkcs12Download(trustPointID string, token string) (*Response, error) {

// 	path := pkiBrokerBasePath + "/trust-point/pkcs12/{trustPointID}/{token}"
// 	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
// 	path = strings.Replace(path, "{"+"token"+"}", fmt.Sprintf("%v", token), -1)
// 	req, err := s.client.NewRequest("GET", path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	resp, err := s.client.Do(req, root)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return root, resp, err
// }

// PkiTrustPointPost is ...
// This method is used to create a trust-point
//
//
//
//  * @param pkiTrustPointInput pkiTrustPointInput
// * @return *TaskIDResult
func (s *PKIBrokerService) PkiTrustPointPost(pkiTrustPointInput PkiTrustPoint) (*TaskIDResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point"
	req, err := s.client.NewRequest("POST", path, pkiTrustPointInput)
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

// PkiTrustPointPush is ...
// This method is used to push a created trust-point to its intended device
//
//
//
//
// * @return *TaskIDResult
func (s *PKIBrokerService) PkiTrustPointPush(trustPointID string) (*TaskIDResult, *Response, error) {

	path := pkiBrokerBasePath + "/trust-point/{trustPointID}"
	path = strings.Replace(path, "{"+"trustPointID"+"}", fmt.Sprintf("%v", trustPointID), -1)
	req, err := s.client.NewRequest("POST", path, nil)
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

// UpdateDefaultCaPem is ...
// This method is used to update the default certificate
//
//
//
//  * @param param param
// * @return *TrustpoolUpdateStatusResult
func (s *PKIBrokerService) UpdateDefaultCaPem(id string, pkiType string, param TrustpoolUpdateParam) (*TrustpoolUpdateStatusResult, *Response, error) {

	path := pkiBrokerBasePath + "/certificate-authority/{id}/{type}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"type"+"}", fmt.Sprintf("%v", pkiType), -1)
	req, err := s.client.NewRequest("PUT", path, param)
	if err != nil {
		return nil, nil, err
	}

	root := new(TrustpoolUpdateStatusResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
