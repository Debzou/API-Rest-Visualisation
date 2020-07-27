# REST-API-GO

The security of the api is based on Gin-JWT.

# Launch

## Start API in container :whale:
change url mongodb in main.go (mongodb://127.0.0.1:27017)
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

change url mongodb in main.go (mongodb://mongo:27017/)

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
# the status

 When a user creates an account on the api, the user is automatically in "mormal_user".
 
- normal_user : this status prevents middleware authorization.

- admin : this status allows to use the middlware.

To have the administrator status, you have to log in to mongodb and modify your profile.


