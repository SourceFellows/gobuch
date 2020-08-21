
```
eval $(minikube docker-env)
docker build -t minikube-golang-sample .

kubectl create -f manifest.yml
kubectl get pod
kubectl expose deployment golang-sample --type=NodePort
kubectl get service
minikube service golang-sample --url
```
