package apicem

import (
	"fmt"
	"strings"
)

const userBasePath = "v1"

// UserService is an interface with the User API
type UserService service

// AutoPassphrase is ... Object used to retrieve the optional auto generated passphrase.
type AutoPassphrase struct {
	AutoPassphrase string `json:"autoPassphrase,omitempty"` // Auto generated passphrase
}

// AutoPassphraseResult is ...
type AutoPassphraseResult struct {
	Version  string         `json:"version,omitempty"`
	Response AutoPassphrase `json:"response,omitempty"`
}

// ScopeRole is ...
type ScopeRole struct {
	Scope string `json:"scope,omitempty"` // Scope of Authorization. Added to support future implementations, supported value is only ALL
	Role  string `json:"role,omitempty"`  // Authorization Role. Supported values are ROLE_ADMIN, ROLE_OBSERVER and ROLE_INSTALLER
}

// UserListResult is ...
type UserListResult struct {
	Version  string       `json:"version,omitempty"`
	Response []UserReqRes `json:"response,omitempty"`
}

// UserLockExpiryTime is ... Object used to retrieve the global user lock expiry timer
type UserLockExpiryTime struct {
	LockExpiryTime int32 `json:"lockExpiryTime,omitempty"` // User Lock Expiry Timer
}

// UserLockExpiryTimeResult is ...
type UserLockExpiryTimeResult struct {
	Version  string             `json:"version,omitempty"`
	Response UserLockExpiryTime `json:"response,omitempty"`
}

// UserLoginInvalidAttemptCount is ... Object used to retrieve the maximum user invalid attempts permitted before user account is locked.
type UserLoginInvalidAttemptCount struct {
	LoginInvalidAttemptCount int32 `json:"loginInvalidAttemptCount,omitempty"` // User Invalid Attempt Count
}

// UserLoginInvalidAttemptCountResult is ...
type UserLoginInvalidAttemptCountResult struct {
	Version  string                       `json:"version,omitempty"`
	Response UserLoginInvalidAttemptCount `json:"response,omitempty"`
}

// UserReqRes is ...
type UserReqRes struct {
	Username      string      `json:"username,omitempty"`      // Username
	Authorization []ScopeRole `json:"authorization,omitempty"` // User Authorization Scope
	AuthSource    string      `json:"authSource,omitempty"`    // User Authentication Source
}

// UserResult is ...
type UserResult struct {
	Version  string     `json:"version,omitempty"`
	Response UserReqRes `json:"response,omitempty"`
}

// UserStatus is ...
type UserStatus struct {
	Username       string `json:"username,omitempty"`
	AccountLocked  bool   `json:"accountLocked,omitempty"`
	LockedAt       string `json:"lockedAt,omitempty"`
	LockExpiration int64  `json:"lockExpiration,omitempty"`
}

// UserStatusResult is ...
type UserStatusResult struct {
	Version  string     `json:"version,omitempty"`
	Response UserStatus `json:"response,omitempty"`
}

// AddUser is ...
// This method is used to add a new user. &lt;b&gt;The password is excluded from the json schema below, but will still need to be included.&lt;/b&gt;
//
//
//
//  * @param user user
// * @return *SuccessResult
func (s *UserService) AddUser(user UserReqRes) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user"
	req, err := s.client.NewRequest("POST", path, user)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// DeleteUser is ...
