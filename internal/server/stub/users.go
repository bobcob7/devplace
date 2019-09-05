package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserByID will get a user object by user ID
func (s Stub) GetUserByID(c *gin.Context) {
	c.String(http.StatusOK, "GetUserByID")
}

// GetUsers will get a list of user objects
func (s Stub) GetUsers(c *gin.Context) {
	c.String(http.StatusOK, "GetUsers")
}
