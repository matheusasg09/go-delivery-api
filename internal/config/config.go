package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Carrega as configurações do .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	viper.AutomaticEnv()
}

// Retorna uma variável específica
func Get(key string) string {
	return viper.GetString(key)
}