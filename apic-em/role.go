package apicem

const roleBasePath = "v1"

// RoleService is an interface with the Userrole API
type RoleService service

// Role is ...
type Role struct {
	Role string `json:"role,omitempty"` // User Role
}

// RoleListResult is ...
type RoleListResult struct {
	Version  string `json:"version,omitempty"`
	Response []Role `json:"response,omitempty"`
}

// GetRoles is ...
// This method is used to get the list of roles
//
//
//
//
// * @return *RoleListResult
func (s *RoleService) GetRoles() (*RoleListResult, *Response, error) {

	path := roleBasePath + "/user/role"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RoleListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
