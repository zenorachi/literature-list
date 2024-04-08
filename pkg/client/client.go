package client

import (
	"bytes"
	"fmt"
	"github.com/zenorachi/literature-list/pkg/client/models"
	"io"
	"net/http"
	"strings"
	"time"
)

const DefaultTimeout = 10 * time.Second

type IClient interface {
	SearchLiterature(endpoint string, literatureList []string) ([]models.LiteratureList, error)
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

func (c *BaseClient) Post(endpoint string, body []byte) ([]byte, error) {
	response, err := c.client.Post(
		fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(endpoint, "/")),
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	return io.ReadAll(response.Body)
}
