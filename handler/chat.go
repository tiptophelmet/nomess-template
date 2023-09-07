package handler

import (
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess-template/internal/logger"
)

func Chat(ws *websocket.Conn) {
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Error("failed to read a message: %v", err.Error())
			break
		}

		logger.Info("reading a message: %v", message)

		err = ws.WriteMessage(mt, message)
		if err != nil {
			logger.Error("failed to write a message: %v", err.Error())
			break
		}
	}
}
