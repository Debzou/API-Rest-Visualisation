#!/bin/bash

#service mongodb stop (if err --> Cannot start service mongo)
docker-compose down
chmod +x ./packaged-api.sh
./packaged-api.sh
docker-compose up --build
