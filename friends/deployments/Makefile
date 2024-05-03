.PHONY: build-image

build-image:
	docker build --tag taupatupatu-friends --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-friends-container -p 8083:8083 taupatupatu-friends:latest
