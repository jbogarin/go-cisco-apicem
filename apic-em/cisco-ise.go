package apicem

import (
	"fmt"
	"strings"
)

const ciscoISEBasePath = "v2"

// CiscoISEService is an interface with the Vciscoise API
type CiscoISEService service

// CiscoISEDto is ...
type CiscoISEDto struct {
	ID                       string `json:"id,omitempty"`                       // id
	Password                 string `json:"password,omitempty"`                 // password
	Description              string `json:"description,omitempty"`              // description
	Username                 string `json:"username,omitempty"`                 // username
	IPAddress                string `json:"ipAddress,omitempty"`                // ipAddress
	ServerState              string `json:"serverState,omitempty"`              // serverState
	SubscriberName           string `json:"subscriberName,omitempty"`           // subscriberName
	KeystoreFileID           string `json:"keystoreFileId,omitempty"`           // keystoreFileId
	KeystoreFilePassPhrase   string `json:"keystoreFilePassPhrase,omitempty"`   // keystoreFilePassPhrase
	TruststoreFileID         string `json:"truststoreFileId,omitempty"`         // truststoreFileId
	TruststoreFilePassPhrase string `json:"truststoreFilePassPhrase,omitempty"` // truststoreFilePassPhrase
}

// CiscoISEListResult is ...
type CiscoISEListResult struct {
	Response []CiscoISEDto `json:"response,omitempty"`
	Version  string        `json:"version,omitempty"`
}

// CiscoISEResult is ...
type CiscoISEResult struct {
	Response CiscoISEDto `json:"response,omitempty"`
	Version  string      `json:"version,omitempty"`
}

// IdentityAuthFileInfoDto is ...
type IdentityAuthFileInfoDto struct {
	FileID   string `json:"fileId,omitempty"`   // fileId
	FileType string `json:"fileType,omitempty"` // fileType
	ID       string `json:"id,omitempty"`       // id
	FileName string `json:"fileName,omitempty"` // fileName
}

// IdentityAuthFileInfoListResult is ...
type IdentityAuthFileInfoListResult struct {
	Response []IdentityAuthFileInfoDto `json:"response,omitempty"`
	Version  string                    `json:"version,omitempty"`
}

// AddCiscoISEConfig is ...
// This method is used to create the ISE Server Configuration
//
//
//
//  * @param ciscoISEDTO CiscoISEDTO Object
// * @return *TaskIDResult
func (s *CiscoISEService) AddCiscoISEConfig(ciscoISEDTO CiscoISEDto) (*TaskIDResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise"
	req, err := s.client.NewRequest("POST", path, ciscoISEDTO)
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

// AddKeystoreFileQueryParams is ...
type AddKeystoreFileQueryParams struct {
	fileType string `url:"fileType,omitempty"` // fileType
}

// AddKeystoreFile is ...
// This method is used to upload Identity Authentication file info
//
//  * @param queryParams
//
//
// * @return *TaskIDResult
func (s *CiscoISEService) AddKeystoreFile(queryParams *AddKeystoreFileQueryParams) (*TaskIDResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/auth-file-info"
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

// Delete is ...
// This method is used to remove the CISCO ISE configuration for a specific ISE Server IP address
//
//
//
//
// * @return *TaskIDResult
func (s *CiscoISEService) Delete(ipAddress string) (*TaskIDResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)
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

// DeleteAuthFileInfo is ...
// This method is used to delete the Identity Auth File that was uploaded using an ID
//
//
//
//
// * @return *TaskIDResult
func (s *CiscoISEService) DeleteAuthFileInfo(id string) (*TaskIDResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/auth-file-info/{id}"
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

// GetAuthFileCount is ...
// This method is used to obtain the number of Identity Auth Files uploaded to APIC-EM
//
//
//
//
// * @return *CountResult
func (s *CiscoISEService) GetAuthFileCount() (*CountResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/auth-file-info/count"
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

// GetAuthFileList is ...
// This method is used to obtain the list of Keystore and Truststore files uploaded for ISE server configuration
//
//
//
//
// * @return *IdentityAuthFileInfoListResult
func (s *CiscoISEService) GetAuthFileList() (*IdentityAuthFileInfoListResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/auth-file-info"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(IdentityAuthFileInfoListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetByIPAddress is ...
// This method is used to get the CISCO ISE configuration for a specific ISE Server IP address
//
//
//
//
// * @return *CiscoISEResult
func (s *CiscoISEService) GetByIPAddress(ipAddress string) (*CiscoISEResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/ip/{ipAddress}/"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CiscoISEResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCount is ...
// This method is used to obtain the number of Cisco ISE configurations known to APIC-EM
//
//
//
//
// * @return *CountResult
func (s *CiscoISEService) GetCount() (*CountResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/count"
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

// GetEndPointGroupbyID is ...
// Get Cisco ISE configuration based on id
//
//
//
//
// * @return *CiscoISEResult
func (s *CiscoISEService) GetEndPointGroupbyID(id string) (*CiscoISEResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CiscoISEResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetList is ...
// This method is used to obtain the list of Cisco ISE Server Configurations
//
//
//
//
// * @return *CiscoISEListResult
func (s *CiscoISEService) GetList() (*CiscoISEListResult, *Response, error) {

	path := ciscoISEBasePath + "/cisco-ise"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CiscoISEListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
