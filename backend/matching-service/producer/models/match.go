package models

import "github.com/labstack/echo/v4"

type MessageQueueRequestPacket struct {
	Ctx         echo.Context `json:"ctx"`
	RequestBody MatchRequest `json:"request_body"`
}

type MessageQueueResponsePacket struct {
}
