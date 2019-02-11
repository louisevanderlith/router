# router
Mango API: Router

The primary function served by the router is to keep a record of every service and it's location.

## Run with Docker
*$ go build
*$ docker build -t avosa/router:dev .
*$ docker rm routerDEV
*$ docker run -d --network host --name routerDEV avosa/router:dev 
*$ docker logs routerDEV
