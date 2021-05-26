#!/usr/bin/env node
const yargs = require("yargs");

// eslint-disable-next-line no-unused-expressions
yargs
  .commandDir("commands")
  .demandCommand(1)
  .strict()
  .recommendCommands()
  .help()
  .alias("h", "help")
  .version()
  .alias("v", "version").argv;
