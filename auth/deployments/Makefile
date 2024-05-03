.PHONY: build-image

build-image:
	docker build --tag taupatupatu-auth --file ./deployments/Dockerfile .
build-image-v1:
	docker build --tag taupatupatu-auth:v1 --file ./deployments/Dockerfile .
build-image-v2:
	docker build --tag taupatupatu-auth:v2 --file ./deployments/Dockerfile .
run-service:
	docker container run --rm --name taupatupatu-auth-container -p 8081:8081 taupatupatu-auth:latest
load-image: build-image
	eval $(minikube docker-env)
	minikube image load taupatupatu-auth
