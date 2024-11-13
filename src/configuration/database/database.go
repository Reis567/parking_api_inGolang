package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"meu-novo-projeto/src/model"
)

var (
	DB *gorm.DB
)

func ConnectDatabase() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao banco de dados: %w", err)
	}

	// Migrar a estrutura do modelo para o banco de dados
	err = db.AutoMigrate(&model.UserDomain{})
	if err != nil {
		log.Fatalf("Erro ao fazer automigração do banco de dados: %v", err)
	}

	DB = db
	return db, nil
}
