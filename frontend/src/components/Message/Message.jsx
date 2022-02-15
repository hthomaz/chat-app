import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    let msgJSON = JSON.parse(this.props.message.data);
    console.log(msgJSON)
    this.state = {
      message: msgJSON
    };
  }

  render() {
    const colorStyle = {color : this.state.message.color}
    return <div className="Message" style={colorStyle} > <strong>{this.state.message.username}</strong>: {this.state.message.text}</div>;
  }
}

export default Message;