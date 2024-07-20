import {promise1, promise2, promise3, promiseRace} from '../task5';

describe('task5', () => {
  it('test 1', async () => {
    const result = await promiseRace([promise1, promise2]);
    expect(result).toEqual('promise-2 resolved');
  });

  it('test 2', async () => {
    const result = await promiseRace([promise1, promise2, promise3]);
    expect(result).toEqual('promise-2 resolved');
  });

  it('test 3', async () => {
    const result = await promiseRace([promise2, promise3]);
    expect(result).toEqual('promise-2 resolved');
  });

  it('test 4', async () => {
    const result = await promiseRace([promise1, promise3]);
    expect(result).toEqual('promise-3 resolved');
  });
});
