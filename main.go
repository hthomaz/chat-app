package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	rdb *redis.Client
)

type ChatMessage struct {
	Username    string `json:"username"`
	Text        string `json:"text"`
	Color       string `json:"color"`
	Destination string `json:"destination"`
}

var possibleColors = []string{"blue", "black", "green"}
var currentColor = "black"
var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan ChatMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	//Load enviroments data
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/websocket", handleConnections)
	go handleMessages()

	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true

	if rdb.Exists("chat_messages").Val() != 0 {
		sendPreviousMessages(ws)
	}

	for {
		var msg ChatMessage

		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcaster <- msg
	}
}

func sendPreviousMessages(ws *websocket.Conn) {
	chatMessages, err := rdb.LRange("chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	for _, chatMessage := range chatMessages {
		var msg ChatMessage
		json.Unmarshal([]byte(chatMessage), &msg)
		sendMessageToClient(ws, msg)
	}
}

func handleMessages() {
	for {
		msg := <-broadcaster
		dealWithCommandMsg(&msg)
		msg.Color = currentColor
		storeInRedis(msg)
		sendMessageToClients(msg)
	}
}

func dealWithCommandMsg(msg *ChatMessage) {

	text := msg.Text
	if strings.HasPrefix(text, "/whisp_") {
		MsgParts := strings.Split(text, " ")
		msg.Destination = strings.Trim(MsgParts[0], "/whisp_")
		msg.Text = strings.Trim(msg.Text, MsgParts[0])
		msg.Text = fmt.Sprintln(msg.Text, "(whisp)")
	} else if strings.HasPrefix(text, "/flood") {
		msg.Text = strings.Trim(msg.Text, "/flood")
		for i := 0; i < 2; i++ {
			storeInRedis(*msg)
			sendMessageToClients(*msg)
		}
	} else if strings.HasPrefix(text, "/color") {
		newColor := strings.Trim(msg.Text, "/color ")
		if stringInSlice(newColor, possibleColors) {
			currentColor = newColor
			msg.Text = fmt.Sprintln("Color changed to", newColor)
		} else {
			msg.Text = "Color not Avaliable"
		}

	}
}

func storeInRedis(msg ChatMessage) {
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	if err := rdb.RPush("chat_messages", json).Err(); err != nil {
		panic(err)
	}
}

func sendMessageToClients(msg ChatMessage) {
	for client := range clients {

		sendMessageToClient(client, msg)
	}
}

func sendMessageToClient(client *websocket.Conn, msg ChatMessage) {
	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error :%v", err)
		client.Close()
		delete(clients, client)
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
