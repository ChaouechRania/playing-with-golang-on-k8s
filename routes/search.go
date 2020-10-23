package routes

import (
	"net/http"
	"playing-with-golang-on-k8s/service"

	"github.com/gin-gonic/gin"
)

//SearchActions represents job postings search actions
type SearchActions struct {
	index *service.Index
}

//NewSearchActions search actions
func NewSearchActions(
	index *service.Index,
) *SearchActions {
	return &SearchActions{
		index: index,
	}
}

//IndexJobs handles job search
func (sa SearchActions) IndexProducts(c *gin.Context) {
	err := sa.index.IndexProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jobs indexed"})
	return
}

//Jobs handles job search
func (sa SearchActions) Products(c *gin.Context) {
	jobs, err := sa.index.SearchProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
	return
}

//GetJob retrieve a job detail
func (sa SearchActions) GetProduct(c *gin.Context) {
	id := c.Param("id")
	job, err := sa.index.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
	return
}
