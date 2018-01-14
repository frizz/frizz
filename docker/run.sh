#!/bin/bash

# Exit on errors.
set -e

docker-compose build
docker-compose up -d
docker-compose exec frizz fish
