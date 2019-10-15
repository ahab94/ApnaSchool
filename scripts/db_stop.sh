#!/bin/bash

if [ "$(docker ps -q -f name=apna-school-mysql-db)" ]; then
  docker rm -f apna-school-mysql-db
fi

if [ "$(docker ps -q -f name=apna-school-mongo-db)" ]; then
  docker rm -f apna-school-mongo-db
fi
