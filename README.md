# MyService

Goals: To create my own service from scratch, will add all my learnings here. 

Language: Go

Tasks: Create a basic http methods for movies information.

Things to DO:
Middleware, Authentication, Loggings, Errors, add metrics, Integrate grafana, cicd, argocd, jaegar, prometheus, Database, Validation, Testing(unit, integration), Readme update

Major goal:
Scale the service for millions requests, note the latency and many more

-- Docker --
```sh
- Build the image
docker build -t jinxankit/my-service:latest .
docker tag jinxankit/my-service:latest jinxankit/my-service:v1.0
docker push jinxankit/my-service:v1.0

- Run 
docker run -p 8080:8080 jinxankit/my-service:v1.0
```

- Let's try to test the health of the service:
```sh
curl -X GET http://localhost:8080/health
```

- Steps to deploy the service to kubernetes:
```sh
minikube start

kubectl apply -f k8s-deployment.yaml
kubectl get svc
kubectl get pods
kubectl describe pod
kubectl port-forward deployment/my-service 8080:8080
OR
kubectl port-forward <pod-name> 8080:8080
```

- Steps to test the service:
```sh
curl -X GET http://localhost:8080/health
```

Port forward is a way to test the service locally, but in production, you would want to expose the pod using services.

Write the k8s service:
```sh
kubectl apply -f k8s-deployment.yaml
minikube service my-service --url
```

- Steps to scale a kubernetes deployment (Move to Top)
```sh
kubectl scale --replicas=4 deployment/my-service
```

```sh
- Delete a pod:
kubectl delete pod <pod-name>
- Delete a deployment:
kubectl delete deployment <deployment-name>
- Delete a service:
kubectl delete service <service-name>

- Stop and Delete the minikube cluster:
minikube stop
minikube delete
```
