#!/bin/sh

platforms=("linux/amd64" "darwin/arm64" "darwin/amd64")

for platform in "${platforms[@]}"
do
	split=(${platform//\// })
	GOOS=${split[0]}
	GOARCH=${split[1]}
	file_name="./bin/scry-"$GOOS"-"$GOARCH
	# if [ $GOOS = "windows" ]; then
	# 	file_name+=".exe"
	# fi
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $file_name
	if [ $? -ne 0 ]; then
   	echo "Error building binaries."
		exit 1
	fi
done
