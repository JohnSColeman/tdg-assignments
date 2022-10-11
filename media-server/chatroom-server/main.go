package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Message string
	Error   error
}

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	// health check
	router.GET("/", func(c *gin.Context) {
		response, err := rdb.Ping(ctx).Result()
		if err != nil {
			c.JSON(http.StatusBadGateway, ErrorResponse{
				Message: err.Error(),
				Error:   err,
			})
		} else if response != "PONG" {
			c.JSON(http.StatusBadGateway, ErrorResponse{
				Message: "NO PONG",
				Error:   nil,
			})
		} else {
			c.JSON(http.StatusOK, "ok")
		}
	})

	// start a new room
	router.GET("/room/:roomId", func(c *gin.Context) {
		validateRequest(c,
			func(roomId string) { c.HTML(http.StatusOK, "index.html", nil) })
	})

	// websocket messages
	router.GET("/ws/:roomId", func(c *gin.Context) {
		validateRequest(c,
			func(roomId string) { serveWs(c.Writer, c.Request, roomId) })
	})

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func validateRequest(c *gin.Context, success func(roomId string)) {
	roomId := c.Param("roomId")
	status, err := validateRoomId(roomId)
	if err != nil {
		c.JSON(status, ErrorResponse{
			Message: err.Error(),
			Error:   err,
		})
	} else {
		success(roomId)
	}
}

func validateRoomId(roomId string) (status int, error error) {
	err := verifyRoomId(roomId)
	if err != nil {
		switch err {
		case invalidRoomError:
		case incorrectServerError:
			return http.StatusForbidden, err
		default:
			return http.StatusBadGateway, err
		}
	}
	return http.StatusOK, nil
}
