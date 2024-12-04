import styled, {css} from 'styled-components';
import {ifProp, switchProp} from 'styled-tools';
import {Side} from './types';

export const Cell = styled.div<{side?: Side; highlight: boolean}>`
  width: 90px;
  height: 90px;
  border: 1px solid black;
  position: relative;
  background-color: ${ifProp('highlight', '#E2F0D9', 'default')};

  ${switchProp('side', {
    [Side.Cross]: css`
      &::before,
      &::after {
        content: '';
        position: absolute;
        left: 50%;
        top: 50%;
        width: 60px; /* Length of the cross lines */
        height: 10px; /* Thickness of the cross lines */
        background-color: red;
      }
    `,
    [Side.Nought]: css`
      &::after {
        content: '';
        position: absolute;
        left: 50%;
        top: 50%;
        background-color: transparent;
        border: 10px solid blue;
        border-radius: 50%;
        height: 40px; /* size of circle */
        width: 40px; /* size of circle */
      }
    `,
  })}

  &::before {
    transform: translate(-50%, -50%) rotate(45deg);
  }

  &::after {
    transform: translate(-50%, -50%) rotate(-45deg);
  }
`;

export const GridContainer = styled.div`
  display: grid;
  grid-template-columns: repeat(3, 1fr); // 3 columns
  grid-template-rows: repeat(3, 1fr); // 3 rows
  gap: 10px;
`;
