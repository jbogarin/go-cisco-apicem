package apicem

import (
	"fmt"
	"strings"
)

const globalCredentialBasePath = "v1"

// GlobalCredentialService is an interface with the GlobalCredential API
type GlobalCredentialService service

// CliCredentialDTO is ...
type CliCredentialDTO struct {
	Username       string `json:"username,omitempty"`       // CLI username
	Password       string `json:"password,omitempty"`       // CLI password
	EnablePassword string `json:"enablePassword,omitempty"` // CLI enable password
	Description    string `json:"description,omitempty"`    // Description of the credential
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`
	ID             string `json:"id,omitempty"`
}

// GlobalCredentialDTO is ...
type GlobalCredentialDTO struct {
	Description    string `json:"description,omitempty"`    // Description of the credential
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`
	ID             string `json:"id,omitempty"`
}

// GlobalCredentialListResult is ...
type GlobalCredentialListResult struct {
	Version  string                `json:"version,omitempty"`
	Response []GlobalCredentialDTO `json:"response,omitempty"`
}

// GlobalCredentialSubTypeResult is ...
type GlobalCredentialSubTypeResult struct {
	Version  string `json:"version,omitempty"`
	Response string `json:"response,omitempty"`
}

// SNMPv2ReadCommunityDTO is ...
type SNMPv2ReadCommunityDTO struct {
	ReadCommunity  string `json:"readCommunity,omitempty"`  // SNMP read community
	Description    string `json:"description,omitempty"`    // Description of the credential
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`
	ID             string `json:"id,omitempty"`
}

// SNMPv2WriteCommunityDTO is ...
type SNMPv2WriteCommunityDTO struct {
	WriteCommunity string `json:"writeCommunity,omitempty"` // SNMP write community
	Description    string `json:"description,omitempty"`    // Description of the credential
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`
	ID             string `json:"id,omitempty"`
}

// SNMPv3CredentialDTO is ...
type SNMPv3CredentialDTO struct {
	Username        string `json:"username,omitempty"`        // SNMP user name
	AuthType        string `json:"authType,omitempty"`        // Authentication type is required if SNMP mode is AuthPriv / AuthNoPriv
	AuthPassword    string `json:"authPassword,omitempty"`    // AuthPassword is required if SNMP mode is AuthPriv / AuthNoPriv
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy password is required if SNMP mode is AuthPriv
	PrivacyType     string `json:"privacyType,omitempty"`     // Privacy type is required if SNMP mode is AuthPriv
	SnmpMode        string `json:"snmpMode,omitempty"`        // SNMP mode
	Description     string `json:"description,omitempty"`     // Description of the credential
	Comments        string `json:"comments,omitempty"`        // Comments to identify the credential
	CredentialType  string `json:"credentialType,omitempty"`  // Credential type to identify the application that uses the credential
	InstanceUUID    string `json:"instanceUuid,omitempty"`
	ID              string `json:"id,omitempty"`
}

// AddCliCredential is ...
// This method is used to add global CLI credential
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param globalCredentialNioList List of CLI credentials
// * @return *TaskIDResult
func (s *GlobalCredentialService) AddCliCredential(globalCredentialNioList *[]CliCredentialDTO, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/cli"
	req, err := s.client.NewRequest("POST", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// AddSnmpReadCommunity is ...
// This method is used to add global SNMP read community
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param globalCredentialNioList List of SNMP read communities
// * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpReadCommunity(globalCredentialNioList []SNMPv2ReadCommunityDTO, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv2-read-community"
	req, err := s.client.NewRequest("POST", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// AddSnmpWriteCommunity is ...
// This method is used to add global SNMP write community
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param globalCredentialNioList List of SNMP write communities
// * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpWriteCommunity(globalCredentialNioList []SNMPv2WriteCommunityDTO, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv2-write-community"
	req, err := s.client.NewRequest("POST", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// AddSnmpv3Credential is ...
// This method is used to add global SNMPv3 credential
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param globalCredentialNioList List of SNMPv3 credentials
// * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpv3Credential(globalCredentialNioList []SNMPv3CredentialDTO, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv3"
	req, err := s.client.NewRequest("POST", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteGlobalCredential is ...
// This method is used to delete global credential for the given ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *GlobalCredentialService) DeleteGlobalCredential(globalCredentialID string, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{"+"globalCredentialId"+"}", fmt.Sprintf("%v", globalCredentialID), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetGlobalCredentialQueryParams is ...
type GetGlobalCredentialQueryParams struct {
	credentialSubType string `url:"credentialSubType,omitempty"` // Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3
}

// GetGlobalCredential is ...
// This method is used to get global credential for the given credential sub type
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *GlobalCredentialListResult
func (s *GlobalCredentialService) GetGlobalCredential(scope string, queryParams *GetGlobalCredentialQueryParams) (*GlobalCredentialListResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(GlobalCredentialListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetGlobalCredentialSubTypeByID is ...
// This method is used to get credential sub type for the given ID
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *GlobalCredentialSubTypeResult
func (s *GlobalCredentialService) GetGlobalCredentialSubTypeByID(id string, scope string) (*GlobalCredentialSubTypeResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(GlobalCredentialSubTypeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
