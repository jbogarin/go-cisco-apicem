package apicem

import "time"

const auditBasePath = "v1"

// AuditService is an interface with the Audit API
type AuditService service

// AugmentedAuditResourceDto is ... See: http://apic-em/wiki/Audit_Design
type AugmentedAuditResourceDto struct {
	InstanceUUID     string    `json:"instanceUuid,omitempty"` // This field is deprecated. Use 'id' instead.
	Severity         string    `json:"severity,omitempty"`
	Tag              string    `json:"tag,omitempty"`
	ApplicationName  string    `json:"applicationName,omitempty"`
	HasParent        bool      `json:"hasParent,omitempty"`
	CreatedDateTime  time.Time `json:"createdDateTime,omitempty"`
	AuditID          string    `json:"auditId,omitempty"`
	AuditParentID    string    `json:"auditParentId,omitempty"`
	AuditRequestor   string    `json:"auditRequestor,omitempty"`
	AuditDescription string    `json:"auditDescription,omitempty"`
	DerivedParentID  string    `json:"derivedParentId,omitempty"`
	DeviceIP         string    `json:"deviceIP,omitempty"`
	DeviceName       string    `json:"deviceName,omitempty"`
	SiteName         string    `json:"siteName,omitempty"`
	HasChildren      bool      `json:"hasChildren,omitempty"`
	AuditParameters  string    `json:"auditParameters,omitempty"`
	PersistDateTime  time.Time `json:"persistDateTime,omitempty"`
}

// ListAuditResourceDtoResponse is ...
type ListAuditResourceDtoResponse struct {
	Version  string                      `json:"version,omitempty"`
	Response []AugmentedAuditResourceDto `json:"response,omitempty"`
}

// DownloadAuditLogs is ...
//
//
//
//
// * @return *TaskIDResult
func (s *AuditService) DownloadAuditLogs() (*TaskIDResult, *Response, error) {

	path := auditBasePath + "/audit/download"
	req, err := s.client.NewRequest("GET", path, nil)
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

// GetAuditCountWithFilter is ...
// Same parameters as search should be passed.. &lt;br/&gt;For example: &#39;/audit/count?auditRecordStartTime&#x3D;10&amp;auditRecordEndTime&#x3D;1529152597971&amp;auditRequestor&#x3D;123a1&amp;auditRequestor&#x3D;some_other_value&#39;. &lt;br/&gt;Request parameter is case sensitive, so &#39;/audit/count?device&lt;b&gt;IP&lt;/b&gt;&#x3D;10.1.1.0&#39; will search for deviceIP equals to 10.1.1.0. &#39;/audit?device&lt;b&gt;ip&lt;/b&gt;&#x3D;10.1.1.0&#39; will be ignored. &lt;br/&gt;Wildcards are supported using %25 as a wildcard, for example: &#39;/audit/count?auditDescription&#x3D;%25success%25&#39; will retrieve all records matching success, successful, successive, etc.&lt;br/&gt;Request values are not case sensitive, so &#39;/audit/count?auditRequestor&#x3D;ADMIN&#39; will retrieve all records with auditRequestor of &#39;admin&#39;, &#39;Admin&#39;, &#39;ADMIN&#39;, etc.
//
//
//
//
// * @return *ListAuditResourceDtoResponse
func (s *AuditService) GetAuditCountWithFilter() (*ListAuditResourceDtoResponse, *Response, error) {

	path := auditBasePath + "/audit/count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ListAuditResourceDtoResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAuditWithFilter is ...
// Free search on any atomic attributes passed as request parameter. &lt;br/&gt;For example: &#39;/audit?auditRecordStartTime&#x3D;10&amp;auditRecordEndTime&#x3D;1529152597971&amp;auditRequestor&#x3D;123a1&amp;auditRequestor&#x3D;some_other_value&#39;. &lt;br/&gt;Request parameter is case sensitive, so &#39;/audit?device&lt;b&gt;IP&lt;/b&gt;&#x3D;10.1.1.0&#39; will search for deviceIP equals to 10.1.1.0. &#39;/audit?device&lt;b&gt;ip&lt;/b&gt;&#x3D;10.1.1.0&#39; will be ignored. &lt;br/&gt;Operation supports offset, limit: &#39;/audit?limit&#x3D;100&amp;offset&#x3D;40&#39; will retrieve 100 records starts from position 40 &lt;br/&gt;Wildcards are supported using %25 as a wildcard, for example: &#39;/audit?auditDescription&#x3D;%25success%25&#39; will retrieve all records matching success, successful, successive, etc.&lt;br/&gt;Request values are not case sensitive, so &#39;/audit?auditRequestor&#x3D;ADMIN&#39; will retrieve all records with auditRequestor of &#39;admin&#39;, &#39;Admin&#39;, &#39;ADMIN&#39;, etc.
//
//
//
//
// * @return *ListAuditResourceDtoResponse
func (s *AuditService) GetAuditWithFilter() (*ListAuditResourceDtoResponse, *Response, error) {

	path := auditBasePath + "/audit"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ListAuditResourceDtoResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
