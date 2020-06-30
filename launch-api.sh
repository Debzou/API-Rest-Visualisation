#!/bin/bash
sudo docker system prune -a
docker build -t api:1.0 .
echo "output api :"
docker run -p 8093:8093 api:1.0 
