package backlog

import (
	"io/ioutil"
	"net/url"
)

type CountNotificationQuery struct {
	AlreadyRead         bool // default: false
	ResourceAlreadyRead bool // default: false
}

func (s *Service) GetNotification() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications"
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

func (s *Service) CountNotification(query CountNotificationQuery) (string, error) {
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
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (s *Service) ResetUnreadNotificationCount() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/markAsRead"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Post(requestUrl+"?"+urlParams.Encode(), "application/json", nil)
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

/*
func (s *Service) ReadNotification(id uint) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/notifications/" + string(id) + "/markAsRead"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Post(requestUrl+"?"+urlParams.Encode(), "application/json", nil)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}*/
