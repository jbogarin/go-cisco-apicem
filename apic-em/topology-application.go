package apicem

import (
	"fmt"
	"strings"
)

const topologyApplicationBasePath = "v1"

// TopologyApplicationService is an interface with the Topology API
type TopologyApplicationService service

// TopologyApplicationDto is ...
type TopologyApplicationDto struct {
	Name        string `json:"name,omitempty"`        // Name for this Application
	ID          string `json:"id,omitempty"`          // Unique identifier for this Application
	Description string `json:"description,omitempty"` // Description for this Application
}

// TopologyApplicationListResult is ...
type TopologyApplicationListResult struct {
	Version  string                   `json:"version,omitempty"`
	Response []TopologyApplicationDto `json:"response,omitempty"`
}

// TopologyApplicationResult is ...
type TopologyApplicationResult struct {
	Version  string                 `json:"version,omitempty"`
	Response TopologyApplicationDto `json:"response,omitempty"`
}

// TopologyPageDto is ...
type TopologyPageDto struct {
	Description     string `json:"description,omitempty"`     // Description for this Page
	Name            string `json:"name,omitempty"`            // Name for this Page
	ID              string `json:"id,omitempty"`              // Unique identifier for this Page
	ApplicationUUID string `json:"applicationUuid,omitempty"` // Application unique identifier for this Page
	DefaultViewID   string `json:"defaultViewId,omitempty"`   // Default View unique identifier for this Page
}

// TopologyPageListResult is ...
type TopologyPageListResult struct {
	Version  string            `json:"version,omitempty"`
	Response []TopologyPageDto `json:"response,omitempty"`
}

// TopologyPageResult is ...
type TopologyPageResult struct {
	Version  string          `json:"version,omitempty"`
	Response TopologyPageDto `json:"response,omitempty"`
}

// TopologyViewDto is ...
type TopologyViewDto struct {
	Description     string   `json:"description,omitempty"`     // View description
	Name            string   `json:"name,omitempty"`            // View name
	ID              string   `json:"id,omitempty"`              // Unique Identifier for View
	Topology        Topology `json:"topology,omitempty"`        // Topology being represented by this view
	ApplicationUUID string   `json:"applicationUuid,omitempty"` // Application unique identifier for this view
	PageUUID        string   `json:"pageUuid,omitempty"`        // Page unique identifier for this view inside the corresponding application
}

// TopologyViewListResult is ...
type TopologyViewListResult struct {
	Version  string            `json:"version,omitempty"`
	Response []TopologyViewDto `json:"response,omitempty"`
}

// TopologyViewResult is ...
type TopologyViewResult struct {
	Version  string          `json:"version,omitempty"`
	Response TopologyViewDto `json:"response,omitempty"`
}

