#!/bin/bash

# Deploy the mongoDB
kubectl apply -f mongodb-configmap.yaml
kubectl apply -f mongodb-pvc.yaml
kubectl apply -f mongodb-service.yaml
kubectl apply -f mongodb-deployment.yaml

# Deploy the shopping list service component
kubectl apply -f shoppinglist-service-service.yaml
kubectl apply -f shoppinglist-service-deployment.yaml

# Deploy the shopping list API component
kubectl apply -f shoppinglist-api-service.yaml
kubectl apply -f shoppinglist-api-deployment.yaml
