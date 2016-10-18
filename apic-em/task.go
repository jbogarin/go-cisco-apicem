package apicem

import (
	"fmt"
	"strings"
	"time"
)

const taskBasePath = "v1"

// TaskService is an interface with the Task API
type TaskService service

// AugmentedTaskDto is ... Represents a task
type AugmentedTaskDto struct {
	ID              string           `json:"id,omitempty"`          // id
	StartTime       time.Time        `json:"startTime,omitempty"`   // startTime
	EndTime         time.Time        `json:"endTime,omitempty"`     // endTime
	Progress        string           `json:"progress,omitempty"`    // progress
	Data            string           `json:"data,omitempty"`        // data
	ErrorCode       string           `json:"errorCode,omitempty"`   // errorCode
	Version         int64            `json:"version,omitempty"`     // version
	ServiceType     string           `json:"serviceType,omitempty"` // serviceType
	OperationIDList Collectionstring `json:"operationIdList,omitempty"`
	ParentID        string           `json:"parentId,omitempty"`      // parentID
	RootID          string           `json:"rootId,omitempty"`        // rootId
	LastUpdate      time.Time        `json:"lastUpdate,omitempty"`    // lastUpdate
	Username        string           `json:"username,omitempty"`      // username
	IsError         bool             `json:"isError,omitempty"`       // isError
	FailureReason   string           `json:"failureReason,omitempty"` // failureReason
}

// Collectionstring is ...
type Collectionstring struct {
}

// TaskDtoListResponse is ...
type TaskDtoListResponse struct {
	Version  string             `json:"version,omitempty"`
	Response []AugmentedTaskDto `json:"response,omitempty"`
}

// TaskDtoResponse is ...
type TaskDtoResponse struct {
	Version  string           `json:"version,omitempty"`
	Response AugmentedTaskDto `json:"response,omitempty"`
}

// GetTask is ...
// This method is used to retrieve a task based on their id
//
//
//
//
// * @return *TaskDtoResponse
func (s *TaskService) GetTask(taskID string) (*TaskDtoResponse, *Response, error) {

	path := taskBasePath + "/task/{taskID}"
	path = strings.Replace(path, "{"+"taskID"+"}", fmt.Sprintf("%v", taskID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskDtoResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTaskByRange is ...
// This method is used to list number of tasks (Pagination)
//
//
//
//
// * @return *TaskDtoListResponse
func (s *TaskService) GetTaskByRange(offset int32, limit int32) (*TaskDtoListResponse, *Response, error) {

	path := taskBasePath + "/task/{offset}/{limit}"
	path = strings.Replace(path, "{"+"offset"+"}", fmt.Sprintf("%v", offset), -1)
	path = strings.Replace(path, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskDtoListResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetTaskTree is ...
// This method is used to retrieve a task with its children tasks based on their id
//
//
//
//
// * @return *TaskDtoListResponse
func (s *TaskService) GetTaskTree(taskID string) (*TaskDtoListResponse, *Response, error) {

	path := taskBasePath + "/task/{taskID}/tree"
	path = strings.Replace(path, "{"+"taskID"+"}", fmt.Sprintf("%v", taskID), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TaskDtoListResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
