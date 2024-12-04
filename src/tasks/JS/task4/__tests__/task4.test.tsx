import {pipe, times, plus, substract, divide} from '../task4';

describe('task4', () => {
  it('test 1', () => {
    const calculation = pipe([times(2), times(3)]);
    expect(calculation(2)).toEqual(12);
  });

  it('test 2', () => {
    const calculation = pipe([times(2), times(3)]);
    expect(calculation(4)).toEqual(24);
  });

  it('test 3', () => {
    const calculation = pipe([times(2), plus(3), times(4)]);
    expect(calculation(2)).toEqual(28);
  });

  it('test 4', () => {
    const calculation = pipe([substract(20), divide(4)]);
    expect(calculation(100)).toEqual(20);
  });

  it('test 5', () => {
    const calculation = pipe([plus(5), times(3), divide(10)]);
    expect(calculation(5)).toEqual(3);
  });
});
