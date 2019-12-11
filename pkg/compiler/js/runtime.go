package js

const runtime = `

const runtime = {
  log: console.log,
  add: x => y => x + y,
  sub: x => y => x - y,
  mul: x => y => x * y,
  div: x => y => x / y,
  floor: x => Math.floor(x),
  map: fn => arr => arr.map(fn),
  reduce: fn => init => arr => arr.reduce((x,y) => fn(y)(x), init),
  
  init: arr => arr.slice(0, arr.length - 1),
  head: arr => arr[0],
  tail: arr => arr.slice(1),
  last: arr => arr[arr.length - 1],
};
const handler = {
  has: () => true,
};
var p = new Proxy(runtime, handler);

with (p) {

%v

}

`
