# Version 
FROM golang:latest 
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