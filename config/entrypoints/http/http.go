package http

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guil95/split-entrypoints/internal/users/infra/server/http/server"
	"go.uber.org/zap"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func RunHTTPServer(quit chan os.Signal) {
	handler := gin.Default()

	s := server.NewHTTPServer(handler)
	s.Api()

	log.Println("Running http server on :8080 port")

	if err := handler.Run(":8080"); err != nil {
		zap.S().Errorf("Error on server run %v", err)
		<-quit
	}
}
