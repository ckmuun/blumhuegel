echo "adding pulsar chart repo to local helm"
helm repo add pulsar https://pulsar.apache.org/charts
helm repo update

echo "pulling helm chart deps"
helm dependency update

echo "helm release will be installed in the followin kubectl k8s context"
kubectl config current-context

echo "setting up namesapce via specific pulsar helm setup script"
prepare_helm_release.sh

helm install --values examples/values-minikube.yaml --set initialize=true --namespace pulsar pulsar-mini apache-pulsar/pulsar

echo "followin pulsar pods created"
kubectl get pods -n pulsar


