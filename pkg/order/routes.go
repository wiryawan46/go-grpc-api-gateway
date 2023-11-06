package order

import (
	"github.com/gin-gonic/gin"
	"github.com/wiryawan46/go-grpc-api-gateway/pkg/auth"
	"github.com/wiryawan46/go-grpc-api-gateway/pkg/config"
	"github.com/wiryawan46/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRouter(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerGroup := r.Group("/order")
	routerGroup.Use(a.AuthRequired)
	routerGroup.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
