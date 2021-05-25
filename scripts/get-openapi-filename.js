#!/usr/bin/node

const path = require('path');

const args = process.argv.slice(2);
const clientPayload = JSON.parse(args[0]);

const openApiFilePath = clientPayload.openapi;

console.log(path.basename(openApiFilePath));
