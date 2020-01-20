package backlog

import (
	"io/ioutil"
	"net/url"
)

func (s *Service) GetProjectList() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/projects"
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
