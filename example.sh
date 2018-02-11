#!/bin/bash

read -p 'Start'

#kubectl create -f prime.yaml
#read -p 'Deployed the app'

kubectl apply -f update-1.yaml
read -p 'Updated to version 3'

kubectl apply -f update-2.yaml
read -p 'Updated to version 2'

kubectl apply -f reset.yaml
echo 'Updated back to version 1'
