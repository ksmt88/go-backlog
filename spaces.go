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

type Space struct {
	SpaceKey           string    `json:"spaceKey"`
	Name               string    `json:"name"`
	OwnerID            int       `json:"ownerId"`
	Lang               string    `json:"lang"`
	Timezone           string    `json:"timezone"`
	ReportSendTime     string    `json:"reportSendTime"`
	TextFormattingRule string    `json:"textFormattingRule"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}

type GetRecentUpdatesQuery struct {
	ActivityTypeId []int
	minId          int
	maxId          int
	count          int
	order          string
}

type Changes struct {
	Field    string `json:"field"`
	NewValue string `json:"new_value"`
	OldValue string `json:"old_value"`
	Type     string `json:"type"`
}

type Content struct {
	ID          int    `json:"id"`
	KeyID       int    `json:"key_id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Comment     struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	} `json:"comment"`
	Changes []Changes `json:"changes"`
}

type RecentUpdate struct {
	ID            int     `json:"id"`
	Project       Project `json:"project"`
	Type          int     `json:"type"`
	Content       Content `json:"content"`
	Notifications []struct {
		ID                  int  `json:"id"`
		AlreadyRead         bool `json:"alreadyRead"`
		Reason              int  `json:"reason"`
		User                User `json:"user"`
		ResourceAlreadyRead bool `json:"resourceAlreadyRead"`
	} `json:"notifications"`
	CreatedUser User      `json:"createdUser"`
	Created     time.Time `json:"created"`
}

type SpaceNotification struct {
	Content string    `json:"content"`
	Updated time.Time `json:"updated"`
}

func (s *Service) GetSpace() (Space, error) {
	requestUrl := s.BaseUrl + "/api/v2/space"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var space Space

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return space, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return space, err
	}

	err = json.Unmarshal(body, &space)
	if err != nil {
		return space, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return space, nil
}

func (s *Service) GetRecentUpdates(query GetRecentUpdatesQuery) ([]RecentUpdate, error) {
	requestUrl := s.BaseUrl + "/api/v2/space/activities"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)
	for _, typeId := range query.ActivityTypeId {
		urlParams.Add("activityTypeId[]", strconv.Itoa(typeId))
	}
	urlParams.Add("minId", strconv.Itoa(query.minId))
	urlParams.Add("maxId", strconv.Itoa(query.maxId))
	urlParams.Add("count", strconv.Itoa(query.count))
	urlParams.Add("order", query.order)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var recentUpdates []RecentUpdate
	err = json.Unmarshal(body, &recentUpdates)
	if err != nil {
		return nil, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return recentUpdates, nil
}

// func (s *Service) GetSpaceLogo() (image, error) {}

func (s *Service) GetSpaceNotification() (SpaceNotification, error) {
	requestUrl := s.BaseUrl + "/api/v2/space/notification"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var spaceNotification SpaceNotification

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return spaceNotification, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return spaceNotification, err
	}

	err = json.Unmarshal(body, &spaceNotification)
	if err != nil {
		return spaceNotification, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return spaceNotification, nil
}

// func (s *Service) UpdateSpaceNotification() (string, error) {}
// func (s *Service) GutSpaceDiskUsage() (string, error) {}
// func (s *Service) PostAttachmentFile(file *File) (string, error) {}
