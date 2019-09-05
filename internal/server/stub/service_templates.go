package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetServiceTemplateByID will get a service template object by service template ID
func (s Stub) GetServiceTemplateByID(c *gin.Context) {
	c.String(http.StatusOK, "GetServiceTemplateByID")
}

// GetServiceTemplates will get a list of service template objects
func (s Stub) GetServiceTemplates(c *gin.Context) {
	c.String(http.StatusOK, "GetServiceTemplates")
}
