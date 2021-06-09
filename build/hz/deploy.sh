#!/bin/bash

kubectl apply -f deployment.yaml -n ifp
kubectl apply -f service.yaml -n ifp
kubectl apply -f ingress.yaml -n ifp
