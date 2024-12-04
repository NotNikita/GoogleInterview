export enum Side {
  Nought = 'Nought',
  Cross = 'Cross',
}
export interface GridCell {
  cellId: number;
  status: Side | undefined;
}
export interface User {
  name: string;
  side: Side;
}

// History of games can be stored in the file and extracted
export interface Grid {
  cells: GridCell[];
  gameFinished: boolean;
  winner: User | undefined;
}
