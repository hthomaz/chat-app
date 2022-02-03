import React, { Component } from "react";
import "./ChatHistory.scss";

class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory?.map((msg, index) => {
        let data = JSON.parse(msg.data)
        const colorStyle = {color : data.color}

        return <p key={index} style={colorStyle}> <strong>{data.username}</strong>: {data.text}</p>
    });

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;