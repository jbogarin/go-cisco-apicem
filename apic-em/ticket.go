package apicem

import (
	"fmt"
	"strings"
)

const ticketBasePath = "v1"

// TicketService is an interface with the Ticket API
type TicketService service

// TicketRbacResult is ...
type TicketRbacResult struct {
	Version string `json:"version,omitempty"`

	Response ServiceTicketRbac `json:"response,omitempty"`
}

// ServiceTicketRbac is used to retrieve the service ticket
type ServiceTicketRbac struct {
	IdleTimeout int32 `json:"idleTimeout,omitempty"`

	// Service Ticket to be used as authentication Ticket
	ServiceTicket string `json:"serviceTicket,omitempty"`

	SessionTimeout int32 `json:"sessionTimeout,omitempty"`
}

// SuccessResult is ...
type SuccessResult struct {
	Version string `json:"version,omitempty"`

	Response string `json:"response,omitempty"`
}

// TicketAttributeResult is ...
type TicketAttributeResult struct {
	Version string `json:"version,omitempty"`

	Response TicketAttribute `json:"response,omitempty"`
}

// TicketAttribute is used to retrieve the ticket attributes
type TicketAttribute struct {

	// Service Ticket Attribute Name
	Attribute string `json:"attribute,omitempty"`

	// Service Ticket Attribute Value
	Value int64 `json:"value,omitempty"`
}

// User is
type User struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// AddTicket is ...
// This method is used to create a new user ticket
//
//  * @param user user
//  * @return *TicketRbacResult
func (s *TicketService) AddTicket(user *User) (*TicketRbacResult, *Response, error) {

	path := ticketBasePath + "/ticket"

	req, err := s.client.NewRequest("Post", path, user)
	if err != nil {
		return nil, nil, err
	}

	root := new(TicketRbacResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// CreateTicketAttribute is ...
// Create and update specific ticket attribute, Idle Timeout Minimum 60 seconds, Max 3600 seconds, default 900 seconds, Session Timeout Minimum 1800 seconds, Max 86400 seconds, default 21600 seconds
//
//  * @param ticketAttribute ticketAttribute
//  * @return *SuccessResult
func (s *TicketService) CreateTicketAttribute(ticketAttribute *TicketAttribute) (*SuccessResult, *Response, error) {

	path := ticketBasePath + "/ticket/attribute"

	req, err := s.client.NewRequest("Post", path, ticketAttribute)
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

// DeleteTicket is ...
// Revoke ticket, effectively sign out
//
//  * @param ticket ticket
//  * @return *SuccessResult
func (s *TicketService) DeleteTicket(ticket string) (*SuccessResult, *Response, error) {

	path := ticketBasePath + "/ticket/{ticket}"
	path = strings.Replace(path, "{"+"ticket"+"}", fmt.Sprintf("%v", ticket), -1)

	req, err := s.client.NewRequest("Delete", path, nil)
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

// DeleteTicketAttribute is ...
// Delete Ticket Attribute
//
//  * @param attribute attribute
//  * @return *SuccessResult
func (s *TicketService) DeleteTicketAttribute(attribute string) (*SuccessResult, *Response, error) {

	path := ticketBasePath + "/ticket/attribute/{attribute}"
	path = strings.Replace(path, "{"+"attribute"+"}", fmt.Sprintf("%v", attribute), -1)

	req, err := s.client.NewRequest("Delete", path, nil)
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

// GetIdleTimeout is ...
// Get Idle timeout
//
//  * @return *TicketAttributeResult
func (s *TicketService) GetIdleTimeout() (*TicketAttributeResult, *Response, error) {

	path := ticketBasePath + "/ticket/attribute/idletimeout"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TicketAttributeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetSessionTimeout is ...
// Get Session timeout
//
//  * @return *TicketAttributeResult
func (s *TicketService) GetSessionTimeout() (*TicketAttributeResult, *Response, error) {

	path := ticketBasePath + "/ticket/attribute/sessiontimeout"

	req, err := s.client.NewRequest("Get", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TicketAttributeResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
