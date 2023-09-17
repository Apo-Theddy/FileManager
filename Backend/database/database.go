package database

import (
	"fmt"

	"github.com/apo-theddy/file-manager/configs"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	Instance *gorm.DB
}

func Connect(dc configs.DatabaseConfiguration) (*Database, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
	gdb, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectarse a la base de datos: %w", err)
	}
	fmt.Println("Conectado correctamente a la base de datos")
	return &Database{Instance: gdb}, nil
}
