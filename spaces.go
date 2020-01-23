package backlog

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"
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
		return space, err
	}

	return space, nil
}

// func (s *Service) GetSpaceActivities() (string, error) {}
// func (s *Service) GetSpaceImage() (*File, error) {}

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
		return spaceNotification, err
	}

	return spaceNotification, nil
}

// func (s *Service) UpdateSpaceNotification() (string, error) {}
// func (s *Service) GutSpaceDiskUsage() (string, error) {}
// func (s *Service) PostAttachmentFile(file *File) (string, error) {}
