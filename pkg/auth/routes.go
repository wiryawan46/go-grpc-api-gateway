package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wiryawan46/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/wiryawan46/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (i *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, i.Client)
}

func (i *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, i.Client)
}
