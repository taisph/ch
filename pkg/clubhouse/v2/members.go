/*
 * Copyright 2018 Tais P. Hansen
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v2

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// https://clubhouse.io/api/rest/v2/#Member
type Member struct {
	CreatedAt  *time.Time `json:"created_at"`
	Disabled   bool       `json:"disabled"`
	EntityType string     `json:"entity_type"`
	Id         uuid.UUID  `json:"id"`
	Profile    Profile    `json:"profile"`
	Role       string     `json:"role"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// https://clubhouse.io/api/rest/v2/#Profile
type Profile struct {
	Deactivated            bool      `json:"deactivated"`
	DisplayIcon            *Icon     `json:"display_icon,omitempty"`
	EmailAddress           *string   `json:"email_address,omitempty"`
	EntityType             string    `json:"entity_type"`
	GravatarHash           *string   `json:"gravatar_hash,omitempty"`
	Id                     uuid.UUID `json:"id"`
	MentionName            string    `json:"mention_name"`
	Name                   string    `json:"name"`
	TwoFactorAuthActivated bool      `json:"two_factor_auth_activated"`
}

// https://clubhouse.io/api/rest/v2/#Icon
type Icon struct {
	CreatedAt  *time.Time `json:"created_at"`
	EntityType string     `json:"entity_type"`
	Id         uuid.UUID  `json:"id"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Url        string     `json:"string"`
}

type Members struct {
	c *Client
}

func (c *Client) Members() *Members {
	return &Members{c: c}
}

func (s *Members) List() ([]*Member, error) {
	var res []*Member
	err := s.c.get("members", nil, &res)
	if err != nil {
		return nil, fmt.Errorf("error listing members: %s", err.Error())
	}
	return res, nil
}

func (s *Members) GetByName(name string) (*Member, error) {
	members, err := s.List()
	if err != nil {
		return nil, fmt.Errorf("error getting member: %s: %s", name, err.Error())
	}

	var res *Member
	for _, member := range members {
		if strings.EqualFold(member.Profile.MentionName, name) {
			res = member
			break
		}
	}
	return res, nil
}
