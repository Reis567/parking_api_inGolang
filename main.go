package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Carregar variáveis do .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Usar as variáveis de ambiente
    appPort := os.Getenv("APP_PORT")
    dbURL := os.Getenv("DATABASE_URL")

    fmt.Printf("Aplicação rodando na porta %s\n", appPort)
    fmt.Printf("Conectando ao banco de dados em %s\n", dbURL)
}
