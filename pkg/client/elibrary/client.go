package elibrary

import (
	c "github.com/zenorachi/literature-list/pkg/client"
	"github.com/zenorachi/literature-list/pkg/client/models"
	"time"
)

type Client struct {
	client *c.BaseClient
}

func NewClient(baseUrl string, timeout time.Duration) c.IClient {
	return &Client{
		client: c.New(baseUrl, timeout),
	}
}

func (c *Client) SearchLiterature(endpoint string, literatureList []string) ([]models.LiteratureList, error) {
	panic("implement me")
}
