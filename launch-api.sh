#!/bin/bash
docker rmi api-rest_api 
chmod +x ./packaged-api.sh
./packaged-api.sh
docker-compose up
