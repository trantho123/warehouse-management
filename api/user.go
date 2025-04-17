package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "createUser route hoạt động!",
	})
}
