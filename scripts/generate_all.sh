
FILES=".openapi/*"
for f in $FILES
do
  OPENAPI_FILENAME=`basename $f`
  echo "Generating API for $OPENAPI_FILENAME "
  . ./scripts/generate_go.sh
  echo "Finished generation for API $OPENAPI_FILENAME "
done