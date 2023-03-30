package handlers

import (
	"github.com/autobar-dev/services/modulerealtime/types"
	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(context *types.AppContext, socket socketio.Conn) error {
	logger := *(*context).AppLogger

	headers := socket.RemoteHeader()

	if authorization_header := headers.Get("Authorization"); authorization_header != "" {
		logger.Info("Socket connected", "authorization_header", authorization_header)
		return nil
	}

	logger.Debug("Socket connected", "socket_id", socket.ID())

	socket.SetContext("")

	return nil
}
