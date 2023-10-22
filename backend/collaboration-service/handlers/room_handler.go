package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
    "log"
    "os"
    "encoding/json"
	"github.com/google/uuid"
    "github.com/gorilla/websocket"
)

type Room struct {
    Id string
    Users map[string]*User
    QuestionId string
}

type Hub struct {
    Rooms map[string]*Room
    BroadcastChannel chan *Message
    DisconnectChannel chan *User
}

type User struct {
    Connection *websocket.Conn
    CodeChannel chan *Message
    ChatChannel chan *Message
    RoomId string
    UserId string
}

type Message struct {
    Content string
    RoomId string
    UserId string
    Type string 
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
        BroadcastChannel: make (chan *Message, 5),
        DisconnectChannel: make(chan *User),
    }
}

func (h *Handler) CreateRoom(c echo.Context) error {
    id := uuid.New().String()
    reqBody := make(map[string]string)
    err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
    if err != nil {
        log.Fatal(err)
    }

	resp, err := http.Get(os.Getenv("QUESTION_SERVICE_URL") + "/questions/complexity/" + reqBody["complexity"])
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    var questionId string
	err = json.NewDecoder(resp.Body).Decode(&questionId)
    if err != nil {
        log.Fatal(err)
    }

    h.hub.Rooms[id] = &Room{
        Id: id,
        Users: make(map[string]*User),
        QuestionId: questionId,
    }

    return c.JSON(http.StatusOK, h.hub.Rooms[id])
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c echo.Context) error {
    roomId := c.Param("roomId")
    if len(h.hub.Rooms[roomId].Users) >= 2 {
        return c.JSON(http.StatusBadRequest, "Error joining room")
    }

    connection, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        log.Printf("error upgrade: %s\n", err.Error())
        return c.JSON(http.StatusBadRequest, "Error joining room")
    }

    log.Println("Attempting to join room")
    
    userId := uuid.New().String()
    user := &User{
        Connection: connection,
        CodeChannel: make(chan *Message, 10),
        ChatChannel: make(chan *Message, 10),
        RoomId: roomId,
        UserId: userId,
    }

    h.hub.Rooms[roomId].Users[userId] = user

    go user.writeCodeMessage()
    go user.writeChatMessage()
    user.readMessage(h.hub)
    return c.JSON(http.StatusOK, "Successfully joined room!")
}


func (h *Handler) GetQuestionId(c echo.Context) error {
    roomId := c.Param("roomId")

    questionId := h.hub.Rooms[roomId].QuestionId
    if questionId == "" {
        return c.JSON(http.StatusBadRequest, "No question is allocated")
    }

    return c.JSON(http.StatusOK, questionId)
}

// Reads data from hub, and emit data to client
func (user *User) writeCodeMessage() {
    defer func() {
        user.Connection.Close()
    }()

    for {
        message, ok := <-user.CodeChannel
        if !ok {
            return
        }

        user.Connection.WriteJSON(message)
    }
}

func (user *User) writeChatMessage() {
    defer func() {
        user.Connection.Close()
    }()

    for {
        message, ok := <-user.ChatChannel
        if !ok {
            return
        }

        user.Connection.WriteJSON(message)
    }
}

// Reads the data from the client, and emit the data to the hub
func (user *User) readMessage(hub *Hub) {
    defer func() {
		hub.DisconnectChannel <- user
		user.Connection.Close()
	}()

    for {
        _, message, err := user.Connection.ReadMessage()
        if err != nil {
            log.Println("Error reading message from socket")
            break
        }

        var msg Message
        err = json.Unmarshal(message, &msg)
        if err != nil {
            log.Println(err)
            break
        }

        msg.RoomId = user.RoomId
        msg.UserId = user.UserId

        hub.BroadcastChannel <- &msg
    }
}

func (hub *Hub) Run() {
    for {
        select {
        case message := <- hub.BroadcastChannel:
            if _, ok := hub.Rooms[message.RoomId]; ok {
                for _, user := range hub.Rooms[message.RoomId].Users {
                    if message.UserId != user.UserId {
                        if message.Type == "code" {
                            user.CodeChannel <- message
                        } else {
                            user.ChatChannel <- message
                        }
                    }
                }
            }
        case user := <- hub.DisconnectChannel:
            delete(hub.Rooms[user.RoomId].Users, user.UserId)
            if len(hub.Rooms[user.RoomId].Users) == 0 {
                delete(hub.Rooms, user.RoomId)
            }
            close(user.CodeChannel)
            close(user.ChatChannel)
        }
    }
}
