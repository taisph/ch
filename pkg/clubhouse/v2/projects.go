package v2

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Abbreviation      *string      `json:"abbreviation"`
	Archived          bool         `json:"archived"`
	Color             *string      `json:"color"`
	CreatedAt         *time.Time   `json:"created_at"`
	DaysToThermometer int64        `json:"days_to_thermometer"`
	Description       *string      `json:"description"`
	EntityType        string       `json:"entity_type"`
	ExternalId        *string      `json:"external_id"`
	FollowerIds       []uuid.UUID  `json:"follower_ids"`
	Id                int64        `json:"id"`
	IterationLength   int64        `json:"iteration_length"`
	Name              string       `json:"name"`
	ShowThermometer   bool         `json:"show_thermometer"`
	StartTime         *time.Time   `json:"start_time"`
	Stats             ProjectStats `json:"stats"`
	TeamId            int64        `json:"team_id"`
	UpdatedAt         *time.Time   `json:"updated_at"`
}

type ProjectStats struct {
	NumPoints  int64 `json:"num_points"`
	NumStories int64 `json:"num_stories"`
}

type Projects struct {
	c *Client
}

func (c *Client) Projects() *Projects {
	return &Projects{c: c}
}

func (s *Projects) List() ([]*Project, error) {
	var res []*Project
	err := s.c.get("projects", nil, &res)
	if err != nil {
		return nil, fmt.Errorf("error listing projects: %s", err.Error())
	}
	return res, nil
}
