# go-dead-drop

[![Build](https://github.com/themoah/go-dead-drop/workflows/Build/badge.svg)](https://github.com/themoah/go-dead-drop/actions)
[![CodeQL](https://github.com/themoah/go-dead-drop/workflows/CodeQL/badge.svg)](https://github.com/themoah/go-dead-drop/actions)

Project state: MVP
--------

Aims to help sharing information in secure, fast and comfortable way. Its security approach based on inability of service to decrypt stored secrets by itself.

Project is in a very early stage. Probably not production ready.

Runs on any possible platform (written in golang, one executable) - standalone, docker, kubernetes, lambda with different storage back-ends: file, redis, rocksdb, mysql, mongodb, cloud bucket and etc.

### Local setup:
Developed with go1.14, but will probably run on earlier versions.
1. Run `bash dev_build.sh` or `docker run -p 8080:8080 themoah:go-dead-drop:latest`

### How to use:

1. Send HTTP post to store , e.g. `curl --data-binary @wow.sh http://localhost:8080/store`
2. Use generated params to retrieve `curl -X POST http://localhost:8080/retrieve/{KEY}/{PASSWORD}`

### Additional configurations:

Basic configuration uses in-memory store. You can also persist data at Redis or local filesystem(not recommended for non-development environments).
Change storageEngine in config.yaml to "redis" and provide connection details.