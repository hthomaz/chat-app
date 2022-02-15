package chatApp

import (
	"fmt"
	"strings"
)

type ChatMessage struct {
	Username    string `json:"username"`
	Text        string `json:"text"`
	Color       string `json:"color"`
	Destination string `json:"destination"`
}

var possibleColors = []string{"blue", "black", "green"}
var currentColor = "black"

func dealWithCommandMsg(msg *ChatMessage) {

	text := msg.Text
	if strings.HasPrefix(text, "/whisp_") {
		MsgParts := strings.Split(text, " ")
		msg.Destination = strings.TrimPrefix(MsgParts[0], "/whisp_")
		msg.Text = strings.TrimPrefix(msg.Text, MsgParts[0])
		msg.Text = fmt.Sprintln(msg.Text, "(whisp)")

	} else if strings.HasPrefix(text, "/flood") {
		msg.Text = strings.TrimPrefix(msg.Text, "/flood ")
		for i := 0; i < 2; i++ {
			StoreInRedis(*msg)
			sendMessageToClients(*msg)
		}

	} else if strings.HasPrefix(text, "/color") {
		newColor := strings.TrimPrefix(msg.Text, "/color ")
		if StringInSlice(newColor, possibleColors) {
			currentColor = newColor
			msg.Color = currentColor
			msg.Text = fmt.Sprintln("Color changed to", newColor)
		} else {
			msg.Text = "Color not Avaliable"
			msg.Username = "System"
		}
	} else if strings.HasPrefix(text, "/count") {
		msg.Text = fmt.Sprintf("Number of active users in chat is: %d", GetNumberOfClients())
		msg.Destination = msg.Username
		msg.Username = "System"
	}
}
