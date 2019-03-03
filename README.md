# router
Mango API: Router

The primary function served by the router is to keep a record of every service and it's location.

#Please note!
The --network-alias MUST be set to 'theRouter'
Also please ensure that all containers are running on the same bridge network.

## Run with Docker
*$ go build
*$ docker build -t avosa/router:latest .
*$ docker rm routerDEV
*$ docker run -d -e RUNMODE=DEV -p 8080:8080 --network mango_net --name routerDEV avosa/router:latest 
*$ docker logs routerDEV
