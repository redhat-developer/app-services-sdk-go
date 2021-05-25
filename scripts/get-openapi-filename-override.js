#!/usr/bin/node

const configs = require('../.config/api-client-metadata.json');

const args = process.argv.slice(2);
const clientPayload = JSON.parse(args[0]);

const mappedConfig = configs[clientPayload.id];

if (mappedConfig.openApiOutputFileName) {
    console.log(mappedConfig.openApiOutputFileName)
}
