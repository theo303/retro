import { BuildEdges, EdgeSide } from "./edges";
import { LocalState, Selected } from "./state";
import { State, Action, Sticky } from "./retro";

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

var localState = new LocalState(ctx);
var userID: string;
var state: State;
var isMoving = false;
var isEditing = false;
var isResizing = false;

ws.onmessage = function (event) {
  if (typeof event.data === "string") {
    userID = event.data;
    return;
  }
  const payload = new Uint8Array(event.data);
  state = State.decode(payload);
  if (!localState.selected) {
    for (let sticky of state.stickies) {
      if (sticky.selectedBy === userID) {
        localState.selected = new Selected(sticky);
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

  let isOnEdge = localState.IsOnEdge(x, y);
  if (isOnEdge) {
    localState.selected = new Selected(isOnEdge.sticky);
    localState.selected.edge = isOnEdge.edgeSide;
    isResizing = true;
    return;
  }

  for (let s of localState.stickies) {
    if (ctx.isPointInPath(s.path, x, y)) {
      if (localState.selected && s.sticky.id == localState.selected.sticky.id) {
        isEditing = true;
      }
      const selectActionMessage = Action.create({
        select: { StickyID: s.sticky.id },
      });
      localState.selected = new Selected(s.sticky);
      localState.selected.offset = {
        x: x - s.sticky.X,
        y: y - s.sticky.Y,
      };
      isMoving = true;
      const bb = Action.encode(selectActionMessage).finish();
      ws.send(bb);
      return;
    }
  }
  localState.selected = undefined;

  const addActionMessage = Action.create({
    add: {
      X: x - STICKY_WIDTH / 2,
      Y: y - STICKY_HEIGHT / 2,
      Width: STICKY_WIDTH,
      Height: STICKY_HEIGHT,
    },
  });
  var bb = Action.encode(addActionMessage).finish();
  ws.send(bb);
});

canvas.addEventListener("mouseup", function (_) {
  isMoving = false;
  isResizing = false;
});

canvas.addEventListener("mousemove", function (e) {
  const x = e.clientX;
  const y = e.clientY;

  let isOnEdge = localState.IsOnEdge(x, y);
  if (isOnEdge) {
    switch (isOnEdge.edgeSide) {
      case EdgeSide.N:
        document.body.style.cursor = "n-resize";
        break;
      case EdgeSide.E:
        document.body.style.cursor = "e-resize";
        break;
      case EdgeSide.S:
        document.body.style.cursor = "s-resize";
        break;
      case EdgeSide.W:
        document.body.style.cursor = "w-resize";
        break;
    }
  } else {
    document.body.style.cursor = "default";
  }

  if (!localState.selected) {
    return;
  }

  if (isMoving) {
    const moveActionMessage = Action.create({
      move: {
        StickyID: localState.selected.sticky.id,
        X: x - localState.selected.offset!.x,
        Y: y - localState.selected.offset!.y,
      },
    });
    const bb = Action.encode(moveActionMessage).finish();
    ws.send(bb);
    return;
  }

  if (isResizing) {
    if (!localState.selected.edge) {
      return;
    }
    let sticky = localState.selected.sticky;
    let newX = sticky.X;
    let newY = sticky.Y;
    let newHeight = sticky.Height;
    let newWidth = sticky.Width;
    let delta: number;
    switch (localState.selected.edge) {
      case EdgeSide.N:
        delta = y - sticky.Y;
        newY = sticky.Y + delta;
        newHeight = sticky.Height - delta;
        break;
      case EdgeSide.E:
        delta = x - sticky.X - sticky.Width;
        newWidth = sticky.Width + delta;
        break;
      case EdgeSide.S:
        delta = y - sticky.Y - sticky.Height;
        newHeight = sticky.Height + delta;
        break;
      case EdgeSide.W:
        delta = x - sticky.X;
        newX = sticky.X + delta;
        newWidth = sticky.Width - delta;
        break;
    }
    const resizeActionMessage = Action.create({
      resize: {
        StickyID: localState.selected.sticky.id,
        X: newX,
        Y: newY,
        Height: newHeight,
        Width: newWidth,
      },
    });
    const bb = Action.encode(resizeActionMessage).finish();
    ws.send(bb);
    return;
  }
});

canvas.addEventListener("keypress", function (e) {
  if (!localState.selected) {
    return;
  }

  switch (e.key) {
    case "Enter":
      isEditing = false;
      localState.selected = undefined;
      return;
    case "Delete":
      const deleteStickyMessage = Action.create({
        delete: {
          StickyID: localState.selected.sticky.id,
        },
      });
      localState.selected = undefined;
      const bb = Action.encode(deleteStickyMessage).finish();
      ws.send(bb);
      return;
  }

  if (!isEditing) {
    return;
  }

  localState.selected.sticky.content += e.key;
  const updateContentMessage = Action.create({
    edit: {
      StickyID: localState.selected.sticky.id,
      content: localState.selected.sticky.content,
    },
  });
  const bb = Action.encode(updateContentMessage).finish();
  ws.send(bb);
});

canvas.addEventListener("keydown", function (e) {
  if (!isEditing || !localState.selected) {
    return;
  }

  if (e.key === "Backspace") {
    const content = localState.selected.sticky.content;
    if (content.length === 0) {
      return;
    }

    localState.selected.sticky.content = content.substring(
      0,
      content.length - 1,
    );
    const updateContentMessage = Action.create({
      edit: {
        StickyID: localState.selected.sticky.id,
        content: localState.selected.sticky.content,
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

  localState.stickies = [];
  for (let s of state.stickies) {
    localState.stickies.push({
      sticky: s,
      path: drawSticky(s),
      edges: BuildEdges(s),
    });
  }
  localState.stickies.reverse();

  requestAnimationFrame(update);
}

function drawSticky(s: Sticky): Path2D {
  ctx.moveTo(s.X, s.Y);
  ctx.beginPath();
  const path = new Path2D();
  path.roundRect(s.X, s.Y, s.Width, s.Height, 2);
  ctx.fillStyle = "yellow";
  if (s.selectedBy === userID) {
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
  ctx.fillText(s.content, s.X + s.Width / 2, s.Y + s.Height / 2);
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
