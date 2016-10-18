package apicem

import (
	"fmt"
	"os"
	"strings"
)

const certificateManagementBasePath = "v1"

// CertificateManagementService is an interface with the  API
type CertificateManagementService service

// CertificateBrief is ...
type CertificateBrief struct {
	CommonName    string `json:"commonName,omitempty"`   // Certificate common name
	Issuer        string `json:"issuer,omitempty"`       // Certificate issuer
	SerialNumber  string `json:"serialNumber,omitempty"` // Certificate serial-number
	Expiry        string `json:"expiry,omitempty"`       // Certificate expiry
	GvSerialID    string `json:"gvSerialId,omitempty"`   // Grapevine certificate serial identification
	SelfSigned    string `json:"selfSigned,omitempty"`   // Set if this is a self-signed certificate
	ProxyEnabled  string `json:"proxyEnabled,omitempty"` // Set if this is a proxy certificate
	AttributeInfo string `json:"attributeInfo,omitempty"`
	ID            string `json:"id,omitempty"`
}

// CertificateBriefListResult is ...
type CertificateBriefListResult struct {
	Response []CertificateBrief `json:"response,omitempty"`
	Version  string             `json:"version,omitempty"`
}

// CertificateBriefResult is ...
type CertificateBriefResult struct {
	Response CertificateBrief `json:"response,omitempty"`
	Version  string           `json:"version,omitempty"`
}

// CertificateProxyStatus is ...
type CertificateProxyStatus struct {
	ProxyEnabled bool `json:"proxyEnabled,omitempty"`
}

// GetCertificateBrief is ...
// This method is used to return a specific certificte by its ID
//
//
//
//
// * @return *CertificateBriefResult
func (s *CertificateManagementService) GetCertificateBrief(certificateID string) (*CertificateBriefResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate/{certificateId}"
	path = strings.Replace(path, "{"+"certificateId"+"}", fmt.Sprintf("%v", certificateID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CertificateBriefResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCertificateBriefList is ...
// This method is used to return list of certificates
//
//
//
//
// * @return *CertificateBriefListResult
func (s *CertificateManagementService) GetCertificateBriefList() (*CertificateBriefListResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CertificateBriefListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCertificateBriefListByRange is ...
// This method is used to return list of certificates by range
//
//
//
//
// * @return *CertificateBriefListResult
func (s *CertificateManagementService) GetCertificateBriefListByRange(startIndex int32, recordsToReturn int32) (*CertificateBriefListResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CertificateBriefListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCertificateCount is ...
// This method is used to return count of certificates
//
//
//
//
// * @return *CountResult
func (s *CertificateManagementService) GetCertificateCount() (*CountResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate/count"
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

// GetProxyCertificate is ...
// This method is used to return the proxy certificate
//
//
//
//
// * @return *CertificateBriefResult
func (s *CertificateManagementService) GetProxyCertificate() (*CertificateBriefResult, *Response, error) {

	path := certificateManagementBasePath + "/proxy-certificate"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CertificateBriefResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// ImportCertificateQueryParams is ...
type ImportCertificateQueryParams struct {
	pkPassword string `url:"pkPassword,omitempty"` // Private Key Passsword
}

// ImportCertificate is ...
// This method is used to upload a certificate
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *CertificateManagementService) ImportCertificate(queryParams *ImportCertificateQueryParams) (*TaskIDResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// ImportCertificateP12QueryParams is ...
type ImportCertificateP12QueryParams struct {
	p12Password string `url:"p12Password,omitempty"` // P12 Passsword
	pkPassword  string `url:"pkPassword,omitempty"`  // Private Key Passsword
}

// ImportCertificateP12 is ...
// This method is used to upload a PKCS#12 file
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *CertificateManagementService) ImportCertificateP12(queryParams *ImportCertificateP12QueryParams) (*TaskIDResult, *Response, error) {

	path := certificateManagementBasePath + "/certificate-p12"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// PostProxyCertificate is ...
// This method is used to upload the proxy certificate
//
//
//
//
// * @return *TaskIDResult
func (s *CertificateManagementService) PostProxyCertificate(certFileUpload *os.File) (*TaskIDResult, *Response, error) {

	path := certificateManagementBasePath + "/proxy-certificate"
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

// UpdateProxyCertState is ...
// This method is used to update the proxy certificate
//
//
//
//  * @param certificateProxyStatus Enable/Disable proxy certificate
// * @return *TaskIDResult
func (s *CertificateManagementService) UpdateProxyCertState(certificateProxyStatus CertificateProxyStatus) (*TaskIDResult, *Response, error) {

	path := certificateManagementBasePath + "/proxy-certificate"
	req, err := s.client.NewRequest("PUT", path, certificateProxyStatus)
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
