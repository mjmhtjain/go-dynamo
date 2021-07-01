### Kubernetes dashboard command

```
kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml

kubectl proxy

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

kubectl create serviceaccount dashboard-admin-sa

kubectl create clusterrolebinding dashboard-admin-sa \
  --clusterrole=cluster-admin \
  --serviceaccount=default:dashboard-admin-sa

kubectl get secrets

kubectl describe secret dashboard-admin-sa-token-4rs96
```