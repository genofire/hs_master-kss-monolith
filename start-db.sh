#!/bin/bash
[[ -n "$DOCKER_HOST" ]] && echo "DOCKER_HOST: $DOCKER_HOST"
if [[ ! -n "$LAB_MONOLITH_DB_IP" ]]; then
  echo 'ERROR: LAB_MONOLITH_DB_IP is not set.'
  exit 1
else
  echo "LAB_MONOLITH_DB_IP: $LAB_MONOLITH_DB_IP"
fi
if [[ ! -n "$LAB_MONOLITH_DB_PASSWORD" ]]; then
  echo 'ERROR: LAB_MONOLITH_DB_PASSWORD is not set.'
  exit 1
fi
echo 'Starting database for integration tests...'
docker run --rm --name monolith-postgres\
 -p 15432:5432\
 -e POSTGRES_USER=monolith\
 -e POSTGRES_PASSWORD=$LAB_MONOLITH_DB_PASSWORD\
 -d postgres:9
