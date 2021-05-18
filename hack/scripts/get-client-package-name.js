const apiPathMappings = require('../../.config/apiMappings.json');

const args = process.argv.slice(2);

// get file name argument
if (!args || !args.length) {
  throw new Error("GitHub Repository argument is required");
}

const githubRepo = args[0];

const apiPath = apiPathMappings[githubRepo];
if (!apiPath) {
  throw new Error(`Unknown GitHub repo "${githubRepo}"`)
}

console.log(apiPath.packageName);