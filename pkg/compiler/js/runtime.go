package js

const runtime = `

const runtime = {
  log: console.log,

  not: x => !x,

  add: x => y => x + y,
  sub: x => y => x - y,
  mul: x => y => x * y,
  div: x => y => x / y,
  floor: x => Math.floor(x),
  gt: x => y => x > y,
  lt: x => y => x < y,
  mod: x => y => x %% y, // temporary double %% as a single causes go to think this is a formatting directive

  filter: fn => arr => arr.filter(fn),
  map: fn => arr => arr.map(fn),
  reduce: fn => init => arr => arr.reduce((x,y) => fn(y)(x), init),
  push: arr => val => [...arr, val],

  do_while: cond => init => fn => {
	let acc = init;
	while (cond(acc)) {
		acc = fn(acc);
	}
	return acc;
  },
  
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
