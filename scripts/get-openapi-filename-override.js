#!/usr/bin/node

const configs = require('../.config/api-client-metadata.json');

const args = process.argv.slice(2);

const mappedConfig = configs[args[0]];

if (mappedConfig.openApiOutputFileName) {
    console.log(mappedConfig.openApiOutputFileName)
}
