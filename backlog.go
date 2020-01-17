package backlog

import (
	"errors"
	"net/http"
)

type Configure struct {
	SpaceId string
	ApiKey  string
	Domain  string
}

type Service struct {
	client  *http.Client
	Config  Configure
	BaseUrl string
}

const (
	DomainJp   = ".backlog.jp"
	DomainCom  = ".backlog.com"
	DomainTool = ".backlogtool.com"
)

func NewClient(config Configure, client *http.Client) (*Service, error) {
	if config.SpaceId == "" {
		return nil, errors.New("SpaceId not found")
	}
	if config.ApiKey == "" {
		return nil, errors.New("ApiKey not found")
	}
	config.Domain = selectDomain(config.Domain)

	s := &Service{
		client:  client,
		Config:  config,
		BaseUrl: "https://" + config.SpaceId + config.Domain,
	}
	return s, nil
}

func selectDomain(domain string) string {
	switch domain {
	case DomainJp:
		return DomainJp
	case DomainCom:
		return DomainCom
	case DomainTool:
		return DomainTool
	default:
		return DomainJp
	}
}
