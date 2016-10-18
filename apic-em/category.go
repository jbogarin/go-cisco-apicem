package apicem

import (
	"fmt"
	"strings"
)

const categoryBasePath = "v1"

// CategoryService is an interface with the Category API
type CategoryService service

// CategoryListResult is ...
type CategoryListResult struct {
	Response []CategoryDto `json:"response,omitempty"`
	Version  string        `json:"version,omitempty"`
}

// CategoryResult is ...
type CategoryResult struct {
	Response CategoryDto `json:"response,omitempty"`
	Version  string      `json:"version,omitempty"`
}

// GetAllCategories is ...
// Get all categories
//
//
//
//
// * @return *CategoryListResult
func (s *CategoryService) GetAllCategories() (*CategoryListResult, *Response, error) {

	path := categoryBasePath + "/category"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CategoryListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCategoryByID is ...
// Get category by id
//
//
//
//
// * @return *CategoryResult
func (s *CategoryService) GetCategoryByID(id string) (*CategoryResult, *Response, error) {

	path := categoryBasePath + "/category/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CategoryResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetCategoryCount is ...
// Get total count for categories
//
//
//
//
// * @return *CountResult
func (s *CategoryService) GetCategoryCount() (*CountResult, *Response, error) {

	path := categoryBasePath + "/category/count"
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
