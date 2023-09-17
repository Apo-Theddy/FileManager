package routes

import (
	"github.com/apo-theddy/file-manager/handlers"
	"github.com/labstack/echo/v4"
)

type FileRoute struct {
	Router  *echo.Echo
	Handler handlers.FileHandler
}

func (fr FileRoute) Create() {
	// GETS
	fr.Router.GET("/media/files", fr.Handler.GetFilesHandler)
	fr.Router.GET("/media/files/:id", fr.Handler.GetFilesHandler)
	fr.Router.GET("/media/dirs", fr.Handler.GetFilesHandler)
	fr.Router.GET("/media/dirs/:id", fr.Handler.GetFilesHandler)

	// POSTS

}
