package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	h.init(mediaServerFirstPort, mediaServerInstances)
	e := echo.New()
	e.POST("api/v1/room/", CreateRoom)
	e.GET("api/v1/room/:roomId", JoinRoom)
	e.DELETE("api/v1/room/:roomId", CancelRoom)
	e.Logger.Fatal(e.Start(":3000"))
}

func CreateRoom(c echo.Context) error {
	room, err := createRoom()
	if err != nil {
		status, err := resolveServiceError(err)("", "")
		return c.JSON(status, err)
	}
	return c.JSON(http.StatusCreated, room)
}

func JoinRoom(c echo.Context) error {
	roomID := c.Param("roomId")
	room, err := joinRoom(roomID)
	if err != nil {
		status, err := resolveServiceError(err)("roomId", roomID)
		return c.JSON(status, err)
	}
	return c.JSON(http.StatusOK, room)
}

func CancelRoom(c echo.Context) error {
	roomID := c.Param("roomId")
	err := deleteRoom(roomID)
	if err != nil {
		status, err := resolveServiceError(err)("roomId", roomID)
		return c.JSON(status, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
