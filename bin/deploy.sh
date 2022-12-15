kubectl create configmap notifier-config  --from-env-file=../default.properties
kubectl apply -f deployment.yaml