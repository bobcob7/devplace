package server

import (
	"github.com/gin-gonic/gin"
)

// UserHandles describe how the user related APIs will be implemented.
type UserHandles interface {
	// GetUserByID will get a user object by user ID
	GetUserByID(c *gin.Context)

	// GetUsers will get a list of user objects
	GetUsers(c *gin.Context)
}

// WorkspaceHandles describe how the workspace related APIs will be implemented.
type WorkspaceHandles interface {
	// GetWorkspaceByID will get a workspace object by workspace ID
	GetWorkspaceByID(c *gin.Context)

	// GetWorkspaces will get a list of workspace objects
	GetWorkspaces(c *gin.Context)
}

// ServiceHandles describe how the service related APIs will be implemented.
type ServiceHandles interface {
	// GetWorkspaceServiceByID will get a service in a workspace object by workspace and service IDs
	GetWorkspaceServiceByID(c *gin.Context)

	// GetWorkspaceServices will get a list of service objects in a workspace by workspace ID
	GetWorkspaceServices(c *gin.Context)
}

// WorkspaceTemplateHandles describe how the workspace template related APIs will be implemented.
type WorkspaceTemplateHandles interface {
	// GetWorkspaceTemplateByID will get a workspace template object by workspace template ID
	GetWorkspaceTemplateByID(c *gin.Context)

	// GetWorkspaceTemplates will get a list of workspace template objects
	GetWorkspaceTemplates(c *gin.Context)
}

// ServiceTemplateHandles describe how the service template related APIs will be implemented.
type ServiceTemplateHandles interface {
	// GetServiceTemplateByID will get a service template object by service template ID
	GetServiceTemplateByID(c *gin.Context)

	// GetServiceTemplates will get a list of service template objects
	GetServiceTemplates(c *gin.Context)
}

// Handles describe how the the APIs will be implemented.
type Handles interface {
	UserHandles
	WorkspaceHandles
	ServiceHandles
	WorkspaceTemplateHandles
	ServiceTemplateHandles
}
