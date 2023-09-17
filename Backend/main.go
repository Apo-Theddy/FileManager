package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/apo-theddy/file-manager/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	// router := echo.New()
	// apiServer := server.ServerManager{
	// 	Router:   router,
	// 	Entities: []interface{}{models.File{}},
	// }

	// apiServer.Start(1111)

	// sigChannel := make(chan os.Signal, 1)
	// signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	// go terminateRouterProcess(sigChannel, router)

	lfm := handlers.LocalFileManipulation{}
	lfm.New()

	lfm.GetFolderContent("test")

	time.Sleep(time.Second * 5)
}

func terminateRouterProcess(sigChannel chan os.Signal, router *echo.Echo) {
	defer router.Close()
	if err := router.Shutdown(context.TODO()); err != nil {
		log.Println("Error al cerrar el servidor Echo:", err)
	}
	log.Println("Servidor cerrado")
}
