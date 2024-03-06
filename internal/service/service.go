package service

import (
	"fmt"
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
	var (
	//result       = make([]models.Literature, len(literatureList))
	//notFoundList = make([]string, 0)
	)

	cyberleninkaList, err := s.cyberleninkaClient.SearchLiterature(cmodels.CyberleninkaEndpoint, literatureList)
	if err != nil {
		return nil, err
	}

	for _, v := range cyberleninkaList {
		fmt.Printf("Title: %s\nContained: %v\n", v.Title, v.IsContained)
	}

	//for _, v := range cyberleninkaList {
	//	if !v.IsContained {
	//		notFoundList = append(notFoundList, v.Title)
	//	} else {
	//		result = append(result, models.Literature{
	//			Title:       v.Title,
	//			IsContained: v.IsContained,
	//		})
	//	}
	//}
	//
	//elibraryList, err := s.elibraryClient.SearchLiterature(notFoundList)
	//if err != nil {
	//	return nil, err
	//}
	//
	//for _, v := range elibraryList {
	//	result = append(result, models.Literature{
	//		Title:       v.Title,
	//		IsContained: v.IsContained,
	//	})
	//}

	return nil, nil
}
