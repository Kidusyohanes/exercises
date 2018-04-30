#!/bin/echo please source using `source setenv.sh`
export MYSQL_ROOT_PASSWORD=$(openssl rand -base64 32)
export MYSQL_DATABASE=blueit
export MYSQL_ADDR=127.0.0.1
export ADDR=localhost:4000
