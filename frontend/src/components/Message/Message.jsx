import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    let msgJSON = JSON.parse(this.props.message.data);
    const username = this.props.username
    //console.log(msgJSON)
    this.state = {
      message: msgJSON,
      username : username
    };
  }

  render() {
    const colorStyle = {color : this.state.message.color}
    if (this.state.message.destination === this.state.username || this.state.message.destination === 'all' || this.state.message.username === this.state.username) {
      return (
        <div className="Message" style={colorStyle} > 
          <strong>{this.state.message.username}</strong>: {this.state.message.text}
        </div>)
    } else {
      return null
    }
  }
}

export default Message;