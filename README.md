# REST-API-GO


## Start API in container :whale:

```
sudo chmod +x ./launch-api.sh
sudo ./launch-api.sh
```

## Start API without docker :space_invader:

Prerequisite : 
- mongo client/server
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
```


