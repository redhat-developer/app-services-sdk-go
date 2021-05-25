#!/usr/bin/node

const path = require('path');
const configs = require('../.config/api-client-metadata.json');

const args = process.argv.slice(2);
const clientPayload = JSON.parse(args[0]);

const mappedConfig = configs[clientPayload.id];
const repo = mappedConfig.source.repo;
const baseBranch = mappedConfig.source.base || 'main';
const openApiFilePath = clientPayload.openapi;

const rawOpenApiURL = path.join("https:///raw.githubusercontent.com", repo,
  baseBranch,
  openApiFilePath
);

console.log(rawOpenApiURL);