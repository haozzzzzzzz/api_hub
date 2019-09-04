#!/usr/bin/env bash

# build frontend
buildDir=$(pwd)
cd ../frontend || exit
npm run build

cd "${buildDir}" || exit
cp -r ../frontend/dist ./

# build backend
cd ../backend
cp -r project/workspace/dev_local/config "${buildDir}"/
export GOROOT=/usr/local/go
export GOOS=linux
export GOARCH=amd64
${GOROOT}/bin/go build -o "${buildDir}"/main .

cd ${buildDir}

docker build -t haozzzzzzzz/api_hub:latest .
docker stop api_hub
docker rm api_hub
docker run -d --name api_hub -p 3000:3000 -p 18000:18000 haozzzzzzzz/api_hub:latest