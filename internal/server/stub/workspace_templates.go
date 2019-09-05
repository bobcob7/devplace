package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWorkspaceTemplateByID will get a workspace template object by workspace template ID
func (s Stub) GetWorkspaceTemplateByID(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaceTemplateByID")
}

// GetWorkspaceTemplates will get a list of workspace template objects
func (s Stub) GetWorkspaceTemplates(c *gin.Context) {
	c.String(http.StatusOK, "GetWorkspaceTemplates")
}
