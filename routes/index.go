package routes

import (
	"net/http"
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
