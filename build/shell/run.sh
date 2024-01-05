#!/bin/bash

APP_NAME="code-go"
Current_Dir="$(pwd)"

echo "Usage: "
echo "    sh ./run.sh start    start $APP_NAME server"
echo "    sh ./run.sh stop     stop $APP_NAME server"
echo "    sh ./run.sh restart  restart $APP_NAME server"
echo "Execute $APP_NAME shell, current dir:"
echo "    $Current_Dir"

DIR="$(cd "$(dirname "$0")" && pwd)"
cd ../../

APP_PATH="$(dirname "$(dirname "$DIR")")"
echo "$APP_PATH"

start(){
  echo "Starting ${APP_NAME}..."
  go build -o $APP_NAME main.go
  # 赋予执行权限
  chmod 777 "$APP_PATH"
  # 重定向到/dev/null设备文件，不产生任何输出
#  nohup ${APP_PATH}/${APP_NAME} > /dev/null 2>&1 &
  nohup "${APP_PATH}"/${APP_NAME} &
  echo "Start ${APP_NAME} successfully. Click me: http://127.0.0.1:28080"
}

stop() {
    echo "Stopping ${APP_NAME}..."
    # 通过 grep 过滤出包含 myapp 的行，通过 grep -v 命令排除自身的 grep 进程，最后通过 awk 命令提取 PID
    pid=$(ps -ef | grep ${APP_NAME} | grep -v grep | awk '{print $2}')
    if [ -n "${pid}" ]; then
        echo "Killing $APP_NAME (PID: $pid)...."
        kill -9 "${pid}"
        ret=0
        #多次循环杀掉进程
        for ((i=1;i<=10;i++)); do
          sleep 1
          pid=$(ps -ef | grep $APP_NAME | grep -v grep | awk '{print $2}')
          if [ "$pid" ]; then
              kill -9 "$pid"
              ret=0
          else
              ret=1
              break
          fi
        done

        if [ "$ret" ]; then
                echo -e $"ok"
        else
                echo -e $"no"
        fi
    else
      echo "Process $APP_NAME not exist"
    fi
    echo "${APP_NAME} stopped."
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    *)
        exit 1
        ;;
esac