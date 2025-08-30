package ping_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	msg string
}

func NewPingController() PingController {
	return PingController{
		msg: "Pong!",
	}
}

func (c PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": c.msg,
	})
}
