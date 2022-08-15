#!/bin/sh

stop(){
  echo "查找 $1 server"
  pid=$(ps -ef|grep ../bin/$1 | grep -v grep |awk '{print $2}')

  echo "关闭 $1 server, Pid = $pid"
  kill -9 $pid
}

stop world
stop login
stop front
stop chat
