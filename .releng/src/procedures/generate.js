const { execSync } = require("child_process");
const { readFileSync } = require("fs");
const path = require("path");

/**
 * Fetch file to local filesystem
 *
 * @param {*} id Of the artifact we fetch
 * @param {*} repositoryFile location in repository
 */
exports.handler = async function (lang) {
  const config = readFileSync(".config/api-client-metadata.json", {
    encoding: "utf8",
  });

  if (!lang) {
    lang = "go";
  }
  const configProperties = JSON.parse(config);
  for (const currentConfig of Object.values(configProperties)) {
    console.log(currentConfig);
    if (lang == "go") {
      generateGo(currentConfig);
    } else {
      throw new Error("Language not implemented " + lang);
    }
  }
};

function generateGo(currentConfig) {
  console.log(
    execSync(
      `npx @openapitools/openapi-generator-cli version-manager set 5.1.1`,
      { encoding: "utf8" }
    )
  );

  if (!currentConfig.go) {
    throw new Error(
      "Mising golang specific config: " +
        JSON.stringify(currentConfig, undefined, 2)
    );
  }
  const configForLang = currentConfig.go;
  const openApiFileName = currentConfig.openApiOutputFileName;
  const outputPath = configForLang.apiGroup + "/" + configForLang.apiVersion;
  const mockFile = outputPath + "/default_api_mock.go";

  console.log("Generating SDK based on " + openApiFileName);

  const cmd =
    `npx @openapitools/openapi-generator-cli generate -g go ` +
    `-i "${openApiFileName}" -o "${outputPath}" --package-name "${configForLang.apiGroup}" ` +
    `--git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/${configForLang.apiGroup}" ` +
    `-p "generateInterfaces=true" --ignore-file-override=.openapi-generator-ignore`;
  console.log(cmd);
  execSync(cmd, { encoding: "utf8" });

  console.log(execSync(`moq -out "${mockFile}" "${outputPath}" DefaultApi`));
}
