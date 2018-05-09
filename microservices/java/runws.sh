#!/usr/bin/env bash
docker rm -f javaws

docker run -d \
-p 80:80 \
-e JAVA_MYSQL_ADDR=tasksdb:3306 \
-e JAVA_MYSQL_DB=$JAVA_MYSQL_DB \
-e JAVA_MYSQL_PASS=$JAVA_MYSQL_PASS \
--network javanet \
--name javaws \
brendankellogg/jwsdemo
