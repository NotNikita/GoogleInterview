import {useState} from 'react';
import {GridCell, Side, User} from './types';
import {getEmptyGrid} from './utils';
import {GridContainer, Cell} from './styles';

const WINNING_SCENARIOS = [
  // horizontal
  [0, 1, 2],
  [3, 4, 5],
  [6, 7, 8],
  // vertical
  [0, 3, 6],
  [1, 4, 7],
  [2, 5, 8],
  // diagonal
  [0, 4, 8],
  [2, 4, 6],
];

export const Task2 = () => {
  const [isFirstPlayer, setIsFirstPlayer] = useState(true);
  const [disabled, setDisabled] = useState(false);
  const [winningScenario, setWinningScenario] = useState<number[]>([]);
  const players: User[] = [
    {
      name: 'Player 1',
      side: Side.Nought,
    },
    {
      name: 'Player 2',
      side: Side.Cross,
    },
  ];
  const [grid, setGrid] = useState(getEmptyGrid());

  const checkIfGameFinished = (cells: GridCell[]): boolean => {
    const startTime = performance.now();
    let isFinished = false;

    for (const side of Object.keys(Side)) {
      for (const scenario of WINNING_SCENARIOS) {
        const elements = scenario.map(cellId => cells[cellId]);
        if (elements.every(element => element.status === side)) {
          isFinished = true;
          setWinningScenario(scenario);
        }
      }
    }

    console.log(`checkIfGameFinished took ${performance.now() - startTime} milliseconds.`);
    return isFinished;
  };

  const handleCellClick = (id: number) => {
    // TODO: can be done better
    const player = players[isFirstPlayer ? 0 : 1];

    const updatedCells = grid.cells.map(cell => {
      if (cell.cellId !== id) return cell;

      return {
        cellId: id,
        status: player.side,
      } as GridCell;
    });
    const finished = checkIfGameFinished(updatedCells);
    setGrid({
      cells: updatedCells,
      gameFinished: finished,
      winner: finished ? player : undefined,
    });
    setIsFirstPlayer(prev => !prev);
    setDisabled(finished);
  };

  return (
    <>
      {grid.winner ?
        <h2>🎉🎉Congratulations to {grid.winner.name} for Winning!🎉🎉🎉</h2>
      : <h2>Its turn of {players[isFirstPlayer ? 0 : 1].name}</h2>}
      <GridContainer>
        {grid.cells.map(({cellId, status}: GridCell) => (
          <Cell
            key={cellId}
            side={status}
            onClick={() => {
              if (status || disabled) return;
              handleCellClick(cellId);
            }}
            highlight={winningScenario.includes(cellId)}
          />
        ))}
      </GridContainer>
    </>
  );
};
