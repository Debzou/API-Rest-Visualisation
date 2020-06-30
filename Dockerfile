# Version 
FROM golang:latest 
ENV GO111MODULE=on
# creating a folder inside the docker
RUN mkdir /app 
# Add fill and folder inside the docker
ADD . /app/ 
# CURRENT DIRECTORY (docker) 
WORKDIR /app 
# Library 
RUN go mod download
RUN go build -o main . 
# Launch the application builder
CMD ["/app/main"]