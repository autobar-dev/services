package handlers

import (
	"fmt"

	"github.com/autobar-dev/services/modulerealtime/types"
	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(s socketio.Conn, app_context *types.AppContext) error {
	s.SetContext(app_context)

	fmt.Println("New connection")

	return nil
}
