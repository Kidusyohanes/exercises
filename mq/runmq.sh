#!/usr/bin/env bash
GREEN='\033[1;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

CONTAINER_NAME=mq

#if there is an existing container with our container name,
#remove it
if [ "$(docker ps -aq --filter name=$CONTAINER_NAME)" ]; then
    echo -e "${YELLOW}removing existing MQ container...${NC}"
    docker rm -f $CONTAINER_NAME
fi

#run the rabbitmq container
echo -e "${YELLOW}running MQ container...${NC}"
docker run -d \
-p 127.0.0.1:5672:5672 \
--name $CONTAINER_NAME \
rabbitmq:3-alpine

echo -e "${GREEN}you are ready to pub/sub!${NC}"
