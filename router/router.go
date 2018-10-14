package router

import (
	"imageConverter/handler"

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/", handler.MainHandler)
	router.POST("/api/convert", handler.ConvertHandler)
	return router
}