// CreateTopologyApplications is ...
// This method is used to create topology applications.
//
//
//
//  * @param topologyApplicationDtoList application
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyApplications(topologyApplicationDtoList *[]TopologyApplicationDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application"
	req, err := s.client.NewRequest("POST", path, topologyApplicationDtoList)
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

// CreateTopologyApplications1 is ...
// This method is used to create topology applications.
//
//
//
//  * @param topologyApplicationDtoList application
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyApplications1(topologyApplicationDtoList *[]TopologyApplicationDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp"
	req, err := s.client.NewRequest("POST", path, topologyApplicationDtoList)
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

// CreateTopologyPages is ...
// This method is used to create topology pages.
//
//
//
//  * @param topologyPageDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyPages(applicationUUID string, topologyPageDtoList *[]TopologyPageDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("POST", path, topologyPageDtoList)
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

// CreateTopologyPages2 is ...
// This method is used to create topology pages.
//
//
//
//  * @param topologyPageDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyPages2(applicationUUID string, topologyPageDtoList *[]TopologyPageDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("POST", path, topologyPageDtoList)
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

// CreateTopologyViews is ...
// This method is used to create topology views.
//
//
//
//  * @param topologyViewDtoList view
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyViews(applicationUUID string, pageUUID string, topologyViewDtoList *[]TopologyViewDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("POST", path, topologyViewDtoList)
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

// CreateTopologyViews3 is ...
// This method is used to create topology views.
//
//
//
//  * @param topologyViewDtoList view
// * @return *TaskIDResult
func (s *TopologyApplicationService) CreateTopologyViews3(applicationUUID string, pageUUID string, topologyViewDtoList *[]TopologyViewDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("POST", path, topologyViewDtoList)
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

// DeleteTopologyApplication is ...
// This method is used to delete a topology application.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyApplication(applicationUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
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

// DeleteTopologyApplication4 is ...
// This method is used to delete a topology application.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyApplication4(applicationUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
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

// DeleteTopologyPage is ...
// This method is used to delete a topology page.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyPage(applicationUUID string, pageUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
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

// DeleteTopologyPage5 is ...
// This method is used to delete a topology page.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyPage5(applicationUUID string, pageUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
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

// DeleteTopologyView is ...
// This method is used to delete a topology view.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyView(applicationUUID string, pageUUID string, viewUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}/view/{viewUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	path = strings.Replace(path, "{"+"viewUUID"+"}", fmt.Sprintf("%v", viewUUID), -1)
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

// DeleteTopologyView6 is ...
// This method is used to delete a topology view.
//
//
//
//
// * @return *TaskIDResult
func (s *TopologyApplicationService) DeleteTopologyView6(applicationUUID string, pageUUID string, viewUUID string) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}/view/{viewUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	path = strings.Replace(path, "{"+"viewUUID"+"}", fmt.Sprintf("%v", viewUUID), -1)
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

// GetTopologyApplication is ...
// This method is used to obtain a topology aplication
//
//
//
//
// * @return *TopologyApplicationResult
func (s *TopologyApplicationService) GetTopologyApplication(applicationUUID string) (*TopologyApplicationResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyApplicationResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPage is ...
// This method is used to obtain a topology aplication page
//
//
//
//
// * @return *TopologyPageResult
func (s *TopologyApplicationService) GetTopologyApplicationPage(applicationUUID string, pageUUID string) (*TopologyPageResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyPageResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPageView is ...
// This method is used to obtain a topology aplication page view
//
//
//
//
// * @return *TopologyViewResult
func (s *TopologyApplicationService) GetTopologyApplicationPageView(applicationUUID string, pageUUID string, viewUUID string) (*TopologyViewResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}/view/{viewUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	path = strings.Replace(path, "{"+"viewUUID"+"}", fmt.Sprintf("%v", viewUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyViewResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPageView7 is ...
// This method is used to obtain a topology aplication page view
//
//
//
//
// * @return *TopologyViewResult
func (s *TopologyApplicationService) GetTopologyApplicationPageView7(applicationUUID string, pageUUID string, viewUUID string) (*TopologyViewResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}/view/{viewUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	path = strings.Replace(path, "{"+"viewUUID"+"}", fmt.Sprintf("%v", viewUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyViewResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPageViews is ...
// This method is used to obtain the list of topology aplication page Views for an topology application Page
//
//
//
//
// * @return *TopologyViewListResult
func (s *TopologyApplicationService) GetTopologyApplicationPageViews(applicationUUID string, pageUUID string) (*TopologyViewListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyViewListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPageViews8 is ...
// This method is used to obtain the list of topology aplication page Views for an topology application Page
//
//
//
//
// * @return *TopologyViewListResult
func (s *TopologyApplicationService) GetTopologyApplicationPageViews8(applicationUUID string, pageUUID string) (*TopologyViewListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyViewListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPage9 is ...
// This method is used to obtain a topology aplication page
//
//
//
//
// * @return *TopologyPageResult
func (s *TopologyApplicationService) GetTopologyApplicationPage9(applicationUUID string, pageUUID string) (*TopologyPageResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyPageResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPages is ...
// This method is used to obtain the list of topology aplication pages for an topology application
//
//
//
//
// * @return *TopologyPageListResult
func (s *TopologyApplicationService) GetTopologyApplicationPages(applicationUUID string) (*TopologyPageListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyPageListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplicationPages10 is ...
// This method is used to obtain the list of topology aplication pages for an topology application
//
//
//
//
// * @return *TopologyPageListResult
func (s *TopologyApplicationService) GetTopologyApplicationPages10(applicationUUID string) (*TopologyPageListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyPageListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplication11 is ...
// This method is used to obtain a topology aplication
//
//
//
//
// * @return *TopologyApplicationResult
func (s *TopologyApplicationService) GetTopologyApplication11(applicationUUID string) (*TopologyApplicationResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyApplicationResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplications is ...
// This method is used to obtain the list of topology aplications
//
//
//
//
// * @return *TopologyApplicationListResult
func (s *TopologyApplicationService) GetTopologyApplications() (*TopologyApplicationListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyApplicationListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTopologyApplications12 is ...
// This method is used to obtain the list of topology aplications
//
//
//
//
// * @return *TopologyApplicationListResult
func (s *TopologyApplicationService) GetTopologyApplications12() (*TopologyApplicationListResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TopologyApplicationListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateTopologyApplication is ...
// This method is used to update topology applications.
//
//
//
//  * @param topologyApplicationDtoList application
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyApplication(topologyApplicationDtoList *[]TopologyApplicationDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application"
	req, err := s.client.NewRequest("PUT", path, topologyApplicationDtoList)
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

// UpdateTopologyApplication13 is ...
// This method is used to update topology applications.
//
//
//
//  * @param topologyApplicationDtoList application
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyApplication13(topologyApplicationDtoList *[]TopologyApplicationDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp"
	req, err := s.client.NewRequest("PUT", path, topologyApplicationDtoList)
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

// UpdateTopologyPages is ...
// This method is used to update topology pages.
//
//
//
//  * @param topologyPageDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyPages(applicationUUID string, topologyPageDtoList *[]TopologyPageDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("PUT", path, topologyPageDtoList)
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

// UpdateTopologyPages14 is ...
// This method is used to update topology pages.
//
//
//
//  * @param topologyPageDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyPages14(applicationUUID string, topologyPageDtoList *[]TopologyPageDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	req, err := s.client.NewRequest("PUT", path, topologyPageDtoList)
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

// UpdateTopologyViews is ...
// This method is used to update topology views.
//
//
//
//  * @param topologyViewDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyViews(applicationUUID string, pageUUID string, topologyViewDtoList *[]TopologyViewDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/application/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("PUT", path, topologyViewDtoList)
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

// UpdateTopologyViews15 is ...
// This method is used to update topology views.
//
//
//
//  * @param topologyViewDtoList page
// * @return *TaskIDResult
func (s *TopologyApplicationService) UpdateTopologyViews15(applicationUUID string, pageUUID string, topologyViewDtoList *[]TopologyViewDto) (*TaskIDResult, *Response, error) {

	path := topologyApplicationBasePath + "/topology/controllerapp/{applicationUUID}/page/{pageUUID}/view"
	path = strings.Replace(path, "{"+"applicationUUID"+"}", fmt.Sprintf("%v", applicationUUID), -1)
	path = strings.Replace(path, "{"+"pageUUID"+"}", fmt.Sprintf("%v", pageUUID), -1)
	req, err := s.client.NewRequest("PUT", path, topologyViewDtoList)
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
