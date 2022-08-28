package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	handler *gin.Engine
}

func NewHTTPServer(handler *gin.Engine) *httpServer {
	return &httpServer{handler}
}

func (server httpServer) Api() {
	server.handler.GET("/", func(ctx *gin.Context) {
		server.listUsers(ctx)
	})
}

func (server httpServer) listUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "list users")
}
