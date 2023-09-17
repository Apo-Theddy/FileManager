package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/apo-theddy/file-manager/configs"
	"github.com/apo-theddy/file-manager/database"
	"github.com/apo-theddy/file-manager/handlers"
	"github.com/apo-theddy/file-manager/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServerManager struct {
	Router   *echo.Echo
	Db       *gorm.DB
	Entities []interface{}
}

func (sm *ServerManager) Start(port int) {
	db, err := sm.initializeDatabase()
	if err != nil {
		log.Println("Error al iniciar la base de datos: ", err)
	}
	sm.Db = db
	sm.initializeRoutes()

	entityChannel := make(chan interface{})

	wg := &sync.WaitGroup{}
	wg.Add(len(sm.Entities))

	go sm.sendEntities(entityChannel, wg)
	go sm.processEntities(entityChannel, wg)
	log.Println(sm.Router.Start(fmt.Sprintf(":%d", port)))
	wg.Wait()
}

func (sm ServerManager) processEntities(entityChannel <-chan interface{}, wg *sync.WaitGroup) {
	for entity := range entityChannel {
		err := sm.Db.AutoMigrate(entity)
		if err != nil {
			log.Println("ocurrio un error al migrar las base de datos")
		}
	}
}

func (sm ServerManager) sendEntities(entityChannel chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entity := range sm.Entities {
		entityChannel <- entity
	}
}

func (ServerManager) initializeDatabase() (*gorm.DB, error) {
	config, err := configs.DatabaseConfiguration{}.Get()
	if err != nil {
		return nil, err
	}
	db, err := database.Connect(config)
	if err != nil {
		return nil, err
	}
	return db.Instance, nil
}

func (sm ServerManager) initializeRoutes() {
	// FileRoutes
	fr := routes.FileRoute{Router: sm.Router, Handler: handlers.FileHandler{Db: sm.Db}}
	fr.Create()
}
