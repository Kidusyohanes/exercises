#!/usr/bin/env bash
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# source the env
source ".env"

echo -e >&2 "${GREEN}starting SongDB...${NC}"

if [ -z "$(docker network ls --filter name=pythonnet --quiet)" ]; then
    echo -e >&2 "${YELLOW}docker network missing; creating it...${NC}"
    docker network create pythonnet
    echo -e >&2 "${YELLOW}done!${NC}"
fi

if [ "$(docker ps -aq --filter name=songdb)" ]; then
    echo -e >&2 "${YELLOW}container exists with name songdb; removing it...${NC}"
	docker rm -f songdb
    echo -e >&2 "${YELLOW}done!${NC}"
fi

docker run -d \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=$MYSQL_PASSWORD \
-e MYSQL_DATABASE=$MYSQL_DATABASE \
--network pythonnet \
--name songdb \
aethan/songdb

# echo -e >&2 "${GREEN}waiting for MySQL to be ready for connections..."
# sleep 7s
echo -e >&2 "complete!${NC}"