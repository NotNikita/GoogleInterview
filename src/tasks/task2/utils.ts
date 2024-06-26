import {Grid, GridCell} from './types';

export const getEmptyGrid = (): Grid => {
  const cells = Array.from(
    {length: 9},
    (_, index) =>
      ({
        cellId: index,
      }) as GridCell
  );
  return {
    cells,
    gameFinished: false,
    winner: undefined,
  };
};
