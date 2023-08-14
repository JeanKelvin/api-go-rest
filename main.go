package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "teste1", Historia: "teste 1"},
		{Id: 2, Nome: "teste2", Historia: "teste 2"},
	}

	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
