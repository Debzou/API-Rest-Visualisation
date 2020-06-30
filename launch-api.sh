#!/bin/bash
rm main
sudo docker system prune -a
docker build -t api:1.0 .
echo "output api :"
docker run -p 8080:8080 api:1.0 
