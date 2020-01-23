package backlog

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type Project struct {
	ID                                int    `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}

func (s *Service) GetProjectList() ([]Project, error) {
	requestUrl := s.BaseUrl + "/api/v2/projects"
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

	var projects []Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
