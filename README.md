# Deploy apps to docker hub

docker login

docker build . -t tuusuariodedocker/nombredelaapp:latest

docker push tuusuariodedocker/nombredelaapp:latest

# Get started

minikube start

kubectl create ns nombredelnamespace

# Deploy

kubectl apply -n nombredelnamespace -f .\kubernetes\rabbit-rbac.yaml

kubectl apply -n nombredelnamespace -f .\kubernetes\rabbit-configmap.yaml

kubectl apply -n nombredelnamespace -f .\kubernetes\rabbit-secret.yaml

kubectl apply -n nombredelnamespace -f .\kubernetes\rabbit-statefulset.yaml

# Acces to rabbit

kubectl -n nombredelnamespace port-forward rabbitmq-0 8080:15672

Go to htttp://localhost:8080
Username: guest
Password: guest

# Deploy apps

cd nombredelacarpetadelaapp

kubectl apply -n nombredelnamespace -f nombredelarchivo.yaml
