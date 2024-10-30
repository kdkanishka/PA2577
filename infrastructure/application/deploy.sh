#!/bin/bash

kubectl apply -f mongodb-configmap.yaml
kubectl apply -f mongodb-pvc.yaml
kubectl apply -f mongodb-deployment.yaml
kubectl apply -f mongodb-service.yaml
kubectl apply -f shoppinglist-service-deployment.yaml
kubectl apply -f shoppinglist-service-service.yaml
kubectl apply -f shoppinglist-api-deployment.yaml
kubectl apply -f shoppinglist-api-service.yaml
