import { hello } from './hello-world';

describe('Hello World', () => {
  test('says hello', () => {
    expect(hello()).toEqual('Hello, World!');
  });
  let b = 3
  console.log("boom")
  switch(b) {
    case 1 || 2:
    console.log("bang")
    break;
  }
});
