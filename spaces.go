package backlog

import (
	"io/ioutil"
	"net/url"
)

func (s *Service) GetSpace() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/space"
	requestParams := url.Values{}
	requestParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + requestParams.Encode())
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
	requestParams := url.Values{}
	requestParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + requestParams.Encode())
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

// func (s *Service) PutSpaceNotification() (string, error) {}
// func (s *Service) GutSpaceDiskUsage() (string, error) {}
// func (s *Service) PostSpaceAttachment(file *File) (string, error) {}
