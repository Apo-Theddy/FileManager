package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfiguration struct {
	User     string
	Password string
	Database string
	Host     string
	Port     int
}

func (dc DatabaseConfiguration) Get() (DatabaseConfiguration, error) {
	dc.loadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return dc.defaultValues(), fmt.Errorf("error al convertir el puerto a un entero: %w, se establecerán valores por defecto", err)
	}
	return DatabaseConfiguration{
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
		Host:     dbHost,
		Port:     port,
	}, nil
}

func (DatabaseConfiguration) loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error cargando el archivo .env:", err)
		return
	}
}

func (DatabaseConfiguration) defaultValues() DatabaseConfiguration {
	return DatabaseConfiguration{
		User:     "sa",
		Password: "5q1@S3rv3r_s3cur3",
		Database: "file_manager",
		Host:     "localhost",
		Port:     1433,
	}
}
