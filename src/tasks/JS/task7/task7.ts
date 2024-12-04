// ========================================== Задача 1
// You need to implement an ATM operations: withdrawal and deposit

export class ATM {
  private vault: {[key: number]: number};
  private bills: string[];

  constructor() {
    this.vault = {
      5000: 1,
      2000: 0,
      1000: 0,
      500: 6,
      100: 5,
      50: 6,
    };
    this.bills = Object.keys(this.vault);
  }

  clearVault() {
    this.vault = {
      5000: 0,
      2000: 0,
      1000: 0,
      500: 0,
      100: 0,
      50: 0,
    };
  }

  deposit(bills: number[]) {
    // code here
    if (!bills || !bills.length) {
      return 'Error: Put money into the bill acceptor';
    }
    const filteredBills = bills.filter(bill => this.bills.includes(bill.toString()));
    const unacceptedBills = bills.filter(bill => !this.bills.includes(bill.toString()));
    const billsSum = filteredBills.reduce((acc, bill) => {
      this.vault[bill] = this.vault[bill] === 0 ? 1 : this.vault[bill] + 1;
      return acc + bill;
    }, 0);
    return unacceptedBills.length ?
        `Contributed ${billsSum}, Collect unrecognized bills ${unacceptedBills}`
      : `Contributed: ${billsSum}`;
  }

  withdraw(amount: number) {
    // code here
    if (amount <= 0) {
      return 'Error: Please enter the correct amount';
    }
    if (this.total < amount) {
      return "Error: I can't issue the required amount, there are not enough funds in the ATM";
    }
    const bills = this.bills.map(Number).sort((a, b) => b - a);
    const result = [];
    let remainingAmount = amount;

    for (const bill of bills) {
      while (remainingAmount >= bill && this.vault[bill] > 0) {
        remainingAmount -= bill;
        this.vault[bill] -= 1;
        result.push(bill);
      }
    }

    if (remainingAmount > 0) {
      // If we cannot dispense the exact amount, roll back the changes
      result.forEach(bill => {
        this.vault[bill] += 1;
      });
      return 'Error: Cannot dispense the exact amount with available bills';
    }
    return result;
  }

  // returns an array of bills, that are accepted for with./dep.
  get acceptedBills() {
    // code here
    return this.bills;
  }

  // returns current balance
  get total() {
    // code here
    const billsSum = Object.entries(this.vault).reduce((acc, [key, value]) => acc + Number(key) * value, 0);
    return billsSum;
  }

  // return the data cassette as an object during cash collection
  get collect() {
    // code here
    return `Cassette ${Object.entries(this.vault)
      .map(([key, value]) => `'${key}' : ${value}`)
      .join(', ')}`;
  }
}
