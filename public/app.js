const sendLogin = document.querySelector('#send-login')
const sendMessage = document.querySelector('#send-message')
let activeUsername
let websocket


sendLogin.onclick = () => {
  activeUsername = document.getElementById("login-username");
  document.getElementById("input-username").value = activeUsername.value
  createWebsocket()
}

sendMessage.onclick = () => {
  let text = document.getElementById("input-text");
    websocket.send(
      JSON.stringify({
        username: activeUsername.value,
        text: text.value,
        color : "black",
        destination: "all"
      })
    );
  text.value = "";
}

function createWebsocket() {
  if (window.location.host === "localhost:4444") {
    websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  } else {
    websocket = new WebSocket("wss://" + window.location.host + "/websocket");
  }
}

if (websocket !== undefined) {
  websocket.onmessage = function(msg){
    let data = JSON.parse(msg.data)
    if (checkMsgDestination(data)) {
      insertMessage(data)
    }
  }
}

function insertMessage (messageObj) {
  let p = document.createElement("p");
  p.innerHTML = `<strong>${data.username}</strong>: ${data.text}`;
  p.style.color = data.color
  room.append(p);
  room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
}

function checkMsgDestination (messageObj) {
  if (data.destination === "all" || data.destination === activeUsername.value || data.username === activeUsername.value) {
    return true
  }
  return false
}

$(document).ready(function() {             
$('#loginModal').modal('show');
  $(function () {
    $('[data-toggle="tooltip"]').tooltip()
  })
});
