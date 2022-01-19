#!/bin/bash

# please run from api/scripts directory
# arg[1] = project name
# arg[2] = repo/full_name/projec name

if [ "$#" -ne 2 ]; then
    echo "Need new repo name and package folder";
    exit 1;
fi

repo_name=$1
path=$2

# copy files
mkdir ../../$1
cp -r ../* ../../$1
cp -r ../.vscode ../../$1

encoded=$(echo $path | sed 's;/;\\/;g')
encoded=$(echo $encoded | sed 's;\.;\\.;g')
cd ../../$repo_name

# rename proto file
mv ./pkg/proto/proto.proto ./pkg/proto/$repo_name.proto

# mac
find . -type f -print0 -exec sed -i '' "s/github\.com\/keenfury\/api/$encoded/g" {} +
# linux or gnu sed
# find . -type f -print0 -exec sed -i "s/github\.com\/keenfury\/api/$encoded/g" {} +