#!/usr/bin/node

const path = require('path');

const args = process.argv.slice(2);

console.log(path.basename(args[0]));
