package client

import "github.com/zenorachi/literature-list/pkg/client/models"

type IClient interface {
	SearchLiterature(literatureList []string) ([]models.LiteratureList, error)
}
