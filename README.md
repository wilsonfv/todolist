# todolist
This is a todolist app. <br/>
Built on golang restful api, mongodb as database backend. <br/>
Both go web app server and mongodb are running as docker container. 
There is a batch script will run the CICD on every git changes update.

# Prerequisite
### Development Setup
##### Install Go
Install go on your workstation [https://golang.org/doc/install](https://golang.org/doc/install)  

##### Install Dep
Dep is go dependency management tool [https://github.com/golang/dep](https://github.com/golang/dep)

##### Install Docker
Install docker CE on your workstation [https://docs.docker.com/install/](https://docs.docker.com/install/)

### Server Setup 
Install docker CE on your server [https://docs.docker.com/install/](https://docs.docker.com/install/)

# Development
On project root folder, run dep ensure to download all dependencies defined in Gopkg.toml
```
dep ensure
```

Then you can start up the go app-server

```go
go run -v app/app_server.go
```

Run the unit test
```go
go test ./...
```

Create a user-defined bridge network so that app-server can access mongodb container by the container name. Refer to [https://docs.docker.com/network/bridge/](https://docs.docker.com/network/bridge/)
```
docker network create app-net
``` 

Startup mongodb container. Create a directory on the host system (Such as on your workstation /opt/dev/mongo) as an external volume mount for the mongodb so that the database content will persist across container lifecycle. Refer to [https://docs.docker.com/storage/#tips-for-using-bind-mounts-or-volumes](https://docs.docker.com/storage/#tips-for-using-bind-mounts-or-volumes)
```
docker run --name mongodb --network app-net --publish 27017:27017 -v /opt/dev/mongo:/data/db -it mongo:3.
```

Startup go web server container. The mongodb url will be passed into the server.
```
docker run --name app-server --network app-net --publish 8181:8181 -d todolist:latest go run -v app/app_server.go -mongodbUrl mongodb:27017
```

#Restful API
To get all the tasks
```
curl -X GET http://localhost:8181/task
```

To create a new task
```
curl -X POST \
  http://localhost:8181/task \
  -d '{
	"name": "task1",
	"description": "task description 1"
}'
```

To delete an existing task
```
curl -X DELETE \
  http://localhost:8181/task \
  -d '    {
        "id": "5bb4baa67d1394005195877e",
        "name": "task8",
        "creation_date": 6608111775386697728,
        "description": "task description 8"
    }'
}'
```