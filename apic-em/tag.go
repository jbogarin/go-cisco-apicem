package apicem

import (
	"fmt"
	"strings"
)

const tagBasePath = "v1"

// TagService is an interface with the Tag API
type TagService service

// TagDto is ...
type TagDto struct {
	ID           string `json:"id,omitempty"`           // Unique identifier for tag
	Tag          string `json:"tag,omitempty"`          // Name of the tag
	ResourceType string `json:"resourceType,omitempty"` // Type of the resource to which the tag to be associated
	ResourceID   string `json:"resourceId,omitempty"`   // Id of the resource to which the tag to be associated
}

// TagDtoListResult is ...
type TagDtoListResult struct {
	Version  string   `json:"version,omitempty"`
	Response []TagDto `json:"response,omitempty"`
}

// TagDtoResult is ...
type TagDtoResult struct {
	Version  string `json:"version,omitempty"`
	Response TagDto `json:"response,omitempty"`
}

// AddTag is ...
// Adds a new tag to the controller. The {tag} field should contain the value of the tag.
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param tagDto TagDto with the tag
// * @return *TaskIDResult
func (s *TagService) AddTag(tagDto TagDto, scope string) (*TaskIDResult, *Response, error) {

	path := tagBasePath + "/tag"
	req, err := s.client.NewRequest("POST", path, tagDto)
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

// AddTagToResource is ...
// Associates a tag to a resource. The {id} field should be the id of the tag. The {resourceID} should the id of the resource. The {resourceType} should be the type of the resource. Supported resourceTypes are network-device, interface.
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param tagDto TagDto with tag ID, resource ID and resource type
// * @return *TaskIDResult
func (s *TagService) AddTagToResource(tagDto TagDto, scope string) (*TaskIDResult, *Response, error) {

	path := tagBasePath + "/tag/association"
	req, err := s.client.NewRequest("POST", path, tagDto)
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

// DeleteTag is ...
// Deletes the tag by its id. The tag cannot be deleted when there is an existing association between the tag and a resource.
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *TagService) DeleteTag(id string, scope string) (*TaskIDResult, *Response, error) {

	path := tagBasePath + "/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
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

// DeleteTagFromResourceQueryParams is ...
type DeleteTagFromResourceQueryParams struct {
	ResourceType string `url:"resourceType,omitempty"` // Type of resource. Supported resourceTypes are network-device, interface.
	ResourceID   string `url:"resourceId,omitempty"`   // Resource ID
}

// DeleteTagFromResource is ...
// Deletes the association of tag by its id.from resource of id {resourceID} which is of type {resourceType}.  Supported resourceTypes are network-device, interface.
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *TaskIDResult
func (s *TagService) DeleteTagFromResource(scope string, id string, queryParams *DeleteTagFromResourceQueryParams) (*TaskIDResult, *Response, error) {

	path := tagBasePath + "/tag/association/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
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

// GetTag is ...
// Gets the details of the tag by its id.
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *TagDtoResult
func (s *TagService) GetTag(id string, scope string) (*TagDtoResult, *Response, error) {

	path := tagBasePath + "/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TagDtoResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTagsQueryParams is ...
type GetTagsQueryParams struct {
	ResourceType string `url:"resourceType,omitempty"` // Type of resource (network-device)
	ResourceID   string `url:"resourceId,omitempty"`   // Resource ID
}

// GetTags is ...
// Gets all the tags if no filters are provided. Gets all the tags that are associated with resources of {resourceType} if resourceType filter is provided. Gets all the tags that are associated with a resource with id {resourceID} and of resource type {resourceType} when resourceId and resourceType filters are provided.
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *TagDtoListResult
func (s *TagService) GetTags(scope string, queryParams *GetTagsQueryParams) (*TagDtoListResult, *Response, error) {

	path := tagBasePath + "/tag"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(TagDtoListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateTag is ...
// Updates the tag identified by {id} to a new value. The {id} field should be id of the old tag. The {tag} field should contain the new value of the tag.
//
//
//  * @param scope Authorization Scope for RBAC
//  * @param tagDto TagDto with the new tag
// * @return *TaskIDResult
func (s *TagService) UpdateTag(tagDto TagDto, scope string) (*TaskIDResult, *Response, error) {

	path := tagBasePath + "/tag"
	req, err := s.client.NewRequest("PUT", path, tagDto)
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
