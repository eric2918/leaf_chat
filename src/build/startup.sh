#!/bin/sh
go mod tidy
start(){
  #echo "编译 $1 server"
  #go build -o bin/$1 cmd/$1/main.go

  echo "启动 $1 server"
#  nohup ./bin/$1 $2 >> /dev/null 2>&1 &
  nohup ../bin/$1 ../bin/conf/$2 >> ../bin/logs/$1/output.log 2>&1 &
}

start world world.json
start login login.json
start front front1.json
start chat chat1.json
