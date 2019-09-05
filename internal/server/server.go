package server

import (
	"github.com/gin-gonic/gin"
)

// Start will start an http server with an implementation of the Handles interface.
func Start(address string, handlers Handles) {
	router := gin.Default()
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// Users
			v1.GET("/user/:userid", handlers.GetUserByID)
			v1.GET("/user/", handlers.GetUsers)

			// Workspaces
			v1.GET("/workspace/:workspaceid", handlers.GetWorkspaceByID)
			v1.GET("/workspace/", handlers.GetWorkspaces)

			// Services
			v1.GET("/workspace/:workspaceid/service/:serviceid", handlers.GetWorkspaceServiceByID)
			v1.GET("/workspace/:workspaceid/service/", handlers.GetWorkspaceServices)

			// Workspace Templates
			v1.GET("/template/workspace/:workspacetemplateid", handlers.GetWorkspaceTemplateByID)
			v1.GET("/template/workspace/", handlers.GetWorkspaceTemplates)

			// Service Templates
			v1.GET("/template/service/:servicetemplateid", handlers.GetServiceTemplateByID)
			v1.GET("/template/service/", handlers.GetServiceTemplates)
		}
	}
	router.Run(address)
}
