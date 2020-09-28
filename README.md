# metrics-server

Metrics server is a REST server that is capable of ingesting metrics from 
clients and the metrics are stored in memory.

The server exposes the following two endpoints: 
- `/metrics`: Client can post metrics in JSON format so that it is ingested
   by the server. 
- `/report`: One can do a `GET` request on this endpoint to see the max cpu
  and memory usage till time by nodes.

To know more on how to use this metrics server please follow the quick start 
guide in the below section.

## Quick Start Guide

**Prerequsites** : 
- Docker should be installed on the node.

The container images for this server is published here : 
https://hub.docker.com/repository/registry-1.docker.io/sonasingh46/metrics-server-amd64/tags?page=1

**Steps**
- **Run the following command to run the docker image**
  ```bash
  docker run -p 8080:8080 sonasingh46/metrics-server-amd64:b5fbd7c7a19d14612440bb94e971312b6d490f5c
  ```

Note: The image tag `b5fbd7c7a19d14612440bb94e971312b6d490f5c` is used while this
doc is being written. You can use the latest docker image pushed from the docker
hub link pasted above.
A merge against master branch builds and publishes a new image to the docker hub. 

- **Steps for posting the metrics:** 

Assuming that it is deployed locally using the above docker run command. Metrics
  can be pushed by doing a POST request.
  Following is an example : (You could use CURL or POSTMAN)
  The following is a sample metric payload/body for POST request.
  ```json
  {
	"node_ip":"10.10.2.4",
	"percentage_cpu_used":24,
	"percentage_memory_used":16
  }
  ```
 The following header should be passed:
 ```bash
 Content-type: application/json
 ```

The following is the URL: 
```bash
localhost:8080/metrics
```

So the complete CURL command is following : 
```bash
curl -X POST \
  http://localhost:8080/metrics \
  -H 'content-type: application/json' \
  -d '{
	"node_ip":"10.10.2.4",
	"percentage_cpu_used":24,
	"percentage_memory_used":16
}'
```


- **Steps for getting  the report:** 

Following is the CURL command: 
```bash
curl -X GET \
  http://localhost:8080/report \
  -H 'content-type: application/json' \

```

**Sample Test Cases:**

To try/test out this server, a bash script has been written.
You can find it [here](hack/test.sh) 
Read the bash script and execute it to verify the results.

## Building it locally

1. Clone this github repo
2. Run `make test` if you have done code changes. 
3. Run `make metrics-server.amd64` to build a docker image

## Limitations and ToDo

The server does not handle the concurrency aspects as of now. If there are
concurrent request for the server to ingest the metrics, there is possibility
of race conditions. This is a TODO item.