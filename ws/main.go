package main

import (
	"github.com/BUGLAN/kit/ms"
	"github.com/BUGLAN/pigeon/services/ws"
	"github.com/gin-gonic/gin"
)

func main() {
	server := ms.NewMicroService(
		ms.WithGinHTTP(handler),
		ms.WithPrometheus(),
	)
	server.ListenAndServer(5000)
}

func handler(engine *gin.Engine) {
	ctrl := ws.NewController()
	engine.GET("/ws", ctrl.ChatHandler)
}
