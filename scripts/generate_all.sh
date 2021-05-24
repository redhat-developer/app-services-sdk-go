
cd .openapi
FILES="*"
for f in $FILES
do
  echo "Generating SDK for $f "
  OPENAPI_FILENAME=$f ../scripts/generate_golang.sh
done