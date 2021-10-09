
echo "install certman via helm"
helm repo add jetstack https://charts.jetstack.io && \
helm repo update && \
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.2.0 \
  --set installCRDs=true

echo "add vectorized redpanda helm repo"
helm repo add redpanda https://charts.vectorized.io/ && \
helm repo update

echo "get redpanda version"
export VERSION=$(curl -s https://api.github.com/repos/vectorizedio/redpanda/releases/latest | jq -r .tag_name)

echo "get redpanda crd"
kubectl apply \
-k https://github.com/vectorizedio/redpanda/src/go/k8s/config/crd?ref=$VERSION

echo "install redpanda operator via helm"
helm install \
--namespace redpanda-system \
--create-namespace redpanda-system \
--version $VERSION \
redpanda/redpanda-operator

echo "create namespace blumhuegel"
kubectl create ns blumhuegel
echo "apply one-one_node_cluster redpanda deployment"
kubectl apply \
-n blumhuegel \
-f https://raw.githubusercontent.com/vectorizedio/redpanda/dev/src/go/k8s/config/samples/one_node_cluster.yaml

echo "create topic 'blumhuegel' in redpanda cluster"
kubectl -n blumhuegel run -ti --rm \
--restart=Never \
--image docker.vectorized.io/vectorized/redpanda:$VERSION \
-- rpk --brokers one-node-cluster-0.one-node-cluster.blumhuegel.svc.cluster.local:9092 \
topic create basic-financials -p 3

echo "do port-forward"
kubectl port-forward -n blumhuegel service/one-node-cluster 9092 9644 8082

