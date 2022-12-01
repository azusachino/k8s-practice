# Go Client

## `k8s.io/api` module

- Huge - 1000+ structs describing Kubernetes API objects.
- Simple - almost no algorithms, only "dumb" data structures.
- Useful - its data types are used by clients, servers, controllers, etc.

## `k8s.io/apimachinery` module

This library is a shared dependency for servers and clients to work with Kubernetes API infrastructure without direct type dependencies. Its first consumers are k8s.io/kubernetes, k8s.io/client-go, and k8s.io/apiserver.

![.](https://iximiuz.com/kubernetes-api-go-types-and-common-machinery/k8s-api-and-apimachinery-2000-opt.png)
