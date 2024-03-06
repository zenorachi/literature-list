package elibrary

import (
	c "github.com/zenorachi/literature-list/pkg/client"
	"github.com/zenorachi/literature-list/pkg/client/models"
	"net/http"
)

type client struct {
	client *http.Client
}

func NewClient() c.IClient {
	return &client{
		client: &http.Client{},
	}
}

func (c client) SearchLiterature(literatureList []string) ([]models.LiteratureList, error) {
	panic("implement me")
}
