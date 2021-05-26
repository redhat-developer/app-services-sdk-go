#!/usr/bin/node

const args = process.argv.slice(2);
const clientPayload = JSON.parse(args[0]);

console.log(clientPayload.download_url);