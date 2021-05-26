const { logError, logInfo } = require("../utils");
const fetch = require("../procedures/fetch").handler;

exports.command = "fetch [id] [location]";

exports.desc = `Fetch Github file from remote repository.
Command expects .config/api-client-metadata.json to exist in current working directory
Env vars: GITHUB_TOKEN, CLIENT_PAYLOAD='{"id":"kafka-mgmt/v1","openapi":"./openapi/kas-fleet-manager.yaml"}'
`;

exports.builder = {};

exports.handler = async function (argv) {
  if (argv.id && argv.location) {
    fetch(argv.id, argv.location);
  } else if (process.env.CLIENT_PAYLOAD) {
    logInfo(`Fetching using CLIENT_PAYLOAD env: ` + process.env.CLIENT_PAYLOAD);
    const data = JSON.parse(process.env.CLIENT_PAYLOAD);
    fetch(data.id, data.openapi);
  } else {
    logError("Missing arguments. Set CLIENT_PAYLOAD or use positional args");
  }
};
