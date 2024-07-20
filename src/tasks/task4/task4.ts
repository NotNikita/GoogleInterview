// Pipe function for vanilla js

export const times = (y: number) => (x: number) => x * y;
export const plus = (y: number) => (x: number) => x + y;
export const substract = (y: number) => (x: number) => x - y;
export const divide = (y: number) => (x: number) => x / y;

export function pipe(funcArray: any[]) {
  return (firstValue: number) => {
    return funcArray.reduce((result, foo) => {
      return foo(result);
    }, firstValue);
  };
}
