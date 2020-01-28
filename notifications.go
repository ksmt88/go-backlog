package backlog

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
	"unsafe"
)

type CountNotificationQuery struct {
	AlreadyRead         bool // default: false
	ResourceAlreadyRead bool // default: false
}

type Notification struct {
	ID                  int     `json:"id"`
	AlreadyRead         bool    `json:"alreadyRead"`
	Reason              int     `json:"reason"`
	ResourceAlreadyRead bool    `json:"resourceAlreadyRead"`
	Project             Project `json:"project"`
	Issue               struct {
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
		Assignee       User          `json:"assignee"`
		Category       []interface{} `json:"category"`
		Versions       []interface{} `json:"versions"`
		Milestone      []interface{} `json:"milestone"`
		StartDate      time.Time     `json:"startDate"`
		DueDate        time.Time     `json:"dueDate"`
		EstimatedHours interface{}   `json:"estimatedHours"`
		ActualHours    interface{}   `json:"actualHours"`
		ParentIssueID  interface{}   `json:"parentIssueId"`
		CreatedUser    User          `json:"createdUser"`
		Created        time.Time     `json:"created"`
		UpdatedUser    User          `json:"updatedUser"`
		Updated        time.Time     `json:"updated"`
		CustomFields   []interface{} `json:"customFields"`
		Attachments    []Attachment  `json:"attachments"`
		SharedFiles    []SharedFile  `json:"sharedFiles"`
		Stars          []interface{} `json:"stars"`
	} `json:"issue"`
	Comment struct {
		ID            int           `json:"id"`
		Content       string        `json:"content"`
		ChangeLog     interface{}   `json:"changeLog"`
		CreatedUser   User          `json:"createdUser"`
		Created       time.Time     `json:"created"`
		Updated       time.Time     `json:"updated"`
		Stars         []interface{} `json:"stars"`
		Notifications []interface{} `json:"notifications"`
	} `json:"comment"`
	PullRequest        interface{} `json:"pullRequest"`
	PullRequestComment interface{} `json:"pullRequestComment"`
	Sender             User        `json:"sender"`
	Created            time.Time   `json:"created"`
}

func (s *Service) GetNotification() ([]Notification, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var notifications []Notification
	err = json.Unmarshal(body, &notifications)
	if err != nil {
		return nil, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return notifications, nil
}

func (s *Service) CountNotification(query CountNotificationQuery) (int, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/count"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	if query.AlreadyRead {
		urlParams.Add("alreadyRead", "true")
	} else {
		urlParams.Add("alreadyRead", "false")
	}
	if query.ResourceAlreadyRead {
		urlParams.Add("resourceAlreadyRead", "true")
	} else {
		urlParams.Add("resourceAlreadyRead", "false")
	}

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var count struct {
		Count int
	}
	err = json.Unmarshal(body, &count)
	if err != nil {
		return 0, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return count.Count, nil
}

func (s *Service) ResetUnreadNotificationCount() (int, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/markAsRead"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Post(requestUrl+"?"+urlParams.Encode(), "application/json", nil)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var count struct {
		Count int
	}
	err = json.Unmarshal(body, &count)
	if err != nil {
		return 0, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return count.Count, nil
}

func (s *Service) ReadNotification(id int) (bool, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/" + strconv.Itoa(id) + "/markAsRead"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Post(requestUrl+"?"+urlParams.Encode(), "application/json", nil)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}