// This method is used to delete a user.&lt;br/&gt; Admin permission is required.&lt;br/&gt; It is possible to delete your own user.
//
//
//
//
// * @return *SuccessResult
func (s *UserService) DeleteUser(username string) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user/{username}"
	path = strings.Replace(path, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAttemptCount is ...
// This method is used to get the max user invalid attempts before the user account is locked.
//
//
//
//
// * @return *UserLoginInvalidAttemptCountResult
func (s *UserService) GetAttemptCount() (*UserLoginInvalidAttemptCountResult, *Response, error) {

	path := userBasePath + "/user/password-policy/invalid-attempt-count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(UserLoginInvalidAttemptCountResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAutoPassphrase is ...
// This method is used to get an auto generated password.
//
//
//
//
// * @return *AutoPassphraseResult
func (s *UserService) GetAutoPassphrase() (*AutoPassphraseResult, *Response, error) {

	path := userBasePath + "/user/passphrase/auto"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AutoPassphraseResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAutoPassphraseWithSeedPhrase is ...
// This method is used to get an auto generated password with a seed phrase.
//
//
//
//
// * @return *AutoPassphraseResult
func (s *UserService) GetAutoPassphraseWithSeedPhrase(seedPhrase string) (*AutoPassphraseResult, *Response, error) {

	path := userBasePath + "/user/passphrase/auto/{seedPhrase}"
	path = strings.Replace(path, "{"+"seedPhrase"+"}", fmt.Sprintf("%v", seedPhrase), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AutoPassphraseResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetLockExpiry is ...
// This method is used to get the lock expiration timer value which is used to unlock a user account.
//
//
//
//
// * @return *UserLockExpiryTimeResult
func (s *UserService) GetLockExpiry() (*UserLockExpiryTimeResult, *Response, error) {

	path := userBasePath + "/user/password-policy/lock-expiry-time"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(UserLockExpiryTimeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetUser is ...
// This method is used to get user data by username.
//
//
//
//
// * @return *UserResult
func (s *UserService) GetUser(username string) (*UserResult, *Response, error) {

	path := userBasePath + "/user/{username}"
	path = strings.Replace(path, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(UserResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetUserStatus is ...
// This method is used to retrieve a user&#39;s lock status.&lt;br/&gt; Admin permission is required.&lt;br/&gt;
//
//
//
//
// * @return *UserStatusResult
func (s *UserService) GetUserStatus(username string) (*UserStatusResult, *Response, error) {

	path := userBasePath + "/user/status/{username}"
	path = strings.Replace(path, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(UserStatusResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetUsersQueryParams is ...
type GetUsersQueryParams struct {
	AuthSource string `url:"authSource,omitempty"` // authSource
}

// GetUsers is ...
// This method is used to get the list of Users. If you are an admin user, this will return a list of all the users, if you have an observer role, it will only show your own user information.
//
//  * @param queryParams
//
//
// * @return *UserListResult
func (s *UserService) GetUsers(queryParams *GetUsersQueryParams) (*UserListResult, *Response, error) {

	path := userBasePath + "/user"
	path, err := addOptions(path, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(UserListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateAttemptCount is ...
// This method is used to update the max user invalid attempts before the user account is locked.
//
//
//
//  * @param attemptCount attemptCount
// * @return *SuccessResult
func (s *UserService) UpdateAttemptCount(attemptCount UserLoginInvalidAttemptCount) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user/password-policy/invalid-attempt-count"
	req, err := s.client.NewRequest("PUT", path, attemptCount)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateLockExpiryTimeout is ...
// This method is used to set the lock expiration timer value which is used to unlock a user account.
//
//
//
//  * @param expiryTime expiryTime
// * @return *SuccessResult
func (s *UserService) UpdateLockExpiryTimeout(expiryTime UserLockExpiryTime) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user/password-policy/lock-expiry-time"
	req, err := s.client.NewRequest("PUT", path, expiryTime)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateUser is ...
// This method is used to update user data.&lt;br /&gt; To update the password, \&quot;oldPassword\&quot; will need to be provided and the new password can be represented as \&quot;password\&quot;.&lt;br /&gt; &lt;b&gt;A user can only update their own password&lt;/b&gt;.&lt;br /&gt; For an admin to update another admin or user&#39;s password, then the process is to delete and add a new user by the same name with the same permissions.
//
//
//
//  * @param user user
// * @return *SuccessResult
func (s *UserService) UpdateUser(user UserReqRes) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user"
	req, err := s.client.NewRequest("PUT", path, user)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateUserStatus is ...
// This method is used to update a user&#39;s lock status from locked to un-locked.&lt;br/&gt; Admin permission is required.&lt;br/&gt;
//
//
//
//  * @param userStatus userStatus
// * @return *SuccessResult
func (s *UserService) UpdateUserStatus(userStatus UserStatus) (*SuccessResult, *Response, error) {

	path := userBasePath + "/user/status"
	req, err := s.client.NewRequest("PUT", path, userStatus)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
