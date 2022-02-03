// App.js
import React, { Component } from "react";
import "./App.css";
import {connect, sendMsg } from "./api";
import Header from './components/Header'
import ChatHistory from "./components/ChatHistory";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    }
  }

  send() {
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <Header/>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory,msg]
      }))
      console.log(this.state);
    });
  }
}

export default App;