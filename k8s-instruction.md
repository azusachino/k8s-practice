# k8s

## minikube

### instructions

```sh
# 1. 指定镜像地址为 aliyun
# 2. 指定 driver
# 3. 指定 runtime
# 4. 输出错误日志，以进行排查
minikube start --image-mirror-country='cn' --driver=docker --container-runtime=containerd --alsologtostderr --v=5

# 若流程出现错误，可以删除后重试
minikube delete
```

## References

- [k8s get start](https://minikube.sigs.k8s.io/docs/start/)
