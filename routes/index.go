package routes

import (
	"net/http"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/service"

	"github.com/gin-gonic/gin"
)

//IndexActions represents job postings controller actions
type IndexActions struct {
	index *service.Index
}

//NewIndexActions create a new UserActions
func NewIndexActions(
	index *service.Index,
) *IndexActions {
	return &IndexActions{
		index: index,
	}
}

//IndexOne indexes a document
func (as IndexActions) IndexOne(c *gin.Context) {
	//product := &store.Product{}
	req := new(api.ProUpsertRequest)
	err := as.index.IndexOne(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Products indexed"})
	return

}

//IndexAll handles indexations
func (is IndexActions) IndexAll(c *gin.Context) {
	err := is.index.IndexProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Products indexed"})
	return
}
