package service

import (
	"context"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/es"
	"playing-with-golang-on-k8s/store"
)

const (
	proIndex = "products"
)

// Create a new index.
const proMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"name":{
				"type":"text"
			},
			"description":{
				"type":"text"
			}
		}
	}
}`

//Index service
type Index struct {
	es           *es.Client
	productStore store.ProductStore
}

//NewIndex create a new UserActions
func NewIndex(
	es *es.Client,
	productStore store.ProductStore,
) *Index {
	return &Index{
		es:           es,
		productStore: productStore,
	}
}

//IndexJobs indexes all jobs
func (i *Index) IndexProducts(ctx context.Context) error {
	err := i.es.CreateIndex(proMapping, proIndex)
	if err != nil {
		return err
	}

	products, err := i.productStore.List(ctx, 0, 10)
	println(len(products))

	if err != nil {
		return err
	}

	for _, product := range products {
		println(product.Name)
		ij := api.ProductStore(&product)
		i.es.Index(proIndex, product.ID, ij)
		println(ij.Name)
		/*if err != nil {
			i.es.Index(proIndex, product.ID, ij)
			println(product.ID)
		}*/
	}
	return nil
}

//SearchJobs search jobs
func (i *Index) SearchProducts(ctx context.Context) (*api.ProductsResults, error) {
	products, err := i.es.Search(proIndex)
	if err != nil {
		return nil, err
	}

	return &api.ProductsResults{Products: products}, nil
}

//GetJob search jobs
func (i *Index) GetProduct(ctx context.Context, id string) (*api.Product, error) {
	job, err := i.es.GetProduct(proIndex, id)
	if err != nil {
		return nil, err
	}

	return job, nil
}
