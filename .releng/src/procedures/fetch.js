const { execSync } = require("child_process");
const { readFileSync } = require("fs");
const path = require("path");

/**
 * Fetch file to local filesystem
 *
 * @param {*} id Of the artifact we fetch
 * @param {*} repositoryFile location in repository
 */
exports.handler = async function (id, repositoryFile) {
  const config = readFileSync(".config/api-client-metadata.json", {
    encoding: "utf8",
  });

  const configProperties = JSON.parse(config);
  const mappedConfig = configProperties[id];
  const repo = mappedConfig.source.repo;
  const baseBranch = mappedConfig.source.base || "main";
  const openApiFilePath = repositoryFile;
  const openApiOutputFileName = mappedConfig.openApiOutputFileName;

  const rawOpenApiURL = path.join(
    "https:///raw.githubusercontent.com",
    repo,
    baseBranch,
    openApiFilePath
  );

  console.log(`Fetch data from ${rawOpenApiURL} to ${openApiOutputFileName}`);

  if (mappedConfig.source.public) {
    console.log(
      execSync(`curl "${rawOpenApiURL}" -o "${openApiOutputFileName}"`)
    );
  } else if (process.env.GITHUB_TOKEN) {
    const command = `curl -H "Authorization: token "${process.env.GITHUB_TOKEN}" "${rawOpenApiURL}" -o "${openApiOutputFileName}"`;
    console.log(execSync(command));
  } else {
    console.log("Repository needs to be public or provide GITHUB_TOKEN");
  }
};
