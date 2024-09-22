import {getRoute} from '../task8';

describe('task8', () => {
  it('should get correct route 1', () => {
    const result = getRoute();
    expect(result).toEqual([
      {from: 'Dubai', to: 'Astana'},
      {from: 'Astana', to: 'Bali'},
      {from: 'Bali', to: 'Dublin'},
    ]);
  });
});
