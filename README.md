## Build
``` shell
$ go build -o msmgr ./src/ # in fish
```

## How to use
``` shell
# ./msmgr http://ms:Microservice123!@20.26.33.122:32010
$ ./msmgr http://<usr>:<passwd>@<Eureka_IP>:<Eureka_Port>
```

## Build docker
``` shell
$ cd msmgr
$ env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o msmgr ./src/ # in fish
$ upx msmgr # compress binary
$ docker build -t msmgr:latest -f ./deploy/Dockerfile .
```

Ref:
https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker

## Run docker
``` shell
$ docker run msmgr "./msmgr http://<usr>:<passwd>@<Eureka_IP>:<Eureka_Port>"
```
