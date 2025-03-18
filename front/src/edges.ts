import { Sticky } from "./retro";

const EDGES_THICKNESS = 5;

export enum EdgeSide {
  N = "north",
  E = "east",
  S = "south",
  W = "west",
}

export interface Edges {
  N: Path2D;
  E: Path2D;
  S: Path2D;
  W: Path2D;
}

export function BuildEdges(sticky: Sticky): Edges {
  let n = new Path2D();
  n.rect(sticky.X, sticky.Y + EDGES_THICKNESS, sticky.Width, EDGES_THICKNESS);

  let e = new Path2D();
  e.rect(
    sticky.X + sticky.Width + EDGES_THICKNESS,
    sticky.Y,
    EDGES_THICKNESS,
    sticky.Height,
  );

  let s = new Path2D();
  s.rect(
    sticky.X,
    sticky.Y + sticky.Height + EDGES_THICKNESS,
    sticky.Width,
    EDGES_THICKNESS,
  );

  let w = new Path2D();
  w.rect(sticky.X + EDGES_THICKNESS, sticky.Y, EDGES_THICKNESS, sticky.Height);

  return { N: n, E: e, S: s, W: w };
}
