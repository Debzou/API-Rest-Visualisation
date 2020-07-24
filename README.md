# REST-API-GO

The security of the api is based on Gin-JWT.

## Start API in container :whale:

```
sudo chmod +x ./install.sh
sudo ./install.sh
```

## Start API without docker :space_invader:

Prerequisite : 
- mongo client/server 

```sh
docker run -d -p 27017-27019:27017-27019 --name RESTmongo  mongo
```

- go lastest version


```sh
go mod init github.com/Debzou/REST-API-GO
go mod vendor
go run main.go
```

or 

```sh
sudo chmod +x ./packaged-api.sh
sudo ./packaged-api.sh
go run main.go
```

