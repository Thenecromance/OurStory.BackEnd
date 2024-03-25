#!/bin/bash


FILE=./node_modules

echo "Building the frontend"
#Build the frontend first
cd FrontEnd || exit

if [ ! -d "$FILE" ];then
  echo "$FILE does not exist"
  npm install
fi

npm run build
cd ..

echo "Building the backend"
#then run the backend
go run main.go