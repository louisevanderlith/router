# router
Mango API: Router

The primary function served by the router is to keep a record of every service and it's location.

## Run with Docker
* $ docker build -t avosa/router:latest .
* $ docker rm RouterDEV
* $ docker run -d -e RUNMODE=DEV -p 8080:8080 --network mango_net --name RouterDEV avosa/router:latest 
* $ docker logs RouterDEV
