#!/usr/bin/env bash

# build frontend
projDir=$(pwd)
cd ${projDir}/frontend || exit
npm run build

cd ${projDir}/frontend_service || exit
npm install

# build backend
cd ${projDir}/backend
export GOROOT=/usr/local/go
export GOOS=linux
export GOARCH=amd64
${GOROOT}/bin/go build -o main .

cd ${projDir}
docker stop api_hub
docker rm api_hub
docker rmi haozzzzzzzz/api_hub:latest
docker build -t haozzzzzzzz/api_hub:latest .
docker run -d --name api_hub -p 3000:3000 -p 18000:18000 haozzzzzzzz/api_hub:latest
