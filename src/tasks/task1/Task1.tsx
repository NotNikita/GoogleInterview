import {useState} from 'react';

export const Task1 = () => {
  const [list1, setlist1] = useState<number[]>([1, 2, 3, 4]);
  const [list2, setlist2] = useState<number[]>([]);
  const [toMoveLeft, settoMoveLeft] = useState<number[]>([]);
  const [toMoveRight, settoMoveRight] = useState<number[]>([]);

  function handleClick(num: number, setQueue: any) {
    setQueue((prev: number[]) => (prev.includes(num) ? prev.filter(p => p !== num) : [...prev, num]));
  }

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'row',
        gap: 20,
        width: '300px',
        height: '400px',
      }}
    >
      <div style={{padding: '20px 0', border: '2px solid black', width: 70, alignSelf: 'center'}}>
        {list1.map(el => (
          <div key={`list1-${el}`}>
            <input type='checkbox' onClick={() => handleClick(el, settoMoveLeft)} />
            {el}
          </div>
        ))}
      </div>

      <div style={{display: 'flex', flexDirection: 'column', height: '100%', gap: 15, justifyContent: 'center'}}>
        <button
          style={{
            width: 60,
            height: 40,
            background: 'peach',
            border: '1px solid black',
          }}
          onClick={() => {
            setlist1(list1.filter(el => !toMoveLeft.includes(el)));
            setlist2(list2.concat(toMoveLeft));
            settoMoveLeft([]);
          }}
        >{`>`}</button>
        <button
          style={{
            width: 60,
            height: 40,
            background: 'peach',
            border: '1px solid black',
          }}
          onClick={() => {
            setlist2(list2.filter(el => !toMoveRight.includes(el)));
            setlist1(list1.concat(toMoveRight));
            settoMoveRight([]);
          }}
        >{`<`}</button>
      </div>

      <div
        style={{
          padding: '20px 0',
          border: '2px solid black',
          width: 70,
          alignSelf: 'center',
        }}
      >
        {list2.map(el => (
          <div key={`list2-${el}`}>
            <input type='checkbox' onClick={() => handleClick(el, settoMoveRight)} />
            {el}
          </div>
        ))}
      </div>
    </div>
  );
};
