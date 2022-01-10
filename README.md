# urlshortener

Assignment-InfraCloud

Github: https://github.com/kudligi/urlshortener

DockerHub: https://hub.docker.com/repository/docker/kudligi97/urlshortener

## Endpoints

 1. **POST** /shorten
    > **Request Body Schema**

    > long_url  string, required, url

    > **Response Body Schema**

    > long_url  string, url

    > short_url string, url

    Sample request body : {
    	"long_url" : "https://go.dev/blog/error-handling-and-go"
    }

    Sample response body : {

    "long_url": "https://go.dev/blog/error-handling-and-go",

    "short_url": "http://localhost:4000/38TH3nxg"

    }

same long url will always result in same short url within app lifetime.  

  2. **GET** shorturl
    > Redirects to the registered long url

 3. **POST** /shortenBenchmark
    > **Request Body Schema**

    > **Response Body Schema**

    > long_url  string, url

    > short_url string, url



## CI/CD

Setup Continuous build-test-promote using github actions.


## Design

 - AppLayer
	 - Config loaded from environement variables
	 - Handlers
	 - DataService
	 - In Memory Data Store:
		 - Plain Maps
		 - sync.Map
		 - Map + RwMutex
		 - Map + persistence checkpoints at intervals
	- Utility to generate well distributed random short urls
		- Seeding RNG with cryptographically secure random number
	- Redis based Data Store  


## Unit Testing
Minimal tests for:

 - Router
 - DataService
 - RandomUrl Utility

 TODO : Add more automated testig

## Load Testing
Load Tested using apache bench tool to simulate conncurrent requests under load conditions to validate correctness and measure impace on performance as implimentation was changing.
https://httpd.apache.org/docs/2.4/programs/ab.html

- simultaneous requests to create shorturl for random long urls
	> $ ab.exe -p sample.json  -T application/json -v -c 10 -n 1000 localhost:4000/shortenBenchmark
- simultaneous requests to create shorturl for same long url (supply in sample.json)
	> $ ab.exe -p sample.json  -T application/json -v -c 10 -n 1000 localhost:4000/shorten
- simultaneous requests to redirect for same short url
	> $ab.exe -c 10 -n 1000 **< shortUrl >**

## Distributed version

- Redis as a distributed/persistent state store.
- Stateless App can be scaled horizontally.
- Nginx loadbalancer to distribute load between the app containers
- Docker Compose as orchestrator

# Running the application



## Distributed version
To Start the application with 3 webcontainers:
> $ git clone https://github.com/kudligi/urlshortener

> $ cd urlshortner

> $ docker compose up --scale web=3

Server will be listening on http://localhost:4000/

## Standalone In Memory
To Start the application with 3 webcontainers:

> $ docker run -p 9090:4000 kudligi97/urlshortener:standalone

Server will be listening on http://localhost:4000/


## Standalone In Memory Persistent*
To Start the application with 3 webcontainers:

> $ docker volume create my-vol
>
> $ docker run -p 4000:9090 -v my-vol:/go/app/persist kudligi97/urlshortener:standalone-persistent

Server will be listening on http://localhost:4000/
