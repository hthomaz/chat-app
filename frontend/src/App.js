// App.js
import React, { useEffect, useState } from "react";
import "./App.css";
import {connect, sendMsg } from "./api";
import Header from './components/Header'
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import LoginPage from "./components/LoginPage/LoginPage";


function App () {
  const [chatHistory, setChatHistory] = useState({messages : []});
  const [username, setUsername] = useState('');
  const [isUsernameDefined, setUsernameDefined] = useState(false);

  const sentUsername = (usernameLogin) => {
    setUsername(usernameLogin);
    setUsernameDefined(true);
  }

  const send = (event) => {
    if (event.keyCode === 13) {
      sendMsg(event.target.value,username)
      event.target.value = "";
    } 
  }

  useEffect(() => {
    console.log("entrou no useEffect")
    connect((msg) => {
      console.log("New Message")
      setChatHistory(chatHistory => ({
        messages : [...chatHistory.messages, msg]
      }))
    });
  },[])

  return (
      <div className="App">
        <Header/>
        { !isUsernameDefined ?
        <div className="LoginPage">
        <LoginPage sentUsername = {sentUsername}/>
        </div>
        : null }
        { isUsernameDefined ?  
          <div className="Chat">
          <ChatHistory chatHistory={chatHistory.messages} username = {username}/>
          <ChatInput send={send} />
          </div> 
        : null }
      </div>
    );
}
export default App;