#!/bin/bash

# setup argocd here, assuming helm3 is installed and kubectl is configured to the correct cluster

helm repo add argo https://argoproj.github.io/argo-helm
     helm install my-argocd argo/argo-cd \
         --namespace argocd \
         --create-namespace \
         --set global.image.repository="alinbalutoiu/argocd" \
         --set installCRDs="false" \
         --wait

