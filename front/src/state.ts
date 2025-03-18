import { Edges, EdgeSide } from "./edges";
import { Sticky } from "./retro";

export interface DrawnSticky {
  sticky: Sticky;
  path: Path2D;
  edges: Edges;
}

export class Selected {
  sticky: Sticky;
  edge: EdgeSide | undefined = undefined;
  offset: { x: number; y: number } | undefined = undefined;

  constructor(sticky: Sticky) {
    this.sticky = sticky;
  }
}

export class LocalState {
  ctx: CanvasDrawPath;
  stickies: DrawnSticky[] = [];

  selected: Selected | undefined;

  constructor(ctx: CanvasDrawPath) {
    this.ctx = ctx;
  }

  IsOnEdge(
    this: LocalState,
    x: number,
    y: number,
  ): { sticky: Sticky; edgeSide: EdgeSide } | undefined {
    for (let s of this.stickies) {
      if (this.ctx.isPointInPath(s.edges.N, x, y)) {
        return { sticky: s.sticky, edgeSide: EdgeSide.N };
      }
      if (this.ctx.isPointInPath(s.edges.E, x, y)) {
        return { sticky: s.sticky, edgeSide: EdgeSide.E };
      }
      if (this.ctx.isPointInPath(s.edges.S, x, y)) {
        return { sticky: s.sticky, edgeSide: EdgeSide.S };
      }
      if (this.ctx.isPointInPath(s.edges.W, x, y)) {
        return { sticky: s.sticky, edgeSide: EdgeSide.W };
      }
    }
    return;
  }
}
