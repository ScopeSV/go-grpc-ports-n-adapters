package httpServer

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/scopesv/todoGrpc/api/internal/ports"
)

type Adapter struct {
	*gin.Engine
	port int
	api  ports.APIPort
}

func (a *Adapter) StartServer() {
	createTodoRoutes(a.api, a.Group("/todos"))
	a.Run(fmt.Sprintf(":%d", a.port))
}

func NewAdapter(api ports.APIPort, httpPort int) *Adapter {
	return &Adapter{
		Engine: gin.Default(),
		port:   httpPort,
		api:    api,
	}
}
