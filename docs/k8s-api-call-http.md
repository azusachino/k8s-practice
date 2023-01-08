# HTTP Kubernetes

## Instruction with `curl`

### 1. Check the kubeconfig

```sh
kubectl config view
```

### 2. Grab the api server

```sh
KUBE_API=$(kubectl config view -o jsonpath='{.clusters[0].cluster.server}')
```

### 3. Try to use `curl`

```sh
# will fail the ssl check
curl $KUBE_API/version

## load the local ca.crt
curl --cacert ~/.minikube/ca.crt $KUBE_API/version
```

But when `curl $KUBE_API/ap1s/apps/v1/deployments`, we got declined, authentication failed.

为访问 k8s 集群，需要知名 CA 证书，以及当前 CA 为(minikube)用户颁发的密钥，才能访问后台的业务数据.

```sh
curl $KUBE_API/apis/apps/v1/deployments \
  --cacert ~/.minikube/ca.crt \
  --cert ~/.minikube/profiles/minikube/client.crt \
  --key ~/.minikube/profiles/minikube/client.key
```

或者通过 kubectl 创建 service account token

```sh
# k8s 1.24+
JWT_TOKEN_DEFAULT_DEFAULT=$(kubectl create token default)

# operation
curl $KUBE_API/apis/apps/v1/ \
  --cacert ~/.minikube/ca.crt  \
  --header "Authorization: Bearer $JWT_TOKEN_DEFAULT_DEFAULT"
```

但，`system:serviceaccount:default:default` 的权限不够，甚至无法查询 Objects 信息，再用 `kube-system` 的 token 尝试一下

```sh
JWT_TOKEN_KUBESYSTEM_DEFAULT=$(kubectl -n kube-system create token default)

curl $KUBE_API/apis/apps/v1/deployments \
  --cacert ~/.minikube/ca.crt  \
  --header "Authorization: Bearer $JWT_TOKEN_KUBESYSTEM_DEFAULT"
```

### 通过 CURL 操作 k8s 数据

```sh
curl $KUBE_API/apis/apps/v1/namespaces/default/deployments \
  --cacert ~/.minikube/ca.crt \
  --cert ~/.minikube/profiles/cluster1/client.crt \
  --key ~/.minikube/profiles/cluster1/client.key \
  -X POST \
  -H 'Content-Type: application/yaml' \
  -d '---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sleep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sleep
  template:
    metadata:
      labels:
        app: sleep
    spec:
      containers:
      - name: sleep
        image: curlimages/curl
        command: ["/bin/sleep", "365d"]
'
```
