#!/bin/bash
docker rm $(docker ps -a -q)
# sudo docker system prune -a
docker build -t api1.0 .
echo "output api :"
# expose port 8080
docker-compose up
