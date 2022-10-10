package controller

import (
	"hacktiv8_fp_1/common"
	"hacktiv8_fp_1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodosController interface {
	GetTodos(ctx *gin.Context)
}

type todosController struct {
	todosService service.TodosService
}

func NewTodosController(ts service.TodosService) TodosController {
	return &todosController{
		todosService: ts,
	}
}

// @BasePath /todos

// @Summary Gets all todos
// @Description Gets all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {Todos} todos
// @Router /todos [get]
func (c *todosController) GetTodos(ctx *gin.Context) {
	result, err := c.todosService.GetTodos(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get todo list", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *todosController) CreateNewTodo(ctx *gin.Context){
	
}
