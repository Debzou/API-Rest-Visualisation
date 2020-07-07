#!/bin/bash
docker rm $(docker ps -a -q)
chmod +x ./packaged-api.sh
./packaged-api.sh
# sudo docker system prune -a
docker build -t api1.0 .
# expose port 8080
docker-compose up
