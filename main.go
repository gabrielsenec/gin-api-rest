package main

import (
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gabriel Costa", CPF: "12312312312", RG: "54270998-50"},
		{Nome: "João", CPF: "59859859859", RG: "9000000"},
		{Nome: "Isabela", CPF: "67867867867", RG: "34423423"},
	}
	routes.HandleRequests()
}
