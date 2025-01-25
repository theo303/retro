const STICKY_WIDTH = 100;
const STICKY_HEIGHT = 100;

canvas = document.getElementById("canvas");
canvas.width = 0.98 * window.innerWidth;
canvas.height = 0.97 * window.innerHeight;
ctx = canvas.getContext("2d");

function drawSticky(x, y, content) {
  ctx.moveTo(x, y);
  ctx.beginPath();
  ctx.rect(x, y, STICKY_WIDTH, STICKY_HEIGHT);
  ctx.fillStyle = "yellow";
  ctx.fill();
  ctx.stroke();
  ctx.fillStyle = "black";
  ctx.textAlign = "center";
  ctx.font = "15px Arial";
  ctx.fillText(content, x + STICKY_WIDTH / 2, y + STICKY_HEIGHT / 2);
}
ws = new WebSocket("http://localhost:8080/connect");
ws.binaryType = "arraybuffer";

drawSticky(300, 300, "test");

protobuf.load("retro.proto", function (err, root) {
  if (err) console.log(err);

  var StateMessage = root.lookupType("retro.State");

  ws.onmessage = function (event) {
    var payload = new Uint8Array(event.data);
    var state = StateMessage.decode(payload);
    console.log(state);

    ctx.beginPath();
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    for (const stickyID in state.stickies) {
      const sticky = state.stickies[stickyID];
      console.log(sticky);
      console.log(`${stickyID}: ${state.stickies[stickyID]}`);
      drawSticky(sticky.X, sticky.Y, sticky.content);
    }
  };
});
