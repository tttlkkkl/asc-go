package asc

import (
	"context"
	"fmt"
)

// AppInfo defines model for AppInfo.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo
type AppInfo struct {
	Attributes    *AppInfoAttributes    `json:"attributes,omitempty"`
	ID            string                `json:"id"`
	Links         ResourceLinks         `json:"links"`
	Relationships *AppInfoRelationships `json:"relationships,omitempty"`
	Type          string                `json:"type"`
}

// AppInfoAttributes defines model for AppInfo.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo/attributes
type AppInfoAttributes struct {
	AppStoreAgeRating *AppStoreAgeRating    `json:"appStoreAgeRating,omitempty"`
	AppStoreState     *AppStoreVersionState `json:"appStoreState,omitempty"`
	BrazilAgeRating   *BrazilAgeRating      `json:"brazilAgeRating,omitempty"`
	KidsAgeBand       *KidsAgeBand          `json:"kidsAgeBand,omitempty"`
}

// AppInfoRelationships defines model for AppInfo.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfo/relationships
type AppInfoRelationships struct {
	App                     *Relationship      `json:"app,omitempty"`
	AppInfoLocalizations    *PagedRelationship `json:"appInfoLocalizations,omitempty"`
	PrimaryCategory         *Relationship      `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   *Relationship      `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   *Relationship      `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       *Relationship      `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *Relationship      `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *Relationship      `json:"secondarySubcategoryTwo,omitempty"`
}

// AppInfoResponse defines model for AppInfoResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinforesponse
type AppInfoResponse struct {
	Data     AppInfo       `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfosresponse
type AppInfosResponse struct {
	Data     []AppInfo          `json:"data"`
	Included []interface{}      `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// appInfoUpdateRequest defines model for AppInfoUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest/data
type appInfoUpdateRequest struct {
	ID            string                             `json:"id"`
	Relationships *appInfoUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppInfoUpdateRequestRelationships are relationships for AppInfoUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appinfoupdaterequest/data/relationships
type appInfoUpdateRequestRelationships struct {
	PrimaryCategory         *relationshipDeclaration `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   *relationshipDeclaration `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   *relationshipDeclaration `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       *relationshipDeclaration `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *relationshipDeclaration `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *relationshipDeclaration `json:"secondarySubcategoryTwo,omitempty"`
}

// AppInfoUpdateRequestRelationships is a public-facing options object for AppInfoUpdateRequest relationships
type AppInfoUpdateRequestRelationships struct {
	PrimaryCategoryID         *string
	PrimarySubcategoryOneID   *string
	PrimarySubcategoryTwoID   *string
	SecondaryCategoryID       *string
	SecondarySubcategoryOneID *string
	SecondarySubcategoryTwoID *string
}

// GetAppInfoQuery are query options for GetAppInfo
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
type GetAppInfoQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        []string `url:"fields[appCategories],omitempty"`
	Include                    []string `url:"include,omitempty"`
	LimitAppInfoLocalizations  int      `url:"limit[appInfoLocalizations],omitempty"`
}

// ListAppInfosForAppQuery are query options for ListAppInfosForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
type ListAppInfosForAppQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsApps                 []string `url:"fields[apps],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        []string `url:"fields[appCategories],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Include                    []string `url:"include,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetAppInfo reads App Store information including your App Store state, age ratings, Brazil age rating, and kids' age band.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_information
func (s *AppsService) GetAppInfo(ctx context.Context, id string, params *GetAppInfoQuery) (*AppInfoResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListAppInfosForApp gets information about an app that is currently live on App Store, or that goes live with the next version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_infos_for_an_app
func (s *AppsService) ListAppInfosForApp(ctx context.Context, id string, params *ListAppInfosForAppQuery) (*AppInfosResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/appInfos", id)
	res := new(AppInfosResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateAppInfo updates the App Store categories and sub-categories for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_info
func (s *AppsService) UpdateAppInfo(ctx context.Context, id string, relationships *AppInfoUpdateRequestRelationships) (*AppInfoResponse, *Response, error) {
	req := appInfoUpdateRequest{
		ID:   id,
		Type: "appInfos",
	}
	if relationships != nil {
		req.Relationships = &appInfoUpdateRequestRelationships{
			PrimaryCategory:         newRelationship(relationships.PrimaryCategoryID, "appCategories"),
			PrimarySubcategoryOne:   newRelationship(relationships.PrimarySubcategoryOneID, "appCategories"),
			PrimarySubcategoryTwo:   newRelationship(relationships.PrimarySubcategoryTwoID, "appCategories"),
			SecondaryCategory:       newRelationship(relationships.SecondaryCategoryID, "appCategories"),
			SecondarySubcategoryOne: newRelationship(relationships.SecondarySubcategoryOneID, "appCategories"),
			SecondarySubcategoryTwo: newRelationship(relationships.SecondarySubcategoryTwoID, "appCategories"),
		}
	}
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.client.patch(ctx, url, req, res)
	return res, resp, err
}
