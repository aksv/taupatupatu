### Canary deployement for auth service

1. In auth service directory build auth service version with `v1` tag and load to minikube

```
make -f deployments/Makefile build-image-v1
minikube image load taupatupatu-auth:v1
```

2. Deploy v1 version

```
make deploy-canary-v1
```

3. In auth service directory build auth service version with `v2` tag and load to minikube

```
make -f deployments/Makefile build-image-v2
minikube image load taupatupatu-auth:v2
```

4. Keep `v1` version running and deploy `v2`, for tests we split traffic 50/50 percent

```
make deploy-canary-v1
```

5. Test auth service, expected result:

```
 aksv@istanbul: curl 127.0.0.1/auth/
{"Status":"OK","Path":"/","Info":"Version 1"}
 aksv@istanbul: curl 127.0.0.1/auth/
{"Status":"OK","Path":"/","Info":"Version 2"}
```
