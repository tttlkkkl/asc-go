package users

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal"
)

// UserInvitation defines model for UserInvitation.
type UserInvitation struct {
	Attributes *struct {
		AllAppsVisible      *bool           `json:"allAppsVisible,omitempty"`
		Email               *internal.Email `json:"email,omitempty"`
		ExpirationDate      *time.Time      `json:"expirationDate,omitempty"`
		FirstName           *string         `json:"firstName,omitempty"`
		LastName            *string         `json:"lastName,omitempty"`
		ProvisioningAllowed *bool           `json:"provisioningAllowed,omitempty"`
		Roles               *[]UserRole     `json:"roles,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		VisibleApps *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"visibleApps,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// UserInvitationCreateRequest defines model for UserInvitationCreateRequest.
type UserInvitationCreateRequest struct {
	Data struct {
		Attributes struct {
			AllAppsVisible      *bool          `json:"allAppsVisible,omitempty"`
			Email               internal.Email `json:"email"`
			FirstName           string         `json:"firstName"`
			LastName            string         `json:"lastName"`
			ProvisioningAllowed *bool          `json:"provisioningAllowed,omitempty"`
			Roles               []UserRole     `json:"roles"`
		} `json:"attributes"`
		Relationships *struct {
			VisibleApps *struct {
				Data *[]internal.RelationshipsData `json:"data,omitempty"`
			} `json:"visibleApps,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

// UserInvitationResponse defines model for UserInvitationResponse.
type UserInvitationResponse struct {
	Data     UserInvitation         `json:"data"`
	Included *[]apps.App            `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
}

// UserInvitationsResponse defines model for UserInvitationsResponse.
type UserInvitationsResponse struct {
	Data     []UserInvitation            `json:"data"`
	Included *[]apps.App                 `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// ListInvitationsQuery is the query params structure for ListInvitations
type ListInvitationsQuery struct {
	FieldsApps            *[]string `url:"fields[apps],omitempty"`
	FieldsUserInvitations *[]string `url:"fields[userInvitations],omitempty"`
	FilterRoles           *[]string `url:"filter[roles],omitempty"`
	FilterEmail           *[]string `url:"filter[email],omitempty"`
	FilterVisibleApps     *[]string `url:"filter[visibleApps],omitempty"`
	Include               *[]string `url:"include,omitempty"`
	Limit                 *int      `url:"limit,omitempty"`
	LimitVisibleApps      *int      `url:"limit[visibleApps],omitempty"`
	Sort                  *[]string `url:"sort,omitempty"`
	Cursor                *string   `url:"cursor,omitempty"`
}

// GetInvitationQuery is the query params structure for GetInvitation
type GetInvitationQuery struct {
	FieldsApps            *[]string `url:"fields[apps],omitempty"`
	FieldsUserInvitations *[]string `url:"fields[userInvitations],omitempty"`
	Include               *[]string `url:"include,omitempty"`
	LimitVisibleApps      *int      `url:"limit[visibleApps],omitempty"`
}

// ListInvitations gets a list of pending invitations to join your team.
func (s *Service) ListInvitations(params *ListInvitationsQuery) (*UserInvitationsResponse, *internal.Response, error) {
	res := new(UserInvitationsResponse)
	resp, err := s.GetWithQuery("userInvitations", params, res)
	return res, resp, err
}

// GetInvitation gets information about a pending invitation to join your team.
func (s *Service) GetInvitation(id string, params *GetInvitationQuery) (*UserInvitationResponse, *internal.Response, error) {
	url := fmt.Sprintf("userInvitations/%s", id)
	res := new(UserInvitationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateInvitation invites a user with assigned user roles to join your team.
func (s *Service) CreateInvitation(body *UserInvitationCreateRequest) (*UserInvitationResponse, *internal.Response, error) {
	res := new(UserInvitationResponse)
	resp, err := s.Post("userInvitations", body, res)
	return res, resp, err
}

// CancelInvitation cancels a pending invitation for a user to join your team.
func (s *Service) CancelInvitation(id string) (*internal.Response, error) {
	url := fmt.Sprintf("userInvitations/%s", id)
	return s.Delete(url, nil)
}

// ListVisibleAppsForInvitation gets a list of apps that will be visible to a user with a pending invitation.
func (s *Service) ListVisibleAppsForInvitation(id string, params ListVisibleAppsQuery) (*apps.AppsResponse, *internal.Response, error) {
	url := fmt.Sprintf("userInvitations/%s/visibleApps", id)
	res := new(apps.AppsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
