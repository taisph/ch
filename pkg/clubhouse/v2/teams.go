package v2

import (
	"fmt"
	"strconv"
	"time"
)

type Team struct {
	CreatedAt   *time.Time `json:"created_at"`
	Description *string    `json:"description"`
	EntityType  string     `json:"entity_type"`
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Position    int64      `json:"position"`
	ProjectIds  []int64    `json:"project_ids"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Workflow    *Workflow  `json:"workflow"`
}

type Workflow struct {
	CreatedAt      *time.Time      `json:"created_at"`
	DefaultStateId int64           `json:"default_state_id"`
	Description    *string         `json:"description"`
	EntityType     string          `json:"entity_type"`
	Id             int64           `json:"id"`
	Name           string          `json:"name"`
	States         []WorkflowState `json:"states"`
	TeamId         int64           `json:"id"`
	UpdatedAt      *time.Time      `json:"updated_at"`
}

type WorkflowState struct {
	Color       string     `json:"color"`
	CreatedAt   *time.Time `json:"created_at"`
	Description *string    `json:"description"`
	EntityType  string     `json:"entity_type"`
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	NumStories  int64      `json:"num_stories"`
	Position    int64      `json:"position"`
	Type        string     `json:"type"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Verb        *string    `json:"string"`
}

type Teams struct {
	c *Client
}

func (c *Client) Teams() *Teams {
	return &Teams{c: c}
}

func (s *Teams) List() ([]*Team, error) {
	var res []*Team
	err := s.c.get("teams", nil, &res)
	if err != nil {
		return nil, fmt.Errorf("error listing teams: %s", err.Error())
	}
	return res, nil
}

func (s *Teams) Get(id int64) (*Team, error) {
	var res *Team
	err := s.c.get("teams/"+strconv.FormatInt(id, 10), nil, &res)
	if err != nil {
		return nil, fmt.Errorf("error getting team id: %d: %s", id, err.Error())
	}
	return res, nil
}
