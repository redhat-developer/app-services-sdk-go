## List of the folders common across different SDK's 
## This file targets to provide list of the directories that are common across all SDKS
## Note you are in source repo. This script is not intended to be used here
source=../app-services-sdk-go
folders="${source}/.chglog ${source}/.config ${source}/.github ${source}/scripts"

cp -r ${folders} ./