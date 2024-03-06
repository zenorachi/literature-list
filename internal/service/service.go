package service

import (
	"github.com/zenorachi/literature-list/internal/models"
	"github.com/zenorachi/literature-list/pkg/client/cyberleninka"
	"github.com/zenorachi/literature-list/pkg/client/elibrary"
)

type IService interface {
	SearchLiterature(literatureList []string) ([]models.Literature, error)
}

type Service struct {
	cyberleninkaClient cyberleninka.IClient
	elibraryClient     elibrary.IClient
}

// SearchLiterature TODO: refactor to goroutines
func (s *Service) SearchLiterature(literatureList []string) ([]models.Literature, error) {
	var (
		result       = make([]models.Literature, len(literatureList))
		notFoundList = make([]string, 0)
	)

	cyberleninkaList, err := s.cyberleninkaClient.SearchLiterature(literatureList)
	if err != nil {
		return nil, err
	}

	for _, v := range cyberleninkaList {
		if !v.IsContained {
			notFoundList = append(notFoundList, v.Title)
		} else {
			result = append(result, models.Literature{
				Title:       v.Title,
				IsContained: v.IsContained,
			})
		}
	}

	elibraryList, err := s.elibraryClient.SearchLiterature(notFoundList)
	if err != nil {
		return nil, err
	}

	for _, v := range elibraryList {
		result = append(result, models.Literature{
			Title:       v.Title,
			IsContained: v.IsContained,
		})
	}

	return result, nil
}
