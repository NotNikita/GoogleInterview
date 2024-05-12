import styled from 'styled-components';

interface ListContainerProps {
  list: number[];
  onClick: (el: number) => void;
  'data-testid': string;
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  gap: 20px;
  width: 300px;
  height: 300px;
  align-self: center;
  border: 2px solid black;
`;

export const ListContainer = ({list, onClick, ...props}: ListContainerProps) => {
  return (
    <Container data-testid={props['data-testid']}>
      {list.map(el => {
        const elementKey = `${props['data-testid']}-${el}`;
        return (
          <div key={elementKey}>
            <input data-testid={elementKey} type='checkbox' onClick={() => onClick(el)} />
            {el}
          </div>
        );
      })}
    </Container>
  );
};
