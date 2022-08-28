package http

import (
	"github.com/guil95/split-entrypoints/internal/users/infra/server/http/server"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func RunHTTPServer(quit chan error) {
	handler := gin.Default()

	s := server.NewHTTPServer(handler)
	s.Api()

	log.Println("Running http server on :8080 port")

	if err := handler.Run(":8080"); err != nil {
		quit <- err
	}
}
