# 自定义资源

1. 添加 CRD
2. 添加鉴权钩子

![.](https://iximiuz.com/kubernetes-api-how-to-extend/extending-kubernetes-api-2000-opt.png)

## CRD

crd 请求处理流程

![.](https://iximiuz.com/kubernetes-api-how-to-extend/controller-2000-opt.png)

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: monkeys.chaos.iximiuz.com
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: chaos.iximiuz.com
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: monkeys
    # singular name to be used as an alias on the CLI and for display
    singular: monkey
    # kind is normally the CamelCased singular name of the schema
    kind: Monkey
  scope: Namespaced
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        # Monkey objects will have just two properties: .count and .selector
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                count:
                  type: number
                  minimum: 1
                selector:
                  type: string
                  maxLength: 1024
```

## 鉴权钩子

- Authentication & Authorization
- Mutating Admission
- Object Schema Validation
- Validating Admission

![.](https://iximiuz.com/kubernetes-api-how-to-extend/webhooks-2000-opt.png)
