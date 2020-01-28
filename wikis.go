package backlog

import (
	"encoding/json"
	"errors"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type GetWikiPageListQuery struct {
	ProjectIdOrKey int
	Keyword        string
}

type WikiPageQuery struct {
	ProjectIdOrKey int
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Attachment struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	CreatedUser User      `json:"createdUser"`
	Created     time.Time `json:"created"`
}

type SharedFile struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"`
	Dir         string    `json:"dir"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	CreatedUser User      `json:"createdUser"`
	Created     time.Time `json:"created"`
	UpdatedUser User      `json:"updatedUser"`
	Updated     time.Time `json:"updated"`
}

type Wiki struct {
	ProjectId  int
	Name       string
	Content    string
	MailNotify bool
}

type WikiListItem struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectId"`
	Name        string    `json:"name"`
	Tags        []Tag     `json:"tags"`
	CreatedUser User      `json:"createdUser"`
	Created     time.Time `json:"created"`
	UpdatedUser User      `json:"updatedUser"`
	Updated     time.Time `json:"updated"`
}

type DetailWiki struct {
	ID          int           `json:"id"`
	ProjectID   int           `json:"projectId"`
	Name        string        `json:"name"`
	Content     string        `json:"content"`
	Tags        []Tag         `json:"tags"`
	Attachments []Attachment  `json:"attachments"`
	SharedFiles []SharedFile  `json:"sharedFiles"`
	Stars       []interface{} `json:"stars"`
	CreatedUser User          `json:"createdUser"`
	Created     time.Time     `json:"created"`
	UpdatedUser User          `json:"updatedUser"`
	Updated     time.Time     `json:"updated"`
}

func (s *Service) GetWikiPageList(query GetWikiPageListQuery) ([]WikiListItem, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))
	urlParams.Add("keyword", query.Keyword)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var wikiListItems []WikiListItem
	err = json.Unmarshal(body, &wikiListItems)
	if err != nil {
		return nil, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return wikiListItems, nil
}

func (s *Service) CountWikiPage(query WikiPageQuery) (int, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/count"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))

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

func (s *Service) GetWikiPageTagList(query WikiPageQuery) ([]Tag, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/tags"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	urlParams.Add("projectIdOrKey", strconv.Itoa(query.ProjectIdOrKey))

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tags []Tag
	err = json.Unmarshal(body, &tags)
	if err != nil {
		return nil, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return tags, nil
}

func (s *Service) AddWikiPage(wiki Wiki) (DetailWiki, error) {
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

	var addWiki DetailWiki

	res, err := http.Post(requestUrl+"?"+urlParams.Encode(), "application/x-www-form-urlencoded", strings.NewReader(requestParams.Encode()))
	if err != nil {
		return addWiki, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return addWiki, err
	}

	err = json.Unmarshal(body, &addWiki)
	if err != nil {
		return addWiki, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return addWiki, nil
}

func (s *Service) GetWikiPage(wikiId int) (DetailWiki, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/" + strconv.Itoa(wikiId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var wiki DetailWiki

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return wiki, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return wiki, err
	}

	err = json.Unmarshal(body, &wiki)
	if err != nil {
		return wiki, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return wiki, nil
}

func (s *Service) UpdateWikiPage(wikiId int, wiki Wiki) (DetailWiki, error) {
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

	var detailWiki DetailWiki

	req, err := http.NewRequest(http.MethodPatch, requestUrl+"?"+urlParams.Encode(), strings.NewReader(requestParams.Encode()))
	if err != nil {
		return detailWiki, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := s.client.Do(req)
	if err != nil {
		return detailWiki, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return detailWiki, err
	}

	err = json.Unmarshal(body, &detailWiki)
	if err != nil {
		return detailWiki, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return detailWiki, nil
}

func (s *Service) DeleteWikiPage(wikiId int) (DetailWiki, error) {
	requestUrl := s.BaseUrl + "/api/v2/wikis/" + strconv.Itoa(wikiId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var wiki DetailWiki

	req, err := http.NewRequest(http.MethodDelete, requestUrl+"?"+urlParams.Encode(), nil)
	if err != nil {
		return wiki, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return wiki, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return wiki, err
	}

	err = json.Unmarshal(body, &wiki)
	if err != nil {
		return wiki, errors.New(*(*string)(unsafe.Pointer(&body)))
	}

	return wiki, nil
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
