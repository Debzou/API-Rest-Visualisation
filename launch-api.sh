#!/bin/bash
rm main
docker rm $(docker ps -a -q)
# sudo docker system prune -a
docker build -t api:1.0 .
echo "output api :"
# expose port 8080
docker run -p 8080:8080 api:1.0 
