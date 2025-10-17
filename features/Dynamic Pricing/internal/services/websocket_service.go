package services

import (
	"VOID/internal/ws"
)

func SendOrderUpdate(manager *ws.Manager, orderID uint, message string) {
	payload := []byte(message)
	manager.BroadcastMessage(payload)
}
