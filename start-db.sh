#!/bin/bash

[[ -n "$DOCKER_HOST" ]] && echo "DOCKER_HOST: $DOCKER_HOST"

if [[ ! -n "$LAB_MONOLITH_DB_HOST" ]]; then
  echo 'ERROR: LAB_MONOLITH_DB_HOST is not set.'
  exit 1
else
  echo "LAB_MONOLITH_DB_HOST: $LAB_MONOLITH_DB_HOST"
fi

if [[ ! -n "$LAB_MONOLITH_DB_PORT" ]]; then
  echo 'ERROR: LAB_MONOLITH_DB_PORT is not set.'
  exit 1
else
  echo "LAB_MONOLITH_DB_PORT: $LAB_MONOLITH_DB_PORT"
fi

if [[ ! -n "$LAB_MONOLITH_DB_PASSWORD" ]]; then
  echo 'ERROR: LAB_MONOLITH_DB_PASSWORD is not set.'
  exit 1
fi

echo 'Starting database...'

docker run --rm --name monolith-postgres\
 -p $LAB_MONOLITH_DB_PORT:5432\
 -e POSTGRES_USER=monolith\
 -e POSTGRES_PASSWORD=$LAB_MONOLITH_DB_PASSWORD\
 -d postgres:9
