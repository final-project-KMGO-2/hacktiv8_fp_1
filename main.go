package main

import (
	"hacktiv8_fp_1/controller"
	"hacktiv8_fp_1/repository"
	"hacktiv8_fp_1/route"
	"hacktiv8_fp_1/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		todosRepository repository.TodosRepository = repository.NewTodosRepository("filepath")
		todosService    service.TodosService       = service.NewTodosService(todosRepository)
		todosController controller.TodosController = controller.NewTodosController(todosService)
	)
	server := gin.Default()
	route.SwaggerRoute(server)
	route.TodosRoute(server, todosController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
