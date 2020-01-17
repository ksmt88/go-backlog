package backlog

import (
	"io/ioutil"
	"net/url"
)

func (s *Service) GetNotifications() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications"
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

func (s *Service) GetNotificationsCount(alreadyRead bool, resourceAlreadyRead bool) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/count"
	requestParams := url.Values{}
	requestParams.Add("apiKey", s.Config.ApiKey)

	// add params
	if alreadyRead {
		requestParams.Add("alreadyRead", "true")
	} else {
		requestParams.Add("alreadyRead", "false")
	}
	if resourceAlreadyRead {
		requestParams.Add("resourceAlreadyRead", "true")
	} else {
		requestParams.Add("resourceAlreadyRead", "false")
	}

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

func (s *Service) PostNotificationsMarkAsRead() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/markAsRead"
	requestParams := url.Values{}
	requestParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Post(requestUrl+"?"+requestParams.Encode(), "application/json", nil)
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
