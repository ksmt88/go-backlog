package backlog

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type GetWikiPageListQuery struct {
	ProjectIdOrKey int
	Keyword        string
}

// func (s *Service) GetWikiPageList(params GetWikiPageListQuery) (string, error) {}
// func (s *Service) CountWikiPage() (string, error) {}
// func (s *Service) GetWikiPageTagList() (string, error) {}
// func (s *Service) AddWikiPage() (string, error) {}
// func (s *Service) GetWikiPage() (string, error) {}
// func (s *Service) UpdateWikiPage() (string, error) {}
// func (s *Service) DeleteWikiPage() (string, error) {}
// func (s *Service) GetListOfWikiAttachments() (string, error) {}
// func (s *Service) AttachFileToWiki() (string, error) {}
// func (s *Service) GetWikiPageAttachment() (string, error) {}
// func (s *Service) RemoveWikiAttachment() (string, error) {}
// func (s *Service) GetListOfSharedFilesOnWiki() (string, error) {}
// func (s *Service) LinkSharedFilesToWiki() (string, error) {}
// func (s *Service) RemoveLinkToSharedFileFromWiki() (string, error) {}
// func (s *Service) GetWikiPageHistory() (string, error) {}
// func (s *Service) GetWikiPageStar() (string, error) {}
