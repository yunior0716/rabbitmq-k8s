# Deploy apps to docker hub

docker login

docker build . -t tuusuariodedocker/nombredelaapp:latest

docker push tuusuariodedocker/nombredelaapp:latest

# Get started

minikube start

kubectl config use-context minikube

kubectl create ns rabbits

# Deploy

kubectl apply -n rabbits -f .\kubernetes\rabbit-rbac.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-configmap.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-secret.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-statefulset.yaml

# Acces to rabbit

kubectl -n rabbits port-forward rabbitmq-0 8080:15672

Go to htttp://localhost:8080
Username: guest
Password: guest

# Deploy apps

cd nombredelacarpetadelaapp

kubectl apply -n rabbits -f nombredelarchivo.yaml
