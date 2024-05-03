.PHONY: build-image

build-image:
	docker build --tag taupatupatu-notifications --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-notifications-container -p 8085:8085 taupatupatu-notifications:latest
