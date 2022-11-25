# k8s-practice

k8s practices

## Structure

![k8s-apis](https://iximiuz.com/working-with-kubernetes-api/kdpv.png)

## Api Resources

![k8s-apis-objects](https://iximiuz.com/kubernetes-api-structure-and-terminology/resource-types-kinds-objects-2000-opt.png)

```sh
$ kubectl api-resources
NAME                     SHORTNAMES   APIVERSION   NAMESPACED   KIND
bindings                              v1           true         Binding
componentstatuses        cs           v1           false        ComponentStatus
configmaps               cm           v1           true         ConfigMap
endpoints                ep           v1           true         Endpoints
events                   ev           v1           true         Event
limitranges              limits       v1           true         LimitRange
namespaces               ns           v1           false        Namespace
nodes                    no           v1           false        Node
persistentvolumeclaims   pvc          v1           true         PersistentVolumeClaim
persistentvolumes        pv           v1           false        PersistentVolume
pods                     po           v1           true         Pod
...
```

### What is Kind?

As per [sig-architecture/api-conventions.md](https://github.com/kubernetes/community/blob/7f3f3205448a8acfdff4f1ddad81364709ae9b71/contributors/devel/sig-architecture/api-conventions.md#types-kinds), kinds are grouped into three categories:

- Objects (Pod, Service, etc) - persistent entities in the system.
- Lists - (PodList, APIResourceList, etc) - collections of resources of one or more kinds.
- Simple - specific actions on objects (status, scale, etc.) or non-persistent auxiliary entities (ListOptions, Policy, etc).

### The K8S Objects?

Entities like ReplicaSet, Namespace, or ConfigMap are called [Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/). Objects are persistent entities in the Kubernetes system that represent an intent (desired state) and the status (actual state) of the cluster.

Most of the Kubernetes API resources represent Objects. Unlike other forms of resources mandating only the kind field, Objects must have more field defined:

- kind - a string that identifies the schema this object should have
- apiVersion - a string that identifies the version of the schema the object should have
- metadata.namespace - a string with the namespace (defaults to "default")
- metadata.name - a string that uniquely identifies this object within the current namespace
- metadata.uid - a unique in time and space value used to distinguish between objects with the same name that have been deleted and recreated.

```sh
kubectl get --raw /api/v1/namespaces/kube-system/pods/etcd-minikube | python3 -m json.tool
``

## References

- [Working with Kubernetes API](https://iximiuz.com/en/series/working-with-kubernetes-api/)
