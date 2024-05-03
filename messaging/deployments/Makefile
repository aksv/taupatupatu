.PHONY: build-image

build-image:
	docker build --tag taupatupatu-messaging --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-messaging-container -p 8084:8084 taupatupatu-messaging:latest
