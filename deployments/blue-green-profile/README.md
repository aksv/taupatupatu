### Blue green deployment for profile service

1. Deploy version 1 of `profile service`

```
make deploy-blue
```

2. Enable port forwarding

```
kubectl --context minikube  port-forward service/profile-service-blue 8086:80
```

3. Make sure that version 1 of the service is running.

```
aksv@istanbul: curl 127.0.0.1:8086
{"Status":"OK","Version":"Profile V1"}
```

4. Deploy service version 2

```
make deploy-green
```

5. Forward traffic to green (v2) version.

```
kubectl --context minikube  port-forward service/profile-service-green 8086:80
```

6. Make sure that V2 is handling our requests

```
aksv@istanbul: curl 127.0.0.1:8086
{"Status":"OK","Version":"Profile V2"}
```
