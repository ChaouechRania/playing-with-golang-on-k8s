package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetHome handles request on the API home.
func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"playing-with-golang-on-k8s-api": "Hello Paris."})
}
