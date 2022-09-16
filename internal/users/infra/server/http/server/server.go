package server

import (
	"github.com/guil95/split-entrypoints/internal/users/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guil95/split-entrypoints/internal/users"
)

type httpServer struct {
	handler *gin.Engine
	uc      usecases.UseCase
}

func NewHTTPServer(handler *gin.Engine, uc usecases.UseCase) *httpServer {
	return &httpServer{handler, uc}
}

func (server httpServer) Api() {
	groupUsers := server.handler.Group("/users")
	groupUsers.GET("/", func(ctx *gin.Context) {
		server.listUsers(ctx)
	})
	groupUsers.POST("/", func(ctx *gin.Context) {
		server.saveUser(ctx)
	})
}

func (server httpServer) listUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "list users")
}

func (server httpServer) saveUser(ctx *gin.Context) {
	var user users.User

	err := ctx.BindJSON(&user)
	if err != nil {
		return
	}

	err = server.uc.Save(ctx, &user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
