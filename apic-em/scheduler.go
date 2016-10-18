package apicem

import (
	"fmt"
	"strings"
	"time"
)

const schedulerBasePath = "v1"

// SchedulerService is an interface with the Scheduledjob API
type SchedulerService service

// ScheduleInfoOutput is ...
type ScheduleInfoOutput struct {
	TaskID              string    `json:"taskId,omitempty"`              // UUID of the Task
	Description         string    `json:"description,omitempty"`         // Simple description to be shown to end-users
	Origin              string    `json:"origin,omitempty"`              // Contextual field used to identify work spcifications originating from the same source
	StartTime           time.Time `json:"startTime,omitempty"`           // The time at which the trigger should first fire. If the schedule has fired and will not fire again, this value will be null
	EndTime             time.Time `json:"endTime,omitempty"`             // The time at which the trigger should quit repeating
	Operation           string    `json:"operation,omitempty"`           // Contextual field used by the service to identify an operation
	GroupName           string    `json:"groupName,omitempty"`           // A grouping name that can be specified by the service to allow for filtered work spec retrieval
	ScheduledWorkSpecID string    `json:"scheduledWorkSpecId,omitempty"` // UUID of the ScheduledWorkSpec associated with the scheduled task
	NextTime            time.Time `json:"nextTime,omitempty"`            // The next time at which the trigger should fire
	PrevTime            time.Time `json:"prevTime,omitempty"`            // The previous time at which the trigger fired. If the trigger has not yet fired, null will be returned
}

// ScheduleInfoOutputListResult is ...
type ScheduleInfoOutputListResult struct {
	Version  string               `json:"version,omitempty"`
	Response []ScheduleInfoOutput `json:"response,omitempty"`
}

// ScheduleInfoOutputResult is ...
type ScheduleInfoOutputResult struct {
	Version  string             `json:"version,omitempty"`
	Response ScheduleInfoOutput `json:"response,omitempty"`
}

// DeleteScheduledTask is ...
// deleteScheduledTask
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *SuccessResult
func (s *SchedulerService) DeleteScheduledTask(scheduledWorkSpecID string, scope string) (*SuccessResult, *Response, error) {

	path := schedulerBasePath + "/scheduled-job/{scheduledWorkSpecID}"
	path = strings.Replace(path, "{"+"scheduledWorkSpecID"+"}", fmt.Sprintf("%v", scheduledWorkSpecID), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAllScheduledTasksQueryParams is ...
type GetAllScheduledTasksQueryParams struct {
	SortBy    string `url:"sortBy,omitempty"`    // sortBy
	Order     string `url:"order,omitempty"`     // orderBy
	GroupName string `url:"groupName,omitempty"` // filterByGroup
	Origin    string `url:"origin,omitempty"`    // filterByOrigin
	Operation string `url:"operation,omitempty"` // filterByOperation
	TaskID    string `url:"taskId,omitempty"`    // filterByTaskID
	Limit     string `url:"limit,omitempty"`     // limit
	Offset    string `url:"offset,omitempty"`    // offset
}

// GetAllScheduledTasks is ...
// getAllScheduledTasks
//
//  * @param queryParams
//  * @param scope Authorization Scope for RBAC
//
// * @return *ScheduleInfoOutputListResult
func (s *SchedulerService) GetAllScheduledTasks(scope string, queryParams *GetAllScheduledTasksQueryParams) (*ScheduleInfoOutputListResult, *Response, error) {

	path := schedulerBasePath + "/scheduled-job"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(ScheduleInfoOutputListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetScheduledTask is ...
// getScheduledTask
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *ScheduleInfoOutputResult
func (s *SchedulerService) GetScheduledTask(scheduledWorkSpecID string, scope string) (*ScheduleInfoOutputResult, *Response, error) {

	path := schedulerBasePath + "/scheduled-job/{scheduledWorkSpecID}"
	path = strings.Replace(path, "{"+"scheduledWorkSpecID"+"}", fmt.Sprintf("%v", scheduledWorkSpecID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(ScheduleInfoOutputResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
