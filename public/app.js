window.addEventListener("DOMContentLoaded", (_) => {
  let websocket
  let activeUsername
  if (window.location.host === "localhost:4444") {
    websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  } else {
    websocket = new WebSocket("wss://" + window.location.host + "/websocket");
  }
  
  let room = document.getElementById("chat-text");

  websocket.addEventListener("message", function (e) {
    let data = JSON.parse(e.data);
    // creating html element
    console.log(data)
    if (data.destination === "all" || data.destination === activeUsername.value || data.username === activeUsername.value) {
      console.log("entrou")
      let p = document.createElement("p");
      p.innerHTML = `<strong>${data.username}</strong>: ${data.text}`;
      p.style.color = data.color
      room.append(p);
      room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
    }
  });

  let input_form = document.getElementById("login-form");
  input_form.addEventListener("submit", function(event) {
    event.preventDefault();
    activeUsername = document.getElementById("login-username");
    document.getElementById("input-username").value = activeUsername.value

  });

  let form = document.getElementById("input-form");
  form.addEventListener("submit", function (event) {
    event.preventDefault();
    //let username = document.getElementById("input-username");
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
  });
});

$(document).ready(function() {             
$('#loginModal').modal('show');
  $(function () {
    $('[data-toggle="tooltip"]').tooltip()
  })
});
