#!/usr/bin/node

const path = require('path');

const args = process.argv.slice(2);
const clientPayload = JSON.parse(args[0]);

const downloadUrl = clientPayload.download_url;

console.log(path.basename(downloadUrl));
