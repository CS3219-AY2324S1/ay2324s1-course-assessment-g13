package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	// "fmt"
    "log"
	"github.com/google/uuid"
    "github.com/gorilla/websocket"
)

type Room struct {
    Id string
    Users map[string]*User
}

type Hub struct {
    Rooms map[string]*Room
    Broadcast chan *Message
}

type User struct {
    Connection *websocket.Conn
    Message chan *Message
    RoomId string
    Username string
}

type Message struct {
    Content string
    RoomId string
    Username string
}

type Handler struct {
    hub *Hub
}

func NewHandler(h *Hub) *Handler {
    return &Handler{
        hub: h,
    }
}

func NewHub() *Hub {
    return &Hub{
        Rooms: make(map[string]*Room),
        Broadcast: make (chan *Message, 5),
    }
}

func (h *Handler) CreateRoom(c echo.Context) error {
    id := uuid.New().String()
    h.hub.Rooms[id] = &Room{
        Id: id,
        Users: make(map[string]*User),
    }

    return c.JSON(http.StatusOK, h.hub.Rooms[id])
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func (h *Handler) JoinRoom(c echo.Context) error {
    connection, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        log.Printf("error upgrade: %s\n", err.Error())
        return c.JSON(http.StatusBadRequest, "Error joining room")
    }

    roomId := c.Param("roomId")
    user := &User{
        Connection: connection,
        Message: make(chan *Message, 10),
        RoomId: roomId,
        // TODO: get username from either token or param
        Username: "",
    }

    go user.writeMessage()
    user.readMessage(h.hub)
    return c.JSON(http.StatusOK, "Successfully joined room!")
}

func (user *User) writeMessage() {
    defer func() {
        user.Connection.Close()
    }()

    for {
        message, ok := <-user.Message
        if !ok {
            return
        }

        user.Connection.WriteJSON(message)
    }
}

func (user *User) readMessage(hub *Hub) {
    for {
        _, message, err := user.Connection.ReadMessage()
        if err != nil {
            log.Println("Error reading message from socker")
            break
        }

        msg := &Message{
            Content: string(message),
            RoomId: user.RoomId,
            Username: "",
        }

        hub.Broadcast <- msg
    }
}

func (hub *Hub) Run() {
    for {
        select {
        case message := <- hub.Broadcast:
            if _, ok := hub.Rooms[message.RoomId]; ok {
                for _, user := range hub.Rooms[message.RoomId].Users {
                    user.Message <- message
                }
            }
        }
    }
}


