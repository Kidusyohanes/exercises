#!/bin/echo please source using `source setenv.sh`
export MYSQL_ROOT_PASSWORD=$(openssl rand -base64 32)
export MYSQL_DATABASE=tasks
export MYSQL_ADDR=127.0.0.1:3306
export ADDR=localhost:4000
