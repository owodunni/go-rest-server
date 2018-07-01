# go-rest-server
A simple rest server writen in golang.

#### Getting started 
Install docker

	docker build -t go-rest-server .
	docker run --publish 6060:8080 --name test --rm go-rest-server

Now goind to localhost:6060/test should give you a response

To stop the docker container run
	
	docker stop test
