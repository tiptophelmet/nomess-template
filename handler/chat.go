package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/logger"
)

func Chat(ws *websocket.Conn) {
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Err(fmt.Sprintf("failed to read a message: %v", err.Error()))
			break
		}

		logger.Info(fmt.Sprintf("reading a message: %v", message))

		err = ws.WriteMessage(mt, message)
		if err != nil {
			logger.Err(fmt.Sprintf("failed to write a message: %v", err.Error()))
			break
		}
	}
}
