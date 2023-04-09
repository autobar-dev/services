package handlers

import (
	"fmt"

	"github.com/autobar-dev/services/modulerealtime/types"
	socketio "github.com/googollee/go-socket.io"
)

func OnDisconnect(s socketio.Conn, reason string) {
	app_ctx := s.Context().(*types.AppContext)

	fmt.Println("Disconnect")

	app_ctx.Log.Info("Socket disconnected", "session_id", s.ID())
}
