package cyberleninka

import (
	"github.com/zenorachi/literature-list/pkg/clients/models"
	"net/http"
)

type IClient interface {
	SearchLiterature(literatureList []string) ([]models.LiteratureList, error)
}

type client struct {
	client *http.Client
}

func NewClient() IClient {
	return &client{
		client: &http.Client{},
	}
}

func (c client) SearchLiterature(literatureList []string) ([]models.LiteratureList, error) {
	panic("implement me")
}
