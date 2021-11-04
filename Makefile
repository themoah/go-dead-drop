
## dev		: local build
dev:
	bash dev_build.sh

## build : build a version
build:
	go build -o main


buildDocker:
	docker build -t dead-drop:t .

runDocker:
	docker run -p 8080:8080 dead-drop:t

buildDockerRelease:
	docker build -t themoah/go-dead-drop:latest