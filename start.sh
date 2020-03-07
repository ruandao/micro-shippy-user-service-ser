#!/bin/bash

echo "try to ping ${DB_HOST}: ${DB_PORT}";
while ! nc -z "${DB_HOST}" "${DB_PORT}"; do
  echo "try to ping ${DB_HOST}: ${DB_PORT} fail";
  sleep 3;
done

./shippy-service-user