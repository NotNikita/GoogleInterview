// Pipe function for vanilla js

export const promise1 = new Promise(res => {
  setTimeout(() => res('promise-1 resolved'), 300);
});

export const promise2 = new Promise(res => {
  setTimeout(() => res('promise-2 resolved'), 200);
});

export const promise3 = new Promise(res => {
  setTimeout(() => res('promise-3 resolved'), 100);
});

export function promiseRace(promises: any[]) {
  // return Promise.race(promises);
  return new Promise((resolve, reject) => {
    promises.forEach(promise => {
      promise
        .then((promiseResult: any) => {
          resolve(promiseResult);
        })
        .catch(reject);
    });
  });
}
