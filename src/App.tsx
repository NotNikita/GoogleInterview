import {useState} from 'react';
import './App.css';
import {Task1} from './tasks/task1/Task1';
import {Pagination} from '@gravity-ui/uikit';
import styled from 'styled-components';
import {Task2} from './tasks/task2/Task2';
import {Task3} from './tasks/task3/Task3';

const Layout = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 30px;
`;

const TASKS_MAP = new Map<number, JSX.Element>([
  [1, <Task1 />],
  [2, <Task2 />],
  [3, <Task3 />],
]);

function App() {
  const [displayedTask, setDisplayedTask] = useState(2);

  const changeTask = (page: number) => {
    setDisplayedTask(page);
  };

  const TaskComponent = TASKS_MAP.get(displayedTask);

  return (
    <Layout>
      <h1>Task {displayedTask}</h1>
      {TaskComponent}
      <Pagination page={1} pageSize={1} total={TASKS_MAP.size} onUpdate={changeTask} />
    </Layout>
  );
}

export default App;
