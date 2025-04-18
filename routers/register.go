package routers

import (
	"github.com/gin-gonic/gin"
)

// ViewSetRouter registers a model's ViewSet with RESTful endpoints.
// Accepts a *gin.RouterGroup to allow grouping of routes under middleware.
func ViewSetRouter(routerGroup *gin.RouterGroup, path string, viewSet interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Retrieve(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}) {
	// Register RESTful endpoints within the RouterGroup
	routerGroup.GET(path, viewSet.List)
	routerGroup.POST(path, viewSet.Create)
	routerGroup.GET(path+"/:id", viewSet.Retrieve)
	routerGroup.PUT(path+"/:id", viewSet.Update)
	routerGroup.DELETE(path+"/:id", viewSet.Delete)
}
