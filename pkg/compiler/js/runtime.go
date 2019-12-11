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
};
const handler = {
  has: () => true,
};
var p = new Proxy(runtime, handler);

with (p) {

%v

}

`
