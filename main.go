package main

import (
	"cinema_digital_go/api/app/permissao/repository"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	carregaDadosIniciais()
	iniciaConfigBanco()
	configuraPermissoes()

	iniciaRotasAPI()
}

func carregaDadosIniciais() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func iniciaConfigBanco() {
	dbConection.ConnectDatabase()
	err := dbConection.RunMigrations()
	if err != nil {
		return
	}
}

func configuraPermissoes() {
	if repository.NewPermissaoRepository(dbConection.DB).GerenciaPermissoes() != nil {
		log.Fatalf("erro ao configurar permissões do sistema")
	}
}

func iniciaRotasAPI() {
	router := gin.Default()

	router.RemoveExtraSlash = true
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupGeneroRoutes(router, dbConection.DB)

	router = routes.SetupRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
