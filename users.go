package backlog

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type User struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Password    string `json:"-"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"` // Administrator(1) Normal User(2) Reporter(3) Viewer(4) Guest Reporter(5) Guest Viewer(6)
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}

func (s *Service) GetUserList() ([]User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users"
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

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUser(userId int) (User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var user User

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return user, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) AddUser(user User) (User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	requestParams := url.Values{}
	requestParams.Add("userId", user.UserID)
	requestParams.Add("password", user.Password)
	requestParams.Add("name", user.Name)
	requestParams.Add("mailAddress", user.MailAddress)
	requestParams.Add("roleType", strconv.Itoa(user.RoleType))

	var addUser User

	res, err := http.Post(requestUrl+"?"+urlParams.Encode(), "application/x-www-form-urlencoded", strings.NewReader(requestParams.Encode()))
	if err != nil {
		return addUser, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return addUser, err
	}

	err = json.Unmarshal(body, &addUser)
	if err != nil {
		return addUser, err
	}

	return addUser, nil
}

func (s *Service) UpdateUser(userId int, user User) (User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	requestParams := url.Values{}
	requestParams.Add("password", user.Password)
	requestParams.Add("name", user.Name)
	requestParams.Add("mailAddress", user.MailAddress)
	requestParams.Add("roleType", strconv.Itoa(user.RoleType))

	var updateUser User

	req, err := http.NewRequest(http.MethodPatch, requestUrl+"?"+urlParams.Encode(), strings.NewReader(requestParams.Encode()))
	if err != nil {
		return updateUser, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := s.client.Do(req)
	if err != nil {
		return updateUser, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return updateUser, err
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *Service) DeleteUser(userId int) (User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var user User

	req, err := http.NewRequest(http.MethodDelete, requestUrl+"?"+urlParams.Encode(), nil)
	if err != nil {
		return user, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return user, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetOwnUser() (User, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/myself"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	var user User

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return user, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetUserIcon(userId int) (image.Image, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId) + "/icon"
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	res, err := s.client.Get(requestUrl + "?" + urlParams.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// func (s *Service) GetUserRecentUpdates(userId int, query QueryActivities) (string, error) {}
// func (s *Service) GetReceivedStarList(userId int, query QueryStar) (string, error) {}
// func (s *Service) CountUserReceivedStars(query QueryStarsCount) (string, error) {}
// func (s *Service) GetListOfRecentlyViewedIssues(query RecentlyViewed) (string, error) {}
// func (s *Service) GetListOfRecentlyViewedProjects(query RecentlyViewed) (string, error) {}
// func (s *Service) GetListOfRecentlyViewedWikis(query RecentlyViewed) (string, error) {}
