package controller

import (
	"hacktiv8_fp_1/common"
	"hacktiv8_fp_1/entity"
	"hacktiv8_fp_1/helpers"
	"hacktiv8_fp_1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

type TodosController interface {
	GetTodos(ctx *gin.Context)
	CreateNewTodo(ctx *gin.Context)
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

// func (c *todosController) GetTodoById (ctx *gin.Context) {
// 	params := ctx.Param("id");
// 	result, err := c.todosService.GetTodoById(ctx.Request.Context())
// }

func (c *todosController) CreateNewTodo(ctx *gin.Context) {

	contentType := helpers.GetContentType(ctx)
	todo := entity.Todos{}
	if contentType == appJson {
		ctx.ShouldBindJSON(&todo)
	} else {
		ctx.ShouldBind(&todo)
	}

	result, err := c.todosService.CreateTodo(ctx.Request.Context()) // testing aja ini
	if err != nil {
		res := common.BuildErrorResponse("Bad create request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusCreated, common.BuildResponse(true, "Success", result))
}
