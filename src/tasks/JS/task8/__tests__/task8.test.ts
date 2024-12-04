import {getRoute, getRouteFaster} from '../task8';

describe('task8', () => {
  it('should get correct route 1', () => {
    const start = performance.now();
    const result = getRoute();
    const end = performance.now();
    const duration = end - start;
    console.log(`getRoute took ${duration}ms`); // 0.215s
    expect(result).toEqual(['Dubai', 'Astana', 'Bali', 'Dublin']);
  });

  it('should get correct route faster', () => {
    const start = performance.now();
    const result = getRouteFaster();
    const end = performance.now();
    const duration = end - start;
    console.log(`getRouteFaster took ${duration}ms`); // 0.094ms
    expect(result).toEqual(['Dubai', 'Astana', 'Bali', 'Dublin']);
  });
});
