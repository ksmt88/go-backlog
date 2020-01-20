package backlog

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WikiListQuery struct {
	ProjectIdOrKey int
	Keyword        string
}

type Wiki struct {
	ProjectId  int
	Name       string
	Content    string
	MailNotify bool
}

func (s *Service) GetWikiPageList(query WikiListQuery) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))
	urlParams.Add("keyword", query.Keyword)

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

func (s *Service) CountWikiPage(query WikiListQuery) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/count"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))

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

func (s *Service) GetWikiPageTagList(query WikiListQuery) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/tags"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))

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

func (s *Service) AddWikiPage(wiki Wiki) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	requestParams := url.Values{}
	requestParams.Add("projectId", strconv.Itoa(wiki.ProjectId))
	requestParams.Add("name", wiki.Name)
	requestParams.Add("content", wiki.Content)
	if wiki.MailNotify {
		requestParams.Add("mailNotify", "true")
	} else {
		requestParams.Add("mailNotify", "false")
	}

	res, err := http.Post(requestUrl+"?"+urlParams.Encode(), "application/x-www-form-urlencoded", strings.NewReader(requestParams.Encode()))
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

func (s *Service) GetWikiPage(wikiId int) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/" + strconv.Itoa(wikiId)
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

func (s *Service) UpdateWikiPage(wikiId int, wiki Wiki) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/" + strconv.Itoa(wikiId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	requestParams := url.Values{}
	requestParams.Add("name", wiki.Name)
	requestParams.Add("content", wiki.Content)
	if wiki.MailNotify {
		requestParams.Add("mailNotify", "true")
	} else {
		requestParams.Add("mailNotify", "false")
	}

	req, err := http.NewRequest(http.MethodPatch, requestUrl+"?"+urlParams.Encode(), strings.NewReader(requestParams.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := s.client.Do(req)
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

func (s *Service) DeleteWikiPage(wikiId int) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/" + strconv.Itoa(wikiId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	req, err := http.NewRequest(http.MethodDelete, requestUrl+"?"+urlParams.Encode(), nil)
	if err != nil {
		return "", err
	}

	res, err := s.client.Do(req)
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

// func (s *Service) GetListOfWikiAttachments() (string, error) {}
// func (s *Service) AttachFileToWiki() (string, error) {}
// func (s *Service) GetWikiPageAttachment() (string, error) {}
// func (s *Service) RemoveWikiAttachment() (string, error) {}
// func (s *Service) GetListOfSharedFilesOnWiki() (string, error) {}
// func (s *Service) LinkSharedFilesToWiki() (string, error) {}
// func (s *Service) RemoveLinkToSharedFileFromWiki() (string, error) {}
// func (s *Service) GetWikiPageHistory() (string, error) {}
// func (s *Service) GetWikiPageStar() (string, error) {}
