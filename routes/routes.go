package routes

import (
	"github.com/daniel/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorID)
	r.PATCH("/alunos", controllers.EditaAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
