package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/atlassian/go-artifactory/v2/artifactory/client"
	"net/http"
	"time"
)

type UsersService Service

type UserStatus string

const (
	UserStatusInvited  UserStatus = "invited"
	UserStatusEnabled  UserStatus = "enabled"
	UserStatusDisabled UserStatus = "disabled"
	UserStatusLocked   UserStatus = "locked"
)

type User struct {
	Username                 *string   	 `json:"username,omitempty"`
	Password                 *string   	 `json:"password,omitempty"`
	Email                    *string   	 `json:"email,omitempty"`
	Groups                   *[]string 	 `json:"groups,omitempty"`
	Admin                    *bool     	 `json:"admin,omitempty"`
	ProfileUpdatable         *bool     	 `json:"profile_updatable,omitempty"`
	InternalPasswordDisabled *bool     	 `json:"internal_password_disabled,omitempty"`
	DisableUiAccess          *bool     	 `json:"disable_ui_access,omitempty"`
	WatchManager             *bool     	 `json:"watch_manager,omitempty"`
	PolicyViewer             *bool     	 `json:"policy_viewer,omitempty"`
	PolicyManager            *bool     	 `json:"policy_manager,omitempty"`
	ReportsManager           *bool     	 `json:"reports_manager,omitempty"`
	ResourcesManager         *bool     	 `json:"resources_manager,omitempty"`
	ManageWebhook            *bool     	 `json:"manage_webhook,omitempty"`
	PlatformAuditor          *bool     	 `json:"platform_auditor,omitempty"`
	Realm                    *string   	 `json:"realm,omitempty"`
	EffectiveAdmin           *bool     	 `json:"effective_admin,omitempty"`
	LastLoggedIn             *time.Time  `json:"last_logged_in,omitempty"`
	Status                   *UserStatus `json:"status,omitempty"`
}

func (u User) String() string {
	res, _ := json.MarshalIndent(u, "", "    ")
	return string(res)
}

type UserUpdateParams struct {
	Password                 *string   `json:"password,omitempty"`
	Email                    *string   `json:"email,omitempty"`
	Groups                   *[]string `json:"groups,omitempty"`
	Admin                    *bool     `json:"admin,omitempty"`
	ProfileUpdatable         *bool     `json:"profile_updatable,omitempty"`
	InternalPasswordDisabled *bool     `json:"internal_password_disabled,omitempty"`
	DisableUiAccess          *bool     `json:"disable_ui_access,omitempty"`
	WatchManager             *bool     `json:"watch_manager,omitempty"`
	PolicyViewer             *bool     `json:"policy_viewer,omitempty"`
	PolicyManager            *bool     `json:"policy_manager,omitempty"`
	ReportsManager           *bool     `json:"reports_manager,omitempty"`
	ResourcesManager         *bool     `json:"resources_manager,omitempty"`
	ManageWebhook            *bool     `json:"manage_webhook,omitempty"`
	PlatformAuditor          *bool     `json:"platform_auditor,omitempty"`
}

func (u UserUpdateParams) String() string {
	res, _ := json.MarshalIndent(u, "", "    ")
	return string(res)
}

type UserGroupsUpdateParams struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

type ListUsersOptions struct {
	Status          UserStatus `url:"status,omitempty"`
	Limit           int        `url:"limit,omitempty"`
	Username        string     `url:"username,omitempty"`
	OnlyAdmins      bool       `url:"onlyAdmins,omitempty"`
	Cursor          string     `url:"cursor,omitempty"`
	Role            string     `url:"role,omitempty"`
	ResourceType    string     `url:"resourceType,omitempty"`
	ResourceName    string     `url:"resourceName,omitempty"`
	ProjectKey      string     `url:"projectKey,omitempty"`
	DescendingOrder bool       `url:"descendingOrder,omitempty"`
}

type UserListItem struct {
	Username *string 		 `json:"username,omitempty"`
	Uri      *string 		 `json:"uri,omitempty"`
	Realm    *string 		 `json:"realm,omitempty"`
	Status   *UserStatus `json:"status,omitempty"`
}

type UserList struct {
	Users  []UserListItem `json:"users"`
	Cursor *string        `json:"cursor,omitempty"`
}

func (s *UsersService) ListUsers(ctx context.Context, opts *ListUsersOptions) (*UserList, *http.Response, error) {
	path, err := client.AddOptions("/access/api/v2/users", opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	list := new(UserList)
	resp, err := s.client.Do(ctx, req, list)
	return list, resp, err
}

func (s *UsersService) ExpireUserPassword(ctx context.Context, username string) (*http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s/password/expire", username)
	req, err := s.client.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) UnlockUser(ctx context.Context, username string) (*http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s/unlock", username)
	req, err := s.client.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) ExpirePasswordForAllUsers(ctx context.Context) (*http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/access/api/v2/users/expire_password_for_all_users", nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) UnexpirePasswordForAllUsers(ctx context.Context) (*http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/access/api/v2/users/unexpire_password_for_all_users", nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) ChangeUserPassword(ctx context.Context, username string, password string) (*http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s/password", username)
	body := struct {
		Password string `json:"password"`
	}{Password: password}
	req, err := s.client.NewJSONEncodedRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) UpdateUserGroups(ctx context.Context, username string, params *UserGroupsUpdateParams) ([]string, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s/groups", username)
	req, err := s.client.NewJSONEncodedRequest(http.MethodPatch, path, params)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	var result struct {
		Groups []string `json:"groups"`
	}
	resp, err := s.client.Do(ctx, req, &result)
	return result.Groups, resp, err
}

func (s *UsersService) DeleteUser(ctx context.Context, username string) (*http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s", username)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *UsersService) GetUser(ctx context.Context, username string) (*User, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s", username)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	user := new(User)
	resp, err := s.client.Do(ctx, req, user)
	return user, resp, err
}

func (s *UsersService) UpdateUser(ctx context.Context, username string, params *UserUpdateParams) (*User, *http.Response, error) {
	path := fmt.Sprintf("/access/api/v2/users/%s", username)
	req, err := s.client.NewJSONEncodedRequest(http.MethodPatch, path, params)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	updated := new(User)
	resp, err := s.client.Do(ctx, req, updated)
	return updated, resp, err
}

func (s *UsersService) CreateUser(ctx context.Context, user *User) (*User, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest(http.MethodPost, "/access/api/v2/users", user)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", client.MediaTypeJson)

	created := new(User)
	resp, err := s.client.Do(ctx, req, created)
	return created, resp, err
}
