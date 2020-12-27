# go-dead-drop

Project state: MVP
--------

Aims to help sharing information in secure, fast and comfortable way. Its security approach based on inability of service to decrypt stored secrets by itself.

Project is in a very early stage. Probably not production ready.

Runs on any possible platform (written in golang, one executable) - standalone, docker, kubernetes, lambda with different storage back-ends: file, redis, rocksdb, mysql, mongodb, cloud bucket and etc.

### Local setup:
Developed with go1.14, but will probably run on earlier versions.
1. Run `docker run -p 6379:6379 --name local_redis -d redis:5-alpine` (or change in config.yaml to use localfile.)
2. Run `bash dev_build.sh` or `docker run -p 8080:8080 themoah:go-dead-drop:latest`

### How to use:

1. Send HTTP post to store , e.g. `curl --data-binary @wow.sh http://localhost:8080/store`
2. Use generated params to retrieve `curl -X POST http://localhost:8080/retrieve/{KEY}/{PASSWORD}`