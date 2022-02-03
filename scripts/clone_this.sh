#!/bin/bash

# please run from project_root/scripts directory
# arg[1] = repo/full_name/project name

if [ "$#" -ne 1 ]; then
    echo "Need new package folder structure (i.e: github.com/<name>/<project>";
    exit 1;
fi

path=$1
name=$(basename $1)
uppername=$(echo "$name" | tr '[:lower:]' '[:upper:]')

# copy files
mkdir ../../$name
cp -r ../* ../../$name
cp -r ../.vscode ../../$name

encoded=$(echo $path | sed 's;/;\\/;g')
encoded=$(echo $encoded | sed 's;\.;\\.;g')
cd ../../$name

# rename proto file
mv ./pkg/proto/proto.proto ./pkg/proto/$name.proto

# replace config
sed -i '' "s/go-api-base/$name/g" ./config/config.go
sed -i '' "s/GO_API_BASE/$uppername/g" ./config/config.go

# replace rest/main.go
sed -i '' "s/GO_API_BASE/$uppername/g" ./cmd/rest/main.go

# mac
find . -type f -print0 -exec sed -i '' "s/github\.com\/keenfury\/go-api-base/$encoded/g" {} +
# linux or gnu sed
# find . -type f -print0 -exec sed -i "s/github\.com\/keenfury\/go-api-base/$encoded/g" {} +
echo -e "\nProject '$name' cloned..."