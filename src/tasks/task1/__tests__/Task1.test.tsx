import {fireEvent, render} from '@testing-library/react';
import {Task1} from '../Task1';

describe('Google task 1', () => {
  const renderConfirmation = (firstList = [1, 2, 3, 4], secondList = [5, 6]) =>
    render(<Task1 firstList={firstList} secondList={secondList} />);

  it('should move elements to right', () => {
    const {getByTestId} = renderConfirmation();

    expect(getByTestId('left-list').children.length).toBe(4);
    fireEvent.click(getByTestId('left-list-1'));
    const button = getByTestId('btn-move-right');
    fireEvent.click(button);

    expect(getByTestId('left-list').children.length).toBe(3);
    expect(getByTestId('right-list').children.length).toBe(3);
  });

  it('should move elements to left', () => {
    const {getByTestId} = renderConfirmation();

    expect(getByTestId('right-list').children.length).toBe(2);
    fireEvent.click(getByTestId('right-list-5'));
    const button = getByTestId('btn-move-left');
    fireEvent.click(button);

    expect(getByTestId('left-list').children.length).toBe(5);
    expect(getByTestId('right-list').children.length).toBe(1);
  });

  it('should not move selected left elements after clicking move-left', () => {
    const {getByTestId} = renderConfirmation();

    expect(getByTestId('left-list').children.length).toBe(4);
    fireEvent.click(getByTestId('left-list-1'));
    const button = getByTestId('btn-move-left');
    fireEvent.click(button);

    expect(getByTestId('left-list').children.length).toBe(4);
    expect(getByTestId('right-list').children.length).toBe(2);
  });

  it('should not move selected right elements after clicking move-right', () => {
    const {getByTestId} = renderConfirmation();

    expect(getByTestId('right-list').children.length).toBe(2);
    fireEvent.click(getByTestId('right-list-5'));
    const button = getByTestId('btn-move-right');
    fireEvent.click(button);

    expect(getByTestId('left-list').children.length).toBe(4);
    expect(getByTestId('right-list').children.length).toBe(2);
  });
});
