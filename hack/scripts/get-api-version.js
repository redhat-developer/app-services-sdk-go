const args = process.argv.slice(2);

// get file name argument
if (!args || !args.length) {
  throw new Error("File name argument is required");
}

const fileName = args[0];

const releasePatternFileSuffix = '?(.)(yaml|yml|json)$'
const releasePattern = `[v]+[1-9](alpha|beta)?`;
const releaseFilePattern = `[-]${releasePattern}${releasePatternFileSuffix}`

if (!fileName.match(releaseFilePattern)) {
  throw new Error(`Invalid syntax in "${fileName}": missing or incorrect version format.`);
}

const match = fileName.match(releasePattern);
if (!match) {
  throw new Error("Missing version number in " + fileName)
}

console.log(match[0]);