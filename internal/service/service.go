package service

import (
	"github.com/zenorachi/literature-list/internal/config"
	"github.com/zenorachi/literature-list/internal/models"
	"github.com/zenorachi/literature-list/pkg/client"
	"github.com/zenorachi/literature-list/pkg/client/cyberleninka"
	"github.com/zenorachi/literature-list/pkg/client/elibrary"
	cmodels "github.com/zenorachi/literature-list/pkg/client/models"
)

type IService interface {
	SearchLiterature(literatureList []string) ([]models.Literature, error)
}

type Service struct {
	cyberleninkaClient client.IClient
	elibraryClient     client.IClient
}

func New(config *config.Config) *Service {
	return &Service{
		cyberleninkaClient: cyberleninka.NewClient(config.CyberleninkaClient.BaseURL, config.CyberleninkaClient.Timeout),
		elibraryClient:     elibrary.NewClient(config.ELibraryClient.BaseURL, config.ELibraryClient.Timeout),
	}
}

func (s *Service) SearchLiterature(literatureList []string) ([]models.Literature, error) {
	cyberleninkaList, err := s.cyberleninkaClient.SearchLiterature(cmodels.CyberleninkaEndpoint, literatureList)
	if err != nil {
		return nil, err
	}

	result := make([]models.Literature, 0, len(cyberleninkaList))
	for _, v := range cyberleninkaList {
		result = append(result, models.Literature{
			Title:       v.Title,
			IsContained: v.IsContained,
		})
	}

	return result, nil
}
