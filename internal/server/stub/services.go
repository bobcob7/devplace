package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWorkspaceServiceByID will get a service object in a workspace by workspace and service ID
func (s Stub) GetWorkspaceServiceByID(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaceServicesByID")
}

// GetWorkspaceServices will get a list of service objects in a workspace by workspace ID
func (s Stub) GetWorkspaceServices(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaceServices")
}
