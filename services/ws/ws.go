package ws

import (
	"github.com/BUGLAN/kit/logutil"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	upGrader *websocket.Upgrader
	logger   zerolog.Logger
}

func NewController() *Controller {
	return &Controller{
		logger: logutil.NewLogger("controller", "ws"),
		upGrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ctrl *Controller) ChatHandler(ctx *gin.Context) {
	c, err := ctrl.upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctrl.logger.Panic().Err(err).Msg("websocket 连接断开")
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil && !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			ctrl.logger.Panic().Err(err).Msg("读取消息失败")
			break
		} else if err != nil {
			ctrl.logger.Info().Err(err).Msg("客户端主动关闭连接")
			return
		}

		ctrl.logger.Info().Msgf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		ctrl.logger.Info().Msgf("send: %s", message)
		if err != nil {
			ctrl.logger.Panic().Err(err).Msg("发送消息失败")
			break
		}
	}
}
