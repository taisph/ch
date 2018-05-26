package v2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CreateStory struct {
	Name                string                  `json:"name"`
	ProjectId           int64                   `json:"project_id"`
	Comments            []CreateCommentParams   `json:"comments,omitempty"`
	CompletedAtOverride *time.Time              `json:"completed_at_override,omitempty"`
	Deadline            *time.Time              `json:"deadline,omitempty"`
	Description         string                  `json:"description,omitempty"`
	EpicId              int64                   `json:"epic_id,omitempty"`
	Estimate            int64                   `json:"estimate,omitempty"`
	ExternalId          string                  `json:"external_id,omitempty"`
	FileIds             []int64                 `json:"file_ids,omitempty"`
	FollowerIds         []uuid.UUID             `json:"follower_ids,omitempty"`
	Labels              []CreateLabelParams     `json:"labels,omitempty"`
	LinkedFileIds       []int64                 `json:"linked_file_ids,omitempty"`
	OwnerIds            []uuid.UUID             `json:"owner_ids,omitempty"`
	RequestedById       *uuid.UUID              `json:"requested_by_id,omitempty"`
	StartedAtOverride   *time.Time              `json:"started_at_override,omitempty"`
	StoryLinks          []CreateStoryLinkParams `json:"story_links,omitempty"`
	StoryType           StoryType               `json:"story_type,omitempty"`
	Tasks               []CreateTaskParams      `json:"tasks,omitempty"`
	UpdatedAt           *time.Time              `json:"updated_at,omitempty"`
	WorkflowStateId     int64                   `json:"workflow_state_id,omitempty"`
}

type CreateCommentParams struct {
	AuthorId   string    `json:"author_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ExternalId string    `json:"external_id,omitempty"`
	Text       string    `json:"text,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type CreateLabelParams struct {
	Color      string `json:"color,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Name       string `json:"name,omitempty"`
}

type CreateStoryLinkParams struct {
	ObjectId  int64 `json:"object_id,omitempty"`
	SubjectId int64 `json:"subject_id,omitempty"`
	Verb      Verb  `json:"verb,omitempty"`
}

type CreateTaskParams struct {
	Complete    bool        `json:"complete,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	Description string      `json:"description,omitempty"`
	ExternalId  string      `json:"external_id,omitempty"`
	OwnerIds    []uuid.UUID `json:"owner_ids,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
}

type StoryType string

const (
	StoryTypeBug     = StoryType("bug")
	StoryTypeFeature = StoryType("feature")
	StoryTypeChore   = StoryType("chore")
)

type Verb string

const (
	VerbBlocks     = Verb("blocks")
	VerbDuplicates = Verb("duplicates")
	VerbRelatesTo  = Verb("relates to")
)

// https://clubhouse.io/api/rest/v2/#Story
type Story struct {
	AppUrl   string   `json:"app_url"`
	Archived bool     `json:"archived"`
	Blocked  bool     `json:"blocked"`
	Blocker  bool     `json:"blocker"`
	Branches []Branch `json:"branches"`
	// TODO: Add remaining fields
}

// https://clubhouse.io/api/rest/v2/#Branch
type Branch struct {
	CreatedAt       time.Time     `json:"created_at"`
	Deleted         bool          `json:"deleted"`
	EntityType      string        `json:"entity_type"`
	Id              *int64        `json:"id"`
	MergedBranchIds []int64       `json:"merged_branch_ids"`
	Name            string        `json:"name"`
	Persistent      bool          `json:"persistent"`
	PullRequests    []PullRequest `json:"pull_requests"`
	RepositoryId    *int64        `json:"repository_id"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Url             string        `json:"url"`
}

// https://clubhouse.io/api/rest/v2/#PullRequest
type PullRequest struct {
	BranchId       *int64    `json:"branch_id"`
	Closed         bool      `json:"closed"`
	CreatedAt      time.Time `json:"created_at"`
	EntityType     string    `json:"entity_type"`
	Id             *int64    `json:"id"`
	NumAdded       int64     `json:"num_added"`
	NumCommits     int64     `json:"num_commits"`
	NumModified    int64     `json:"num_modified"`
	NumRemoved     int64     `json:"num_removed"`
	Number         int64     `json:"number"`
	TargetBranchId int64     `json:"target_branch_id"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
}

// https://clubhouse.io/api/rest/v2/#StorySlim
type StorySlim struct {
	AppUrl   string `json:"app_url"`
	Archived bool   `json:"archived"`
	Blocked  bool   `json:"blocked"`
	Blocker  bool   `json:"blocker"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	// TODO: Add remaining fields
}

func (c *Client) Stories() *Stories {
	return &Stories{c: c}
}

type Stories struct {
	c *Client
}

func (s *Stories) Create(story *CreateStory) (*Story, error) {
	res := &Story{}
	err := s.c.post("stories", story, res)
	if err != nil {
		return nil, fmt.Errorf("error creating story: %s", err.Error())
	}
	return res, nil
}

func (s *Stories) List(projectId int64) ([]*StorySlim, error) {
	var res []*StorySlim
	err := s.c.get("projects/"+strconv.FormatInt(projectId, 10)+"/stories", nil, &res)
	if err != nil {
		return nil, fmt.Errorf("error listing stories: %s", err.Error())
	}
	return res, nil
}
