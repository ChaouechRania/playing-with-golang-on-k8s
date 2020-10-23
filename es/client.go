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
func (c *Client) Search(index string) ([]*api.Product, error) {
	query := elastic.NewMatchAllQuery()
	searchResult, err := c.client.Search().
		Index(index).
		Type("_doc").
		Query(query).
		Do(c.ctx)
	if err != nil {
		return []*api.Product{}, err
	}
	var job api.Product
	jobs := make([]*api.Product, 0)
	for _, item := range searchResult.Each(reflect.TypeOf(job)) {
		j := item.(api.Product)
		jobs = append(jobs, &j)
	}
	return jobs, nil
}

//GetJob a job from index
func (c *Client) GetProduct(index string, id string) (*api.Product, error) {
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
	var job api.Product
	err = json.Unmarshal(getJob.Source, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

//IndexLocations indexes locations in bulk.
// func (c *Client) IndexLocations(locations []store.Product) (int, error) {
// 	bulkRequest := c.cfg.EsClient.Bulk()

// 	for _, location := range locations {
// 		bulkRequest.Add(
// 			elastic.NewBulkIndexRequest().
// 				Index(c.cfg.IndexLocations.Name).
// 				Type(c.cfg.IndexLocations.Type).
// 				Id(location.Slug).
// 				Doc(location))
// 	}

// 	if bulkRequest.NumberOfActions() != len(locations) {
// 		return 0, fmt.Errorf("upsertBatch: number of bulkable documents is %v instead of %v", bulkRequest.NumberOfActions(), len(locations))
// 	}

// 	bulkResponse, err := bulkRequest.Do(c.ctx)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return len(bulkResponse.Indexed()), nil
// }
