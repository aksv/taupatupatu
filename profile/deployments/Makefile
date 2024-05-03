.PHONY: build-image

IMAGE_TAG ?= latest

build-image:
	docker build --tag taupatupatu-profile:${IMAGE_TAG} --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-profile-container -p 8086:8086 taupatupatu-profile:latest
