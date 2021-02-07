#!/bin/bash

kubectl apply -f secrets.yaml
kubectl apply -f configs.yaml
kubectl apply -f producer.yaml
kubectl apply -f consumer.yaml

kubectl get all
