package js

const runtime = `

const runtime = {
  log: console.log,
  add: x => y => x + y,
};
const handler = {
  has: () => true,
};
var p = new Proxy(runtime, handler);

with (p) {

%v

}

`
