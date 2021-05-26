const logError = require("../utils").logError;
const generate = require("../procedures/generate").handler;

exports.command = "generate <language>";
exports.desc = "Generate SDKs for specific language";

exports.builder = {};

exports.handler = async function (argv) {
  logError("Generating " + argv.language);
  generate(argv.language);
};
