echo "Patching service registry instances"

jq 'del(.paths."x-codegen-contextRoot")' ./.openapi/registry-instance.json
jq 'walk( if type == "object" and has("tags") 
      then .tags |= select(.[0])
      else . end )' ./.openapi/registry-instance.json
