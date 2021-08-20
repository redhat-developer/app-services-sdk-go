echo "Patching service registry instances"

cd .openapi
echo "Removing codegen "
cat registry-instance.json | jq 'del(.paths."x-codegen-contextRoot")' > registry-instance-tmp.json
mv -f registry-instance-tmp.json registry-instance.json

echo "Ensuring only single tag is created "
cat registry-instance.json | jq 'walk( if type == "object" and has("tags") 
       then .tags |= select(.[0])
       else . end )' > registry-instance-tmp.json
mv -f registry-instance-tmp.json registry-instance.json

echo "Removing invalid datetime definitions"
sed -i '' 's/date-time/utc-date/' registry-instance.json

