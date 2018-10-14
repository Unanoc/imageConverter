package main

import (
	"imageConverter/logger"
	"imageConverter/router"
	"log"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const (
	port = ":8000"
	addr = "localhost"
)

func main() {
	logger.Logger.Info("Starting server",
		zap.String("host", addr),
		zap.String("port", port),
	)
	router := router.NewRouter()
	log.Fatal(fasthttp.ListenAndServe(port, logger.LoggerHandler(router.Handler)))
}