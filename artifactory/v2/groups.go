package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/artifactory/client"
	"net/http"
)

type GroupsService Service

type Group struct {
	Name 						 *string   `json:"name,omitempty"`
	Description      *string   `json:"description,omitempty"`
	AutoJoin         *bool     `json:"auto_join,omitempty"`
	AdminPrivileges  *bool     `json:"admin_privileges,omitempty"`
	Realm            *string   `json:"realm,omitempty"`
	RealmAttributes  *string   `json:"realm_attributes,omitempty"`
	ExternalId       *string   `json:"external_id,omitempty"`
	Members          *[]string `json:"members,omitempty"`
	ReportsManager   *bool     `json:"reports_manager,omitempty"`
	WatchManager     *bool     `json:"watch_manager,omitempty"`
	PolicyManager    *bool     `json:"policy_manager,omitempty"`
	PolicyViewer     *bool     `json:"policy_viewer,omitempty"`
	ManageResources  *bool     `json:"manage_resources,omitempty"`
	ManageWebhook    *bool     `json:"manage_webhook,omitempty"`
}

func (g Group) String() string {
	res, _ := json.MarshalIndent(g, "", "    ")
	return string(res)
}

type GroupUpdateParams struct {
	Description     *string `json:"description,omitempty"`
	AutoJoin        *bool   `json:"auto_join,omitempty"`
	AdminPrivileges *bool   `json:"admin_privileges,omitempty"`
	ExternalId      *string `json:"external_id,omitempty"`
	ReportsManager  *bool   `json:"reports_manager,omitempty"`
	WatchManager    *bool   `json:"watch_manager,omitempty"`
	PolicyManager   *bool   `json:"policy_manager,omitempty"`
	PolicyViewer    *bool   `json:"policy_viewer,omitempty"`
	ManageResources *bool   `json:"manage_resources,omitempty"`
	ManageWebhook   *bool   `json:"manage_webhook,omitempty"`
}

func (g GroupUpdateParams) String() string {
	res, _ := json.MarshalIndent(g, "", "    ")
	return string(res)
}

type GroupMembersUpdateParams struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

type ListGroupsOptions struct {
	Limit           int    `url:"limit,omitempty"`
	DescendingOrder bool   `url:"descendingOrder,omitempty"`
	GroupName       string `url:"groupName,omitempty"`
	Cursor          string `url:"cursor,omitempty"`
	Role            string `url:"role,omitempty"`
	ResourceType    string `url:"resourceType,omitempty"`
	ResourceName    string `url:"resourceName,omitempty"`
	ProjectKey      string `url:"projectKey,omitempty"`
}

type GroupListItem struct {
	GroupName *string `json:"group_name,omitempty"`
	Uri       *string `json:"uri,omitempty"`
}

type GroupList struct {
	Groups []GroupListItem `json:"groups"`
	Cursor *string         `json:"cursor,omitempty"`
}


func (s *GroupsService) CreateGroup(ctx context.Context, group *Group) (*Group, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest(http.MethodPost, "/access/api/v2/groups", group)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	created := new(Group)
	resp, err := s.client.Do(ctx, req, created)
	return created, resp, err
}

func (s *GroupsService) UpdateGroup(ctx context.Context, groupName string, params *GroupUpdateParams) (*Group, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/groups/%s", groupName)
	req, err := s.client.NewJSONEncodedRequest(http.MethodPatch, path, params)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	updated := new(Group)
	resp, err := s.client.Do(ctx, req, updated)
	return updated, resp, err
}

func (s *GroupsService) GetGroup(ctx context.Context, groupName string) (*Group, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/groups/%s", groupName)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	group := new(Group)
	resp, err := s.client.Do(ctx, req, group)
	return group, resp, err
}

func (s *GroupsService) ListGroups(ctx context.Context, opts *ListGroupsOptions) (*GroupList, *http.Response, error) {
	path, err := client.AddOptions("/access/api/v2/groups", opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	list := new(GroupList)
	resp, err := s.client.Do(ctx, req, list)
	return list, resp, err
}

func (s *GroupsService) UpdateGroupMembers(ctx context.Context, groupName string, params *GroupMembersUpdateParams) ([]string, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/groups/%s/members", groupName)
	req, err := s.client.NewJSONEncodedRequest(http.MethodPatch, path, params)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	var result struct {
		Members []string `json:"members"`
	}
	resp, err := s.client.Do(ctx, req, &result)
	return result.Members, resp, err
}

func (s *GroupsService) DeleteGroup(ctx context.Context, groupName string) (*http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/groups/%s", groupName)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
