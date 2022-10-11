package route

import (
	"hacktiv8_fp_1/controller"

	"github.com/gin-gonic/gin"
)

func TodosRoute(router *gin.Engine, todosController controller.TodosController) {
	todosRoutes := router.Group("/todos")
	{
		todosRoutes.GET("", todosController.GetTodos)
		todosRoutes.POST("", todosController.CreateNewTodo)
	}
}
