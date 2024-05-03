deploy-canary-v1:
	minikube addons enable ingress
	kubectl --context minikube apply -f ./deployment-v1.yaml
deploy-canary-v2:
	minikube addons enable ingress
	kubectl --context minikube apply -f ./deployment-canary.yaml