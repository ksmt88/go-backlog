package backlog

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/url"
	"strconv"
)

type User struct {
	UserId      string
	Password    string
	Name        string
	MailAddress string
	RoleType    int // Administrator(1) Normal User(2) Reporter(3) Viewer(4) Guest Reporter(5) Guest Viewer(6)
}

func (s *Service) GetUserList() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/users"
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

func (s *Service) GetUser(userId int) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId)
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

// func (s *Service) AddUser(userId string, password string, name string, mailAddress string, roleType int) (string, error) {}
/*
func (s *Service) UpdateUser(userId int, user User) (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/" + strconv.Itoa(userId)
	urlParams := url.Values{}
	urlParams.Add("apiKey", s.Config.ApiKey)

	// add params
	requestParams := url.Values{}
	if user.Password != "" {
		requestParams.Add("password", user.Password)
	}
	if user.Name != "" {
		requestParams.Add("name", user.Name)
	}
	if user.MailAddress != "" {
		requestParams.Add("mailAddress", user.MailAddress)
	}
	if user.RoleType != 0 {
		requestParams.Add("roleType", strconv.Itoa(user.RoleType))
	}

	res, err := http.NewRequest(http.MethodPatch, requestUrl+"?"+urlParams.Encode(), strings.NewReader(requestParams.Encode()))
	if err != nil {
		return "", err
	}

	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = s.client.Do(res)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("test: %v", string(body))
	time.Sleep(5 * time.Second)

	return string(body), nil
}*/

// func (s *Service) DeleteUser(userId int) (string, error) {}

func (s *Service) GetOwnUser() (string, error) {
	requestUrl := s.BaseUrl + "/api/v2/users/myself"
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
