package backlog

import (
	"io/ioutil"
	"net/url"
)

func (s *Service) GetSpace() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/space"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// func (s *Service) GetSpaceActivities() (string, error) {}
// func (s *Service) GetSpaceImage() (*File, error) {}

func (s *Service) GetSpaceNotification() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/space/notification"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// func (s *Service) UpdateSpaceNotification() (string, error) {}
// func (s *Service) GutSpaceDiskUsage() (string, error) {}
// func (s *Service) PostAttachmentFile(file *File) (string, error) {}
