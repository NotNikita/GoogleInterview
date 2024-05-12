import {useState} from 'react';
import {ListContainer} from './components/List';
import {Button} from './components/Button';
import {Wrapper} from './styles';

interface Task1Props {
  firstList: number[];
  secondList: number[];
}
export const Task1 = ({firstList = [1, 2, 3, 4], secondList = []}: Task1Props) => {
  const [list1, setlist1] = useState<number[]>(firstList);
  const [list2, setlist2] = useState<number[]>(secondList);
  const [toMoveLeft, settoMoveLeft] = useState<number[]>([]);
  const [toMoveRight, settoMoveRight] = useState<number[]>([]);

  function handleClick(num: number, setQueue: any) {
    setQueue((prev: number[]) => (prev.includes(num) ? prev.filter(p => p !== num) : [...prev, num]));
  }

  function handleMoveRight() {
    setlist1(list1.filter(el => !toMoveRight.includes(el)));
    setlist2(list2.concat(toMoveRight));
    settoMoveRight([]);
  }

  function handleMoveLeft() {
    setlist2(list2.filter(el => !toMoveLeft.includes(el)));
    setlist1(list1.concat(toMoveLeft));
    settoMoveLeft([]);
  }

  return (
    <>
      <h1>Task 1</h1>
      <Wrapper>
        <ListContainer data-testid='left-list' list={list1} onClick={element => handleClick(element, settoMoveRight)} />

        <div>
          <Button data-testid='btn-move-left' label={'<'} onClick={handleMoveLeft} />
          <Button data-testid='btn-move-right' label={'>'} onClick={handleMoveRight} />
        </div>

        <ListContainer data-testid='right-list' list={list2} onClick={element => handleClick(element, settoMoveLeft)} />
      </Wrapper>
    </>
  );
};
