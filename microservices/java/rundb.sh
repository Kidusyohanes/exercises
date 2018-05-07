#!/usr/bin/env bash
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e >&2 "${GREEN}starting TasksDB...${NC}"

if [ -z "$(docker network ls --filter name=javanet --quiet)" ]; then
    echo -e >&2 "${YELLOW}docker network missing; creating it...${NC}"
    docker network create javanet
    echo -e >&2 "${YELLOW}done!${NC}"
fi

if [ "$(docker ps -aq --filter name=tasksdb)" ]; then
    echo -e >&2 "${YELLOW}container exists with name tasksdb; removing it...${NC}"
	docker rm -f tasksdb
    echo -e >&2 "${YELLOW}done!${NC}"
fi

docker run -d \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=$JAVA_MYSQL_PASS \
-e MYSQL_DATABASE=$JAVA_MYSQL_DB \
--name tasksdb \
--network javanet \
drstearns/tasksdb

echo -e >&2 "${GREEN}waiting for MySQL to be ready for connections..."
sleep 7s
echo -e >&2 "complete!${NC}"