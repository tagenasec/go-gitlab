//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"net/http"
	"time"
)

// PersonalAccessTokensService handles communication with personal access tokens
// of the Gitlab API.
//
// Gitlab API docs : https://docs.gitlab.com/ee/api/personal_access_tokens.html
type PersonalAccessTokensService struct {
	client *Client
}

// Application represents a GitLab application
type PersonalAccessTokenMember struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Revoked   bool       `json:"revoked"`
	CreatedAt *time.Time `json:"created_at"`
	Scopes    []string   `json:"scopes"`
	UserID    int        `json:"user_id"`
	LastUsed  *time.Time `json:"createlast_used_at"`
	Active    bool       `json:"active"`
	ExpiresAt *ISOTime   `json:"expires_at"`
}

// ListPersonalAccessTokensOptions represents the available
// ListPersonalAccessTokens() options.
type ListPersonalAccessTokensOptions ListOptions

// ListPersonalAccessTokens get a list of personal access tokens by the authenticated user
//
// Gitlab API docs : https://docs.gitlab.com/ce/api/applications.html#list-all-applications
func (s *PersonalAccessTokensService) ListPersonalAccessTokens(opt *ListPersonalAccessTokensOptions, options ...RequestOptionFunc) ([]*PersonalAccessTokenMember, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "personal_access_tokens", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pat []*PersonalAccessTokenMember
	resp, err := s.client.Do(req, &pat)
	if err != nil {
		return nil, resp, err
	}

	return pat, resp, err
}
