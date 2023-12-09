package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	//paginas para serem renderizadas
	r.LoadHTMLGlob("templates/*")
	//carregar arquivos estáticos como css por exemplo
	r.Static("/styles", "./styles")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.DetalhaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/rg/:rg", controllers.BuscaAlunoPorRG)
	r.GET("/index", controllers.ExibePaginaIndex)
	//usamos No Routes quando não tem nenhum caminho para a requisição, nosso famoso 404
	r.NoRoute(controllers.RotaNaoEncontra)
	r.Run(":8000")

}
