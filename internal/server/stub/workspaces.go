package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWorkspaceByID will get a workspace object by workspace ID
func (s Stub) GetWorkspaceByID(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaceByID")
}

// GetWorkspaces will get a list of workspace objects
func (s Stub) GetWorkspaces(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaces")
}
