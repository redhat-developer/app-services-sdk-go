echo "Patching service registry instances"

echo "Removing codegen "
cat ./.openapi/registry-instance.json | jq 'del(.paths."x-codegen-contextRoot")'  > ./.openapi/registry-instance-tmp.json
echo "Ensuring only single tag is created "
cat ./.openapi/registry-instance.json | jq 'walk( if type == "object" and has("tags") 
       then .tags |= select(.[0])
       else . end )' > ./.openapi/registry-instance-tmp.json

mv -f ./.openapi/registry-instance-tmp.json  ./.openapi/registry-instance.json