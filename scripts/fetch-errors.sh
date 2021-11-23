ERRORS_FILE_URL=$1
ERRORS_FILE=$(dirname $0)/../.errors/errors_${CLIENT_ID}.go

echo "Fetching remote file to " + $ERRORS_FILE
# download the OpenAPI file
curl -H "Authorization: token $BF2_TOKEN" "$ERRORS_FILE_URL" --output $ERRORS_FILE
if [ $? != 0 ]; then
  exit $?
fi