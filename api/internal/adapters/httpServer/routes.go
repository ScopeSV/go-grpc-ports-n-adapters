package httpServer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scopesv/todoGrpc/api/internal/application/domain"
	"github.com/scopesv/todoGrpc/api/internal/ports"
)

func createTodoRoutes(api ports.APIPort, todoRoute *gin.RouterGroup) {
	todoRoute.GET("/", func(c *gin.Context) {
		todos, err := api.GetAllTodos()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	todoRoute.POST("/", func(c *gin.Context) {
		var todoReq domain.TodoRequest

		if err := c.BindJSON(&todoReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := api.SaveTodo(&todoReq)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo created"})
	})

	todoRoute.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")

		i, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = api.DeleteTodo(i)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
	})

	todoRoute.PATCH("/:id/complete", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = api.SetTodoCompleted(i)
	})
}
