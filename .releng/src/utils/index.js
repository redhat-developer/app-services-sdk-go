const chalk = require('chalk');

// eslint-disable-next-line no-console
const log = console.log;
exports.logError = (s) => log(chalk.red(s));
exports.logInfo = (s) => log(chalk.green(s));
exports.logDetail = (s) => log(chalk.dim(s));