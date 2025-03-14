import { State, Action } from "./retro";

const STICKY_WIDTH = 100;
const STICKY_HEIGHT = 100;

var canvas = document.getElementById("canvas")! as HTMLCanvasElement;
canvas.width = 0.98 * window.innerWidth;
canvas.height = 0.97 * window.innerHeight;
var ctx = canvas.getContext("2d")!;

var serverAddress = `${window.location.protocol === "https:" ? "wss:" : "ws:"}//${window.location.host}`;
if (import.meta.env.DEV) {
  serverAddress = "ws://localhost:8080";
}

var ws = new WebSocket(`${serverAddress}/connect?name=browser`);
ws.binaryType = "arraybuffer";

var stickiesPaths: { id: string; path: Path2D }[] = [];
var userID: string;
var state: State;
var isMoving = false;
var isEditing = false;
var selected: { id: string; offset?: { x: number; y: number } } | undefined;

ws.onmessage = function (event) {
  if (typeof event.data === "string") {
    userID = event.data;
    return;
  }
  const payload = new Uint8Array(event.data);
  state = State.decode(payload);
  if (!selected) {
    for (let id in state.stickies) {
      if (state.stickies[id].selectedBy === userID) {
        selected = { id: id };
      }
    }
  }
};

canvas.addEventListener("mousedown", function (e) {
  if (e.button !== 0) {
    return;
  }

  const x = e.clientX;
  const y = e.clientY;

  isEditing = false;
  for (let s of stickiesPaths) {
    if (ctx.isPointInPath(s.path, x, y)) {
      if (selected && s.id == selected.id) {
        isEditing = true;
        return;
      }
      const selectActionMessage = Action.create({
        select: { StickyID: s.id },
      });
      selected = {
        id: s.id,
        offset: {
          x: x - state.stickies[s.id].X,
          y: y - state.stickies[s.id].Y,
        },
      };
      isMoving = true;
      const bb = Action.encode(selectActionMessage).finish();
      ws.send(bb);
      return;
    }
  }
  selected = undefined;

  const addActionMessage = Action.create({
    add: {
      X: x - STICKY_WIDTH / 2,
      Y: y - STICKY_HEIGHT / 2,
    },
  });
  var bb = Action.encode(addActionMessage).finish();
  ws.send(bb);
});

canvas.addEventListener("mouseup", function (_) {
  isMoving = false;
});

canvas.addEventListener("mousemove", function (e) {
  if (isMoving === false || selected === undefined) {
    return;
  }

  const moveActionMessage = Action.create({
    move: {
      StickyID: selected.id,
      X: e.clientX - selected.offset!.x,
      Y: e.clientY - selected.offset!.y,
    },
  });
  const bb = Action.encode(moveActionMessage).finish();
  ws.send(bb);
  return;
});

canvas.addEventListener("keypress", function (e) {
  if (!isEditing || selected === undefined) {
    return;
  }

  const updateContentMessage = Action.create({
    edit: {
      StickyID: selected.id,
      content: state.stickies[selected.id].content + e.key,
    },
  });
  const bb = Action.encode(updateContentMessage).finish();
  ws.send(bb);
});

canvas.addEventListener("keydown", function (e) {
  if (!isEditing || selected === undefined) {
    return;
  }

  if (e.key === "Backspace") {
    const content = state.stickies[selected.id].content;
    if (content.length === 0) {
      return;
    }

    const updateContentMessage = Action.create({
      edit: {
        StickyID: selected.id,
        content: content.substring(0, content.length - 1),
      },
    });
    const bb = Action.encode(updateContentMessage).finish();
    ws.send(bb);
  }
});

var lastTime: DOMHighResTimeStamp;

function update(time: DOMHighResTimeStamp) {
  if (!state) {
    requestAnimationFrame(update);
    return;
  }
  clear(time - lastTime);
  lastTime = time;

  var stickiesInOrder = [];
  for (const stickyID in state.stickies) {
    stickiesInOrder.push({ id: stickyID, sticky: state.stickies[stickyID] });
  }

  stickiesInOrder.sort((a, b) => a.sticky.height - b.sticky.height);

  stickiesPaths = [];
  for (const s of stickiesInOrder) {
    const path = drawSticky(
      s.sticky.X,
      s.sticky.Y,
      s.sticky.content,
      s.sticky.selectedBy,
    );
    stickiesPaths.push({ id: s.id, path: path });
  }
  stickiesPaths.reverse();

  requestAnimationFrame(update);
}

function drawSticky(
  x: number,
  y: number,
  content: string,
  selectedBy: string | undefined,
) {
  ctx.moveTo(x, y);
  ctx.beginPath();
  const path = new Path2D();
  path.rect(x, y, STICKY_WIDTH, STICKY_HEIGHT);
  ctx.fillStyle = "yellow";
  if (selectedBy === userID) {
    if (isEditing) {
      ctx.strokeStyle = "red";
    } else {
      ctx.strokeStyle = "blue";
    }
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

function clear(deltaTime: number) {
  ctx.beginPath();
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  ctx.fillStyle = "black";
  ctx.textAlign = "left";
  ctx.textBaseline = "top";
  ctx.font = "15px Arial";
  ctx.fillText(userID, 0, 0);
  ctx.fillText(Math.round((1 / deltaTime) * 1000).toString(), 0, 20);
}

requestAnimationFrame(update);
