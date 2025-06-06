package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusasg09/go-delivery-api/internal/config"
	"github.com/matheusasg09/go-delivery-api/internal/infra/db"
)

func main() {
	// Carrega configs
	config.LoadConfig()

	// Conecta ao banco
	db.Conectar()

	// Inicia o servidor
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API de Delivery rodando com sucesso!"})
	})

	porta := config.Get("PORT")
	log.Println("Servidor rodando na porta", porta)
	r.Run(":" + porta)
}