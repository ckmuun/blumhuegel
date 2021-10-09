#!/bin/bash
# steps from https://vectorized.io/docs/quick-start-kubernetes/

# setup redpanda helm repo
helm repo add redpanda https://charts.vectorized.io/ && \
helm repo update

export REDPANDA_VERSION
REDPANDA_VERSION=$(curl -s https://api.github.com/repos/vectorizedio/redpanda/releases/latest | jq -r .tag_name)


kubectl apply \
-k https://github.com/vectorizedio/redpanda/src/go/k8s/config/crd?ref=REDPANDA_VERSION


helm install \
--namespace redpanda-system \
--create-namespace redpanda-system \
--version "$REDPANDA_VERSION" \
redpanda/redpanda-operator
