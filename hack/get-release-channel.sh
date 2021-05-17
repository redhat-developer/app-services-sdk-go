filename=$1

release_pat="[v]+[1-9](alpha|beta)?(.)(yaml|yml|json)$"
file_pat="[-]$release_pat"

if [[ ! $filename =~ $file_pat ]]; then
  echo "Error: OpenAPI spec file \"$filename\" is missing a version suffix or has invalid syntax."
  exit 1
fi

sed 's'