package client

import (
	"encoding/json"
	"fmt"
	"github.com/zenorachi/literature-list/pkg/client/models"
	"net/http"
	"strings"
	"time"
)

const DefaultTimeout = 10 * time.Second

type IClient interface {
	SearchLiterature(literatureList []string) ([]models.LiteratureList, error)
}

type BaseClient struct {
	baseURL string
	client  *http.Client
	timeout time.Duration
}

func New(baseURL string, timeout time.Duration) *BaseClient {
	return &BaseClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		client:  &http.Client{},
		timeout: timeout,
	}
}

func (c *BaseClient) Get(endpoint string) ([]byte, error) {
	response, err := c.client.Get(fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(endpoint, "/")))
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	return json.Marshal(response.Body)
}
