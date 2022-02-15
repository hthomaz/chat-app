// api/index.js
if(window.location.host === "localhost:3000"){
  var socket = new WebSocket("ws://" + "localhost:4444" + "/websocket");
} else {
  var socket = new WebSocket("wss://" + window.location.host + "/websocket");
}


let connect = cb => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log("on message");
    cb(msg)
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg,username) => {
  //console.log("sending msg: ", msg);
  
  socket.send(
      JSON.stringify({
          username: username,
          text: msg,
          color : "black",
          destination: "all"
      })
  )
};

export { connect, sendMsg };