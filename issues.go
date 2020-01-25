package backlog

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

type GetIssueListQuery struct {
	ProjectId      []int
	IssueTypeId    []int
	CategoryId     []int
	VersionId      []int
	MilestoneId    []int
	StatusId       []int
	PriorityId     []int
	AssigneeId     []int
	CreatedUserId  []int
	ResolutionId   []int
	ParentChild    int // 0: All, 1: Exclude Child Issue, 2: Child Issue, 3: Neither Parent Issue nor Child Issue, 4: Parent Issue
	Attachment     bool
	SharedFile     bool
	Sort           string // issueType, category, version, milestone, summary, status, priority, attachment, sharedFile, created, createdUser, updated, updatedUser, assignee, startDate, dueDate, estimatedHours, actualHours, childIssue, customField_${id}
	Order          string
	Offset         int
	Count          int
	CreatedSince   string
	CreatedUntil   string
	UpdatedSince   string
	UpdatedUntil   string
	StartDateSince string
	StartDateUntil string
	DueDateSince   string
	DueDateUntil   string
	Id             []int
	ParentIssueId  []int
	Keyword        string
}

type Issue struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"projectId"`
	IssueKey  string `json:"issueKey"`
	KeyID     int    `json:"keyId"`
	IssueType struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		DisplayOrder int    `json:"displayOrder"`
	} `json:"issueType"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Resolutions interface{} `json:"resolutions"`
	Priority    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Status struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		DisplayOrder int    `json:"displayOrder"`
	} `json:"status"`
	Assignee  User          `json:"assignee"`
	Category  []interface{} `json:"category"`
	Versions  []interface{} `json:"versions"`
	Milestone []struct {
		ID             int         `json:"id"`
		ProjectID      int         `json:"projectId"`
		Name           string      `json:"name"`
		Description    string      `json:"description"`
		StartDate      interface{} `json:"startDate"`
		ReleaseDueDate interface{} `json:"releaseDueDate"`
		Archived       bool        `json:"archived"`
		DisplayOrder   int         `json:"displayOrder"`
	} `json:"milestone"`
	StartDate      interface{}   `json:"startDate"`
	DueDate        interface{}   `json:"dueDate"`
	EstimatedHours interface{}   `json:"estimatedHours"`
	ActualHours    interface{}   `json:"actualHours"`
	ParentIssueID  interface{}   `json:"parentIssueId"`
	CreatedUser    User          `json:"createdUser"`
	Created        time.Time     `json:"created"`
	UpdatedUser    User          `json:"updatedUser"`
	Updated        time.Time     `json:"updated"`
	CustomFields   []interface{} `json:"customFields"`
	Attachments    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Size int    `json:"size"`
	} `json:"attachments"`
	SharedFiles []SharedFile `json:"sharedFiles"`
	Stars       []struct {
		ID        int         `json:"id"`
		Comment   interface{} `json:"comment"`
		URL       string      `json:"url"`
		Title     string      `json:"title"`
		Presenter struct {
			ID          int    `json:"id"`
			UserID      string `json:"userId"`
			Name        string `json:"name"`
			RoleType    int    `json:"roleType"`
			Lang        string `json:"lang"`
			MailAddress string `json:"mailAddress"`
		} `json:"presenter"`
		Created time.Time `json:"created"`
	} `json:"stars"`
}

func (s *Service) GetIssueList(query GetIssueListQuery) ([]Issue, error) {
	requestUrl := s.BaseUrl + "/api/v2/issues"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)
	for _, projectId := range query.ProjectId {
		urlParams.Add("projectId[]", strconv.Itoa(projectId))
	}
	for _, issueTypeId := range query.IssueTypeId {
		urlParams.Add("issueTypeId[]", strconv.Itoa(issueTypeId))
	}
	for _, categoryId := range query.CategoryId {
		urlParams.Add("categoryId[]", strconv.Itoa(categoryId))
	}
	for _, versionId := range query.VersionId {
		urlParams.Add("versionId[]", strconv.Itoa(versionId))
	}
	for _, milestoneId := range query.MilestoneId {
		urlParams.Add("milestoneId[]", strconv.Itoa(milestoneId))
	}
	for _, statusId := range query.StatusId {
		urlParams.Add("statusId[]", strconv.Itoa(statusId))
	}
	for _, priorityId := range query.PriorityId {
		urlParams.Add("priorityId[]", strconv.Itoa(priorityId))
	}
	for _, assigneeId := range query.AssigneeId {
		urlParams.Add("assigneeId[]", strconv.Itoa(assigneeId))
	}
	for _, createdUserId := range query.CreatedUserId {
		urlParams.Add("createdUserId[]", strconv.Itoa(createdUserId))
	}
	for _, resolutionId := range query.ResolutionId {
		urlParams.Add("resolutionId[]", strconv.Itoa(resolutionId))
	}
	urlParams.Add("parentChild", strconv.Itoa(query.ParentChild))
	if query.Attachment {
		urlParams.Add("attachment", "true")
	} else {
		urlParams.Add("attachment", "false")
	}
	if query.SharedFile {
		urlParams.Add("sharedFile", "true")
	} else {
		urlParams.Add("sharedFile", "false")
	}
	urlParams.Add("sort", query.Sort)
	urlParams.Add("order", query.Order)
	urlParams.Add("offset", strconv.Itoa(query.Offset))
	urlParams.Add("count", strconv.Itoa(query.Count))
	urlParams.Add("createdSince", query.CreatedSince)
	urlParams.Add("createdUntil", query.CreatedUntil)
	urlParams.Add("updatedSince", query.UpdatedSince)
	urlParams.Add("updatedUntil", query.UpdatedUntil)
	urlParams.Add("startDateSince", query.StartDateSince)
	urlParams.Add("startDateUntil", query.StartDateUntil)
	urlParams.Add("dueDateSince", query.DueDateSince)
	urlParams.Add("dueDateUntil", query.DueDateUntil)
	for _, id := range query.Id {
		urlParams.Add("id[]", strconv.Itoa(id))
	}
	for _, parentIssueId := range query.ParentIssueId {
		urlParams.Add("parentIssueId[]", strconv.Itoa(parentIssueId))
	}
	urlParams.Add("keyword", query.Keyword)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
