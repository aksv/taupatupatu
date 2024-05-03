.PHONY: build-image

build-image:
	docker build --tag taupatupatu-content --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-content-container -p 8082:8082 taupatupatu-content:latest
