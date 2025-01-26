const STICKY_WIDTH = 100;
const STICKY_HEIGHT = 100;

canvas = document.getElementById("canvas");
canvas.width = 0.98 * window.innerWidth;
canvas.height = 0.97 * window.innerHeight;
ctx = canvas.getContext("2d");

ws = new WebSocket("http://localhost:8080/connect");
ws.binaryType = "arraybuffer";

var stickiesPaths = new Map();
var userID;

protobuf.load("retro.proto", function (err, root) {
  if (err) console.log(err);

  const State = root.lookupType("retro.State");
  const Action = root.lookupType("retro.Action");

  var firstMessage = true;
  ws.onmessage = function (event) {
    if (firstMessage) {
      firstMessage = false;
      userID = event.data;
    }
    const payload = new Uint8Array(event.data);
    const state = State.decode(payload);

    clear();

    for (const stickyID in state.stickies) {
      const sticky = state.stickies[stickyID];
      const path = drawSticky(
        sticky.X,
        sticky.Y,
        sticky.content,
        sticky.selectedBy,
      );
      stickiesPaths.set(stickyID, path);
    }
  };

  canvas.addEventListener("mousedown", function (e) {
    const rect = canvas.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    console.log("x: " + x + " y: " + y);

    for (let [stickyID, path] of stickiesPaths) {
      if (ctx.isPointInPath(path, x, y)) {
        const selectActionMessage = Action.create({
          select: { StickyID: stickyID },
        });
        const bb = Action.encode(selectActionMessage).finish();
        ws.send(bb);
        return;
      }
    }

    const addActionMessage = Action.create({
      add: { X: x, Y: y },
    });
    var bb = Action.encode(addActionMessage).finish();
    ws.send(bb);
  });
});

function drawSticky(x, y, content, selectedBy) {
  ctx.moveTo(x, y);
  ctx.beginPath();
  const path = new Path2D();
  path.rect(x, y, STICKY_WIDTH, STICKY_HEIGHT);
  ctx.fillStyle = "yellow";
  if (selectedBy === userID) {
    ctx.strokeStyle = "blue";
    ctx.lineWidth = 5;
  } else {
    ctx.strokeStyle = "black";
    ctx.lineWidth = 2;
  }
  ctx.fill(path);
  ctx.stroke(path);
  ctx.fillStyle = "black";
  ctx.textAlign = "center";
  ctx.font = "15px Arial";
  ctx.fillText(content, x + STICKY_WIDTH / 2, y + STICKY_HEIGHT / 2);
  return path;
}

function clear() {
  ctx.beginPath();
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  ctx.fillStyle = "black";
  ctx.textAlign = "left";
  ctx.textBaseline = "top";
  ctx.font = "15px Arial";
  ctx.fillText(userID, 0, 0);
}
