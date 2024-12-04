import {ATM} from '../task7';

describe('task7', () => {
  let mockAtm = new ATM();

  beforeEach(() => {
    // reseting ATM state
    mockAtm = new ATM();
  });

  it('should work for depositing 1', () => {
    const testResult = mockAtm.deposit([5000, 1000, 5000, 500, 100, 50, 50]);
    expect(testResult).toEqual('Contributed: 11700');
    expect(mockAtm.total).toEqual(20500);
  });

  it('should work for depositing 2', () => {
    const testResult = mockAtm.deposit([500, 10, 5]);
    expect(testResult).toEqual('Contributed 500, Collect unrecognized bills 10,5');
    expect(mockAtm.total).toEqual(9300);
  });

  it('should work for depositing empty array edge case', () => {
    const testResult = mockAtm.deposit([]);
    expect(testResult).toEqual('Error: Put money into the bill acceptor');
    expect(mockAtm.total).toEqual(8800);
  });

  it('should return correct accepted bills', () => {
    const testResult = mockAtm.acceptedBills;
    expect(testResult).toEqual(['50', '100', '500', '1000', '2000', '5000']);
  });

  it('should handle withdrawing 3500 with insufficient funds', () => {
    mockAtm.clearVault();
    const testResult = mockAtm.withdraw(3500);
    expect(testResult).toEqual("Error: I can't issue the required amount, there are not enough funds in the ATM");
  });

  it('should be able to withdraw 2100 correctly', () => {
    mockAtm.deposit([1000, 500, 500, 100]);
    const testResult = mockAtm.withdraw(2100);
    expect(testResult).toEqual([1000, 500, 500, 100]);
  });

  it('should handle withdrawing 0 or less error', () => {
    const testResult = mockAtm.withdraw(0);
    expect(testResult).toEqual('Error: Please enter the correct amount');
  });

  it('should handle withdrawing, when no available bills error', () => {
    const testResult = mockAtm.withdraw(9999);
    expect(testResult).toEqual("Error: I can't issue the required amount, there are not enough funds in the ATM");
  });

  it('should handle total amount', () => {
    const testResult = mockAtm.total;
    expect(testResult).toEqual(8800);
  });

  it('should handle collect for incasation services', () => {
    const testResult = mockAtm.acceptedBills;
    expect(testResult).toEqual(['50', '100', '500', '1000', '2000', '5000']);
  });
});
