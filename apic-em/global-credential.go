package apicem

import (
	"fmt"
	"strings"
)

const globalCredentialBasePath = "v1"

// GlobalCredentialService is an interface with the Globalcredential API
type GlobalCredentialService service

// CliCredentialDto is ...
type CliCredentialDto struct {

	// CLI username
	Username string `json:"username,omitempty"`

	// CLI password
	Password string `json:"password,omitempty"`

	// CLI enable password
	EnablePassword string `json:"enablePassword,omitempty"`

	// Description of the credential
	Description string `json:"description,omitempty"`

	// Comments to identify the credential
	Comments string `json:"comments,omitempty"`

	// Credential type to identify the application that uses the credential
	CredentialType string `json:"credentialType,omitempty"`

	InstanceUUID string `json:"instanceUuid,omitempty"`

	ID string `json:"id,omitempty"`
}

// CredentialQueryParams is ...
type CredentialQueryParams struct {
	CredentialSubType string `url:"credentialSubType,omitempty"`
}

// The GlobalCredentialDto is ...
type GlobalCredentialDto struct {
	// Description of the credential
	Description string `json:"description,omitempty"`
	// Comments to identify the credential
	Comments string `json:"comments,omitempty"`
	// Credential type to identify the application that uses the credential
	CredentialType string `json:"credentialType,omitempty"`
	InstanceUUID   string `json:"instanceUuid,omitempty"`
	ID             string `json:"id,omitempty"`
}

// The GlobalCredentialListResult is ...
type GlobalCredentialListResult struct {
	Version string `json:"version,omitempty"`

	Response []GlobalCredentialDto `json:"response,omitempty"`
}

// The GlobalCredentialSubTypeResult is ...
type GlobalCredentialSubTypeResult struct {
	Version string `json:"version,omitempty"`

	Response string `json:"response,omitempty"`
}

// The SNMPv2ReadCommunityDto is ...
type SNMPv2ReadCommunityDto struct {

	// SNMP read community
	ReadCommunity string `json:"readCommunity,omitempty"`

	// Description of the credential
	Description string `json:"description,omitempty"`

	// Comments to identify the credential
	Comments string `json:"comments,omitempty"`

	// Credential type to identify the application that uses the credential
	CredentialType string `json:"credentialType,omitempty"`

	InstanceUUID string `json:"instanceUuid,omitempty"`

	ID string `json:"id,omitempty"`
}

// SNMPv2WriteCommunityDto is ...
type SNMPv2WriteCommunityDto struct {

	// SNMP write community
	WriteCommunity string `json:"writeCommunity,omitempty"`

	// Description of the credential
	Description string `json:"description,omitempty"`

	// Comments to identify the credential
	Comments string `json:"comments,omitempty"`

	// Credential type to identify the application that uses the credential
	CredentialType string `json:"credentialType,omitempty"`

	InstanceUUID string `json:"instanceUuid,omitempty"`

	ID string `json:"id,omitempty"`
}

// SNMPv3CredentialDto is ...
type SNMPv3CredentialDto struct {
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
//  * @param globalCredentialNioList List of CLI credentials
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *GlobalCredentialService) AddCliCredential(globalCredentialNioList *[]CliCredentialDto, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/cli"

	req, err := s.client.NewRequest("Post", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param globalCredentialNioList List of SNMP read communities
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpReadCommunity(globalCredentialNioList *[]SNMPv2ReadCommunityDto, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv2-read-community"

	req, err := s.client.NewRequest("Post", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param globalCredentialNioList List of SNMP write communities
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpWriteCommunity(globalCredentialNioList *[]SNMPv2WriteCommunityDto, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv2-write-community"

	req, err := s.client.NewRequest("Post", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param globalCredentialNioList List of SNMPv3 credentials
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *GlobalCredentialService) AddSnmpv3Credential(globalCredentialNioList *[]SNMPv3CredentialDto, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/snmpv3"

	req, err := s.client.NewRequest("Post", path, globalCredentialNioList)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
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
//  * @param globalCredentialID ID of global-credential
//  * @param scope Authorization Scope for RBAC
//  * @return *TaskIDResult
func (s *GlobalCredentialService) DeleteGlobalCredential(globalCredentialID string, scope string) (*TaskIDResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/{globalCredentialID}"
	path = strings.Replace(path, "{"+"globalCredentialID"+"}", fmt.Sprintf("%v", globalCredentialID), -1)

	req, err := s.client.NewRequest("Delete", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
	req.Header.Add("scope", scope)

	root := new(TaskIDResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetGlobalCredential is ...
// This method is used to get global credential for the given credential sub type
//
//  * @param scope Authorization Scope for RBAC
//  * @param credentialSubType Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3
//  * @return *GlobalCredentialListResult
func (s *GlobalCredentialService) GetGlobalCredential(scope string, credentialQueryParams *CredentialQueryParams) (*GlobalCredentialListResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential"
	path, err := addOptions(path, credentialQueryParams)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
	req.Header.Add("scope", scope)

	root := new(GlobalCredentialListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetGlobalCredentialSubTypeByID is ...
// This method is used to get credential sub type for the given Id
//
//  * @param id Global Credential ID
//  * @param scope Authorization Scope for RBAC
//  * @return *GlobalCredentialSubTypeResult
func (s *GlobalCredentialService) GetGlobalCredentialSubTypeByID(id string, scope string) (*GlobalCredentialSubTypeResult, *Response, error) {

	path := globalCredentialBasePath + "/global-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}
	// header params "scope"
	req.Header.Add("scope", scope)

	root := new(GlobalCredentialSubTypeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
