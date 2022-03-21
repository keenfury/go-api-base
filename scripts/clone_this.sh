#!/bin/bash

# please run from project_root/scripts directory
# arg[1] = repo/full_name/project name

if [ "$#" -ne 1 ]; then
    echo "Need new package folder structure (i.e: github.com/<name>/<project>";
    exit 1;
fi

path=$1
full_path=$GOPATH/src/$1
name=$(basename $1)
upper_name=$(echo "$name" | tr '[:lower:]' '[:upper:]')
upper_name_env="${upper_name//-/_}"

# copy files
mkdir -p $full_path
cp -r ../* $full_path
cp -r ../.vscode $full_path

encoded=$(echo $path | sed 's;/;\\/;g')
encoded=$(echo $encoded | sed 's;\.;\\.;g')
cd $full_path

# rename proto file
mv ./pkg/proto/proto.proto ./pkg/proto/$name.proto

# replace config
sed -i '' "s/go-api-base/$name/g" ./config/config.go
sed -i '' "s/GO_API_BASE/$upper_name_env/g" ./config/config.go

# replace rest/main.go
sed -i '' "s/GO_API_BASE/$upper_name/g" ./cmd/rest/main.go

# mac
find . -type f -print0 -exec sed -i '' "s/github\.com\/keenfury\/go-api-base/$encoded/g" {} +
# linux or gnu sed
# find . -type f -print0 -exec sed -i "s/github\.com\/keenfury\/go-api-base/$encoded/g" {} +
echo -e "\nProject '$name' cloned..."