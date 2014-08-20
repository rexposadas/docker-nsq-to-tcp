An nsq consumer inside a docker container.  This was created primarily because 
LogStash did not have built-in support for nsq. Basically, this allows us to use nsq
as a broker for LogStash.  

This image is *not* limited to being a LogStash broker. Any service which accepts TCP messages can 
make use of this image.

This image uses the google/golang [image](https://registry.hub.docker.com/u/google/golang/) 
to ensure the usage of the latest Go version.


Given:

    my docker host is 172.17.42.1
    LogStash (or any TCP server) is listening on 7000

Then my run command would look like this: 

    docker run -e TOPIC=test 
      -e LOOKUPD_ADDR=172.17.42.1:4161 
      -e OUTPUT_TCP_ADDR=172.17.42.1:7000 
      rexposadas/docker-nsq-to-tcp 

Note that I am using the default port for nsqlookupd.  Your configuration might differ.

