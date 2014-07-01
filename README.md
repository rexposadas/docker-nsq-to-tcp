An nsq consumer inside a docker container.  This was created primarily because 
LogStash did not have built-in support for nsq. 

It makes use of the google/golang [imagea](https://registry.hub.docker.com/u/google/golang/) to ensure the usage of the latest Go version.


Given:

    my docker host is 172.17.42.1
    LogStash (or any TCP server) is listening on 7000

Then my run command would look like this: 

    docker run --name consumer -e TOPIC=test \
    -e LOOKUPD_ADDR=172.17.42.1:4161 \
    -e OUTPUT_TCP_ADDR=172.17.42.1:7000 \
    rexposadas/docker-nsq-to-tcp 

