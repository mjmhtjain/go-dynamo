#!/bin/bash
set -e

# Setup K8 cluster and registry
NAME=goDynamo
KUBE_RESOURCE_GROUP="${NAME}ResourceGroup"
KUBE_LOCATION=eastus

KUBE_ACR="${NAME}Registry"
DOCKER_IMAGE_TAG=go-dynamo:latest

KUBE_AKS="${NAME}Cluster"
KUBE_DNS_PREFIX="${NAME}DNS"
AKS_POD="${NAME}Pod"

az group create \
    --name $KUBE_RESOURCE_GROUP \
    --location $KUBE_LOCATION \
    | jq

# Create a registry
az acr create --resource-group $KUBE_RESOURCE_GROUP \
  --location $KUBE_LOCATION \
  --name $KUBE_ACR \
  --sku Basic \
  --admin-enabled true

# push image to ACR
az acr import  \
  -n $KUBE_ACR \
  --source docker.io/mjmhtjain/$DOCKER_IMAGE_TAG \
  --image $DOCKER_IMAGE_TAG

# create AKS cluster
az aks create \
  --resource-group $KUBE_RESOURCE_GROUP \
  --name $KUBE_AKS \
  --vm-set-type VirtualMachineScaleSets \
  --attach-acr $KUBE_ACR \
  --dns-name-prefix $KUBE_DNS_PREFIX \
  --generate-ssh-keys \
  --load-balancer-sku standard

az aks get-credentials \
  --resource-group=$KUBE_RESOURCE_GROUP \
  --name=$KUBE_AKS \
  --overwrite-existing

# deploy image to AKS instance
kubectl apply -f deployment.yaml

kubectl get pods --watch