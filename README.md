# router
Mango API: Router

The primary function served by the router is to keep a record of every service and it's location.

The router API is used hold all versions of an application or API's URLs
For example, we can use the same router for all of our environments like 'LIVE', 'UAT', and 'DEV'
This means that we can ask the router API for our e-mail API (Comms.API) and depending on the environment of the caller
the API, we will get the correct URL.
The functionality provided by this API also ensures that we can't all anything from 'LIVE' when running on 'DEV',
and may be seen as a way of keeping developers safe.

In order for the router to know about a Database, API or Application, we have to register with the router API on start-up.
We don't need to store URLs anywhere in our applications or JavaScript.
We just need to know what application we want.
This decreases effort when deploying our applications and also keeps the code free from URLs which could cause problems in the future.

You will see 'srv.Register(port)' within the main.go of every application.
This function is used to register an application.

## Run with Docker
* $ docker build -t avosa/router:dev .
* $ docker rm RouterDEV
* $ docker run -d -p 8080:8080 --network mango_net --name RouterDEV avosa/router:dev 
* $ docker logs RouterDEV
