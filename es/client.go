package es

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"playing-with-golang-on-k8s/api"
	"reflect"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

var (
	ErrIndexAlreadyExists = errors.New("index already exists")
)

//Client Elastic client
type Client struct {
	client *elastic.Client
	ctx    context.Context
}

//NewClient new elastic client
func NewClient(ctx context.Context) (*Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		ctx:    ctx,
		client: client,
	}, nil
}

//CreateIndex creates a new index
func (c *Client) CreateIndex(mapping string, index string) error {
	// Check if the index called "twitter" exists
	exists, err := c.client.IndexExists(index).Do(c.ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	_, err = c.client.CreateIndex(index).BodyString(mapping).Do(c.ctx)
	if err != nil {
		return err
	}
	return nil
}

//Index a doc
func (c *Client) Index(index string, id string, doc interface{}) error {
	_, err := c.client.Index().
		Index(index).
		Type("_doc").
		Id(id).
		BodyJson(doc).
		Do(c.ctx)

	if err != nil {
		return err
	}
	return nil
}

//Search docs
func (c *Client) Search(index string) ([]*api.Job, error) {
	query := elastic.NewMatchAllQuery()
	searchResult, err := c.client.Search().
		Index(index).
		Type("_doc").
		Query(query).
		Do(c.ctx)
	if err != nil {
		return []*api.Job{}, err
	}
	var job api.Job
	jobs := make([]*api.Job, 0)
	for _, item := range searchResult.Each(reflect.TypeOf(job)) {
		j := item.(api.Job)
		jobs = append(jobs, &j)
	}
	return jobs, nil
}

//GetJob a job from index
func (c *Client) GetJob(index string, id string) (*api.Job, error) {
	getJob, err := c.client.Get().
		Index(index).
		Type("_doc").
		Id(id).
		Do(c.ctx)
	if err != nil {
		return nil, err
	}

	if !getJob.Found {
		return nil, errors.New("not found")
	}
	var job api.Job
	err = json.Unmarshal(getJob.Source, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
