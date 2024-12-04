import styled from 'styled-components';

interface ButtonContainerProps {
  label: string;
  onClick: () => void;
  'data-testid': string;
}

const SButton = styled.button`
  width: 60px;
  height: 40px
  background: peach;
  border: 1px solid black;
  margin: 10px;
`;

export const Button = ({label, onClick, ...props}: ButtonContainerProps) => {
  return (
    <SButton data-testid={props['data-testid']} onClick={onClick}>
      {label}
    </SButton>
  );
};
