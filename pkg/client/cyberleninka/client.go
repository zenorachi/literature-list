package cyberleninka

import (
	"encoding/json"
	c "github.com/zenorachi/literature-list/pkg/client"
	"github.com/zenorachi/literature-list/pkg/client/models"
	"golang.org/x/sync/errgroup"
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
	var (
		eg      = &errgroup.Group{}
		resChan = make(chan *models.LiteratureList, len(literatureList))
		result  = make([]models.LiteratureList, 0, len(literatureList))
	)

	for _, title := range literatureList {
		title := title
		eg.Go(func() error {
			res, err := c.getOneArticle(endpoint, title)
			resChan <- res

			return err
		})
	}

	go func() {
		if err := eg.Wait(); err != nil {
			//logger.Error(err)
		}
		close(resChan)
	}()

	for v := range resChan {
		result = append(result, *v)
	}

	return result, nil
}

func (c *Client) getOneArticle(endpoint string, title string) (*models.LiteratureList, error) {
	req := models.CyberleninkaRequest{
		Mode: models.Mode,
		Size: models.Size,
		Q:    title,
	}
	jsonReq, _ := json.Marshal(req)

	var res models.CyberleninkaResponse
	body, err := c.client.Post(endpoint, jsonReq)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &models.LiteratureList{
		Title:       title,
		IsContained: res.Found > 0,
	}, nil
}
