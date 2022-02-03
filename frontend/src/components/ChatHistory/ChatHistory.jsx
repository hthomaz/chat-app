import React, { Component } from "react";
import Message from "../Message";
import "./ChatHistory.scss";

class ChatHistory extends Component {
  render() {
    // const messages = this.props.chatHistory?.map((msg, index) => {
    //     let data = JSON.parse(msg.data)
    //     const colorStyle = {color : data.color}
    //     return <p key={index} style={colorStyle}> <strong>{data.username}</strong>: {data.text}</p>
    // });
    const messages = this.props.chatHistory.map((msg, index) => <Message key = {index} message={msg} />);

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;