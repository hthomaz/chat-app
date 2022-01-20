package chatApp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var rdb *redis.Client
var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan ChatMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MiddleWare = func(*websocket.Conn, ChatMessage) ChatMessage

func Obverser(redisURL, port string) {
	err := ConnectToDataBase(redisURL)
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/websocket", handleConnections)
	go handleMessages()
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}

func ConnectToDataBase(redisURL string) error {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return err
	}
	rdb = redis.NewClient(opt)
	return err
}

func handleMessages() {
	for {
		msg := <-broadcaster
		dealWithCommandMsg(&msg)
		StoreInRedis(msg)
		sendMessageToClients(msg)
	}
}

func jsonMiddleWare(ws *websocket.Conn, _ ChatMessage) ChatMessage {
	var msg ChatMessage
	ws.ReadJSON(&msg)

	return msg
}

func setColorMiddleWare(_ *websocket.Conn, msg ChatMessage) ChatMessage {
	msg.Color = currentColor
	return msg
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

	middleWares := []MiddleWare{jsonMiddleWare, setColorMiddleWare}

	for {
		var msg ChatMessage

		for _, middleWare := range middleWares {
			msg = middleWare(ws, msg)
		}
		// if err != nil {
		// 	delete(clients, ws)
		// 	break
		// }

		broadcaster <- msg
	}
}

func GetNumberOfClients() int {
	return len(clients)
}

func StoreInRedis(msg ChatMessage) {
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
